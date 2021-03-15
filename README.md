# Simple, Stateful JSON Configs in Go

[![GoDoc](https://godoc.org/github.com/rwxrob/conf-go?status.svg)](https://godoc.org/github.com/rwxrob/conf-go)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/github.com/rwxrob/conf-go)](https://goreportcard.com/report/github.com/rwxrob/conf-go)

"It's Redis without the Redis."

## Examples 

The following projects use `conf-go` in different ways:

Project|Description
|:-:|-
[`cmdtab-config`] | Access to convention JSON config data 
[`cmdtab-pomo`] | Pomodoro countdown suitable for TMUX
[`cmdtab-timer`] | Generic timer suitable for TMUX
[`kn`] | KEG knowledge management utility

[`config`]: https://github.com/rwxrob/cmdtab-config

## Testing

Given the nature of testing required for something that fundamentally
depends on a user account home directory and the proper environment
normal Go unit testing has been replaced with test cases to be executed
by humans with the desired result (until, and if, a more suitable
container-based test suite can be incorporated into this repo). The
dependent [Examples](#examples) provide further opportunity for this testing.
