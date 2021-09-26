/* Package conf provided a simple, persistable configuration map that is
safe for concurrency and parses much more effectively than JSON or YAML
for simple key=value data. It provides reasonable defaults based on
existing standards and allows some flexibility in how configuration data
is organized under these standards (notable XDG_CONFIG_HOME).

When a new Map is returns from NewMap it saves the current information
about itself within its cached metadata (which can be changed later).

The default Home configuration directory is that returned by
os.UserConfigDir (a well recognized standard that observes HOME if set).

The default Name associated with a Map is the name of the
os.Executable itself. This Name matches the name of the subdirectory
within the home configuration directory.

The File is the name of the file within the subdirectory associated with
Name. The canonical default file name is "values". Multiple Maps can
have different files with the same Name allowing them to effectively be
grouped under the same configuration subdirectory.

The values file format is 100% compatible with the Java properties
specification enabling other parsers to be used when necessary. The
format is, however, a very limited subset of the full Java properties
specification. This makes for easy splitting using simple bash parameter
expansion and other split methods from any language. There is no need
for complicated escaping. Everything is verbatim. However, no actual
line return is allowed as a key or value (but can be used if escaped
before being saved using something like the Escape utility function).

*/
package conf
