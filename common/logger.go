/*
Handle logging within all packages in the program in a central place.
 */

package common

import (
   "log"
   "bytes"
   "fmt"
)

var Logger *log.Logger
var LogBuffer bytes.Buffer

func init() {
   Logger = log.New(&LogBuffer, "logger: ", log.Lshortfile)
}

func PrintLog() {
   fmt.Println(&LogBuffer)
}
