FROM golang:1.13.8-buster
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
LABEL maintainer="Siddhartha Basu <siddhartha-basu@northwestern.edu>"
ENV GOPROXY https://proxy.golang.org
RUN mkdir -p /graphql-authserver
WORKDIR /graphql-authserver
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD cmd cmd
ADD internal internal
RUN CGO_ENABLED=0 go build -o app cmd/graphql-authserver/main.go

FROM gcr.io/distroless/static
COPY --from=0 /graphql-authserver/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]