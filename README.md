# Go Playground

Some Go code whilst I get to grips with Go and integration with C libraries.

## Example

    $ export GOPATH=/Users/paul/Work/go-stuff/sqlite
    $ cd $GOPATH
    $ go install example
    $ tree
    .
    ├── bin
    │   └── example
    └── src
        └── example
            └── example.go
    $ otool -L bin/example 
    bin/example:
     /usr/lib/libsqlite3.dylib (compatibility version 0.0.0, current version 0.0.0)
     /usr/lib/libSystem.B.dylib (compatibility version 0.0.0, current version 0.0.0)
    $ ./bin/example 
    libsqlite3 version 3.7.12
