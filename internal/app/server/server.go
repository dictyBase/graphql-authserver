package server

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	loggerMw "github.com/dictyBase/go-middlewares/middlewares/logrus"
	"gopkg.in/urfave/cli.v1"

	"github.com/dgrijalva/jwt-go"
	"github.com/dictyBase/graphql-authserver/internal/handlers"
	mw "github.com/dictyBase/graphql-authserver/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

// Runs the http server
func RunServer(c *cli.Context) error {
	jt, err := parseJwtKeys(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("unable to parse keys %q\n", err), 2)
	}
	loggerMw, err := getLoggerMiddleware(c)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("unable to get logger middleware %s", err), 2)
	}
	cors := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
	})
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(loggerMw.Middleware)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)
	// Default health check
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("okay"))
	})
	r.Route("/watchmen", func(r chi.Router) {
		tokenAuth := jwtauth.New("RS512", jt.SignKey, jt.VerifyKey)
		r.With(mw.AuthorizeMiddleware).
			With(jwtauth.Verifier(tokenAuth)).
			Post("/", jt.JwtFinalHandler)
	})
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), r))
	return nil
}

// Reads the public and private keys and creates a new JWTAuth instance.
func parseJwtKeys(c *cli.Context) (*handlers.Jwt, error) {
	jh := &handlers.Jwt{}
	private, err := base64.StdEncoding.DecodeString(c.String("private-key"))
	if err != nil {
		return jh, err
	}
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return jh, err
	}
	public, err := base64.StdEncoding.DecodeString(c.String("public-key"))
	if err != nil {
		return jh, err
	}
	pubkey, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return jh, err
	}
	jh.VerifyKey = pubkey
	jh.SignKey = pkey
	return jh, err
}

// GetLoggerMiddleware gets a net/http compatible instance of logrus
func getLoggerMiddleware(c *cli.Context) (*loggerMw.Logger, error) {
	var logger *loggerMw.Logger
	var w io.Writer
	if c.IsSet("log-file") {
		fw, err := os.Create(c.String("log-file"))
		if err != nil {
			return logger,
				fmt.Errorf("could not open log file  %s %s", c.String("log-file"), err)
		}
		w = io.MultiWriter(fw, os.Stderr)
	} else {
		w = os.Stderr
	}
	if c.String("log-format") == "json" {
		logger = loggerMw.NewJSONFileLogger(w)
	} else {
		logger = loggerMw.NewFileLogger(w)
	}
	return logger, nil
}
