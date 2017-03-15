/*
Random utility functions used in various packages
 */

package common

import (
   "bytes"
   "os/exec"
)

/*
CheckFatal checks an error and crashes the program if it exists with a message to the log.
 */
func CheckFatal(e error){
   if e != nil{
      Logger.Fatal(e.Error())
   }
}

/*
CheckFatal checks an error and notes a message to the log if it exists.
 */
func CheckLog(e error){
   if e != nil{
      Logger.Println(e.Error())
   }
}

/*
RunCmd runs a command using exec, and returns what was printed to Stdout and Stderr as strings. If an error was
encountered, it is returned as well.
 */
func RunCmd(cmd string)(err error, outputStr, errorStr string){
   var outputBuf bytes.Buffer
   var errorBuf bytes.Buffer
   cmdVar := exec.Command(cmd)
   cmdVar.Stdout = outputBuf
   cmdVar.Stderr = errorBuf
   errVar := cmdVar.Run()
   return errVar, outputBuf.String(), errorBuf.String()
}
