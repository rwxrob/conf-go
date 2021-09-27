/*

Package conf provided a simple, persistable configuration map that is
safe for concurrency and parses much more efficiently than JSON or YAML
for simple key=value data. It provides reasonable defaults based on
existing standards and allows some flexibility in how configuration data
is organized under these standards (notably XDG_CONFIG_HOME).

Overview

The core component of this package is the Map interface which is
composed of several smaller interfaces to facilitate documentation and
extensibility. This allows the implementation of structs that fulfill
the Map interface in whatever way makes sense for a given application.

A default Map reference implementation is provided and returned by the
NewMap function. The returned struct contains current information about
itself within its cached metadata (which can be changed later). This
includes the Home, Name, and File all of which are accessible via their
respective accessor/mutator interfaces.

The default Home configuration directory is the returned value of
os.UserConfigDir (a well recognized standard that observes HOME if set).

The default Name associated with a Map is the name of the
os.Executable itself. This Name matches the name of the subdirectory
within the home configuration directory.

The default File name is the string "values". Multiple Maps can have
different files with the same Name allowing them to be grouped under the
same configuration subdirectory.  The values file format is simply one
key=value pair per line with a simple equal sign delimiter. (See the
Parse interface for a more detailed description.) This format can be
easily handled by most any simple scripting language and is 100%
compatible with the Java properties specification. When line and
carriage returns are expected the Escape utility function can be used
before setting any key or value.

Editing

Files can be edited directly with any text editor. The Edit method will
also pass the Map.Path to any editor accessible from the command line.

*/
package conf
