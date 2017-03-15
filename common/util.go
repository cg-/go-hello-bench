/*
Random utility functions used in various packages
 */

package common

import (
   "bytes"
   "os/exec"
   "github.com/op/go-logging"
   "fmt"
)

var log = logging.MustGetLogger("go-hello-bench")

/*
CheckFatal checks an error and crashes the program if it exists with a message to the log.
 */
func CheckFatal(e error){
   if e != nil{
      log.Fatal(e.Error())
   }
}

/*
CheckFatal checks an error and notes a message to the log if it exists.
 */
func CheckLog(e error){
   if e != nil{
      log.Error(e.Error())
   }
}

/*
RunCmd runs a command using exec, and returns what was printed to Stdout and Stderr as strings. If an error was
encountered, it is returned as well.
 */
func RunCmd(cmd string, args []string)(err error, outputStr, errorStr string){
   var outputBuf bytes.Buffer
   var errorBuf bytes.Buffer
   argsStr := ""
   for i := range args{
      argsStr += args[i]
      if i != len(args) - 1{
         argsStr += " "
      }
   }
   fullCmd := cmd + " " + argsStr
   log.Infof("About to exec: [%s]", fullCmd)
   cmdVar := exec.Command(fullCmd)
   cmdVar.Stdout = &outputBuf
   cmdVar.Stderr = &errorBuf
   errVar := cmdVar.Run()
   return errVar, outputBuf.String(), errorBuf.String()
}
