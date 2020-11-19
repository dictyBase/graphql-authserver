# GraphQL Authserver
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictyBase/graphql-authserver/workflows/Build/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/21ed283a6186cfa3d003/maintainability)](https://codeclimate.com/github/dictyBase/graphql-authserver/maintainability)  
![Last commit](https://badgen.net/github/last-commit/dictyBase/graphql-authserver/develop)
[![Funding](https://badgen.net/badge/Funding/Rex%20L%20Chisholm,dictyBase,DCR/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=10024726&icde=0)

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

# Misc badges
![Issues](https://badgen.net/github/issues/dictyBase/graphql-authserver)
![Open Issues](https://badgen.net/github/open-issues/dictyBase/graphql-authserver)
![Closed Issues](https://badgen.net/github/closed-issues/dictyBase/graphql-authserver)  
![Total PRS](https://badgen.net/github/prs/dictyBase/graphql-authserver)
![Open PRS](https://badgen.net/github/open-prs/dictyBase/graphql-authserver)
![Closed PRS](https://badgen.net/github/closed-prs/dictyBase/graphql-authserver)
![Merged PRS](https://badgen.net/github/merged-prs/dictyBase/graphql-authserver)  
![Commits](https://badgen.net/github/commits/dictyBase/graphql-authserver/develop)
![Branches](https://badgen.net/github/branches/dictyBase/graphql-authserver)
![Tags](https://badgen.net/github/tags/dictyBase/graphql-authserver/?color=cyan)  
![GitHub repo size](https://img.shields.io/github/repo-size/dictyBase/graphql-authserver?style=plastic)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dictyBase/graphql-authserver?style=plastic)
[![Lines of Code](https://badgen.net/codeclimate/loc/dictyBase/graphql-authserver)](https://codeclimate.com/github/dictyBase/graphql-authserver/code)  

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://www.erichartline.net/"><img src="https://avatars3.githubusercontent.com/u/13489381?v=4" width="100px;" alt=""/><br /><sub><b>Eric Hartline</b></sub></a><br /><a href="https://github.com/dictyBase/graphql-authserver/issues?q=author%3Awildlifehexagon" title="Bug reports">🐛</a> <a href="https://github.com/dictyBase/graphql-authserver/commits?author=wildlifehexagon" title="Code">💻</a> <a href="#content-wildlifehexagon" title="Content">🖋</a> <a href="https://github.com/dictyBase/graphql-authserver/commits?author=wildlifehexagon" title="Documentation">📖</a> <a href="#maintenance-wildlifehexagon" title="Maintenance">🚧</a></td>
    <td align="center"><a href="http://cybersiddhu.github.com/"><img src="https://avatars3.githubusercontent.com/u/48740?v=4" width="100px;" alt=""/><br /><sub><b>Siddhartha Basu</b></sub></a><br /><a href="https://github.com/dictyBase/graphql-authserver/issues?q=author%3Acybersiddhu" title="Bug reports">🐛</a> <a href="https://github.com/dictyBase/graphql-authserver/commits?author=cybersiddhu" title="Code">💻</a> <a href="#content-cybersiddhu" title="Content">🖋</a> <a href="https://github.com/dictyBase/graphql-authserver/commits?author=cybersiddhu" title="Documentation">📖</a> <a href="#maintenance-cybersiddhu" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!