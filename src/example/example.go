package main

// #cgo LDFLAGS: -lsqlite3
// #include <stdio.h>
// #include <stdlib.h>
// #include <sqlite3.h>
import "C"
import "fmt"

func main() {
  version := C.GoString(C.sqlite3_libversion())
  
  fmt.Println("libsqlite3 version", version)
}
