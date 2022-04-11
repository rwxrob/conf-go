# Simple, Persistent, Configuration Maps in Go

***DEPRECATED:*** This module is no longer maintained and inferior to
the new [conf](https://github.com/rwxrob/conf) and
[vars](https://github.com/rwxrob/vars) Bonzai branches.

"It's Redis without the Redis."

[![GoDoc](https://godoc.org/github.com/rwxrob/conf-go?status.svg)](https://godoc.org/github.com/rwxrob/conf-go)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/conf-go)](https://goreportcard.com/report/github.com/rwxrob/conf-go)
[![Coverage](https://gocover.io/_badge/github.com/rwxrob/conf-go)](https://gocover.io/github.com/rwxrob/conf-go)


> ⚠️
> A previous version of this package (with the same name) was based on
> a JSON approach.

Rather than use JSON this library assumes one string configuration key
and value pair per line. This not only simplifies the keys, but parses
much more quickly than JSON and is compatible with traditional
UNIX-style configuration parsing from shell scripts and the command line
as well as full compatibility with Java Properties parsers. In fact, it
is exactly the same method used in the [`template-bash-command`][bash]
allowing bash scripts and Go utilities using this package to be 100%
compatible in their behavior.

Click on the GoDoc badge above for more information.

[bash]: <https://github.com/rwxrob/template-bash-command>

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).
