# `conf-go` Lightweight Configuration Library 

[![GoDoc](https://godoc.org/conf-go?status.svg)](https://godoc.org/conf-go)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Rather than use JSON and fight with JSON map keys this library assumes
one string configuration key and value pair per line. This not only
simplifies the keys, but parses much more quickly than JSON and is
compatible with traditional UNIX-style configuration parsing from shell
scripts and the command line. In fact, it is exactly the same method
used in the [`template-bash-command`][bash] allowing bash scripts and Go
utilities using this package to be 100% compatible in their behavior.

[bash]: <https://github.com/rwxrob/template-bash-command>

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).
