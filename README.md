# GraphQL Authserver

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictyBase/graphql-authserver/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dictyBase/graphql-authserver)](https://goreportcard.com/report/github.com/dictyBase/graphql-authserver)
[![Technical debt](https://badgen.net/codeclimate/tech-debt/dictyBase/graphql-authserver)](https://codeclimate.com/github/dictyBase/graphql-authserver/trends/technical_debt)
[![Issues](https://badgen.net/codeclimate/issues/dictyBase/graphql-authserver)](https://codeclimate.com/github/dictyBase/graphql-authserver/issues)
[![Maintainability](https://api.codeclimate.com/v1/badges/21ed283a6186cfa3d003/maintainability)](https://codeclimate.com/github/dictyBase/graphql-authserver/maintainability)  
![Issues](https://badgen.net/github/issues/dictyBase/graphql-authserver)
![Open Issues](https://badgen.net/github/open-issues/dictyBase/graphql-authserver)
![Closed Issues](https://badgen.net/github/closed-issues/dictyBase/graphql-authserver)  
![Total PRS](https://badgen.net/github/prs/dictyBase/graphql-authserver)
![Open PRS](https://badgen.net/github/open-prs/dictyBase/graphql-authserver)
![Closed PRS](https://badgen.net/github/closed-prs/dictyBase/graphql-authserver)
![Merged PRS](https://badgen.net/github/merged-prs/dictyBase/graphql-authserver)  
![Commits](https://badgen.net/github/commits/dictyBase/graphql-authserver/develop)
![Last commit](https://badgen.net/github/last-commit/dictyBase/graphql-authserver/develop)
![Branches](https://badgen.net/github/branches/dictyBase/graphql-authserver)
![Tags](https://badgen.net/github/tags/dictyBase/graphql-authserver/?color=cyan)  
![GitHub repo size](https://img.shields.io/github/repo-size/dictyBase/graphql-authserver?style=plastic)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dictyBase/graphql-authserver?style=plastic)
[![Lines of Code](https://badgen.net/codeclimate/loc/dictyBase/graphql-authserver)](https://codeclimate.com/github/dictyBase/graphql-authserver/code)  
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,dictyBase/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9476993)
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,DSC/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9438930)

HTTP server for authorizing GraphQL mutations.

## Usage

```
NAME:
   graphql-authserver - cli for graphql-authserver

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   start-server  starts the graphql authserver
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-format value  format of the logging out, either of json or text. (default: "json")
   --log-level value   log level for the application (default: "error")
   --help, -h          show help
   --version, -v       print the version
```

## Subcommand

```
NAME:
   main start-server - starts the graphql authserver

USAGE:
   main start-server [command options] [arguments...]

OPTIONS:
   --private value, --pr value  output file name for private key
   --public value, --pub value  output file name for public key
   --port value, -p value       server port (default: 9099)
```

