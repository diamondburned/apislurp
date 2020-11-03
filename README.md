# apislurp

[![pipeline](https://gitlab.com/diamondburned/apislurp/badges/purpleslush/pipeline.svg?style=flat-square)][pipeline]
[![coverage](https://gitlab.com/diamondburned/apislurp/badges/purpleslush/coverage.svg?style=flat-square)][pipeline]
[![godoc](https://img.shields.io/badge/godoc-reference-blue)][godoc]

[pipeline]: https://gitlab.com/diamondburned/apislurp/pipelines
[godoc]: https://beta.pkg.go.dev/github.com/diamondburned/apislurp

A small Discord documentation parser.

## Coverage

This package covers the structures inside `discord-api-docs/resources/` as well
as the `Gateway.md` file inside `discord-api-docs/topics/`.

## Usage

Refer to `internal/tests/{resources,topics}/`.

## License

apislurp is licensed under the Lesser GNU General Public License version 3.0
(LGPLv3). As a special exception (regardless if LGPLv3 covers this or not), all
code generated using apislurp may be licensed without restriction. This includes
proprietary licenses. In other words, the use of apislurp makes no requirements
about the licenses of generated code.
