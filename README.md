# Simple, Persistent, Configuration Maps in Go

"It's Redis without the Redis."

[![GoDoc](https://godoc.org/conf-go?status.svg)](https://godoc.org/conf-go)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

> ⚠️
> A previous version of this package (with the same name) was based on
> a JSON approach. This version breaks from the JSON dependency.

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
