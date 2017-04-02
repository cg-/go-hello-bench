/*
Random utility functions used in various packages
*/

package common

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("go-hello-bench")

/*
CheckIfSuperuser returns true if the user is running as Superuser, otherwise
it will return false.
*/
func CheckIfSuperuser() bool {
	if os.Getegid() != 0 {
		return false
	}
	return true
}

/*
IsInStringSlice will return true if a string matches a string inside a slice,
otherwise it will return false.
*/
func IsInStringSlice(s string, slice []string) bool {
	for i := range slice {
		if slice[i] == s {
			return true
		}
	}
	return false
}

/*
CheckFatal checks an error and crashes the program if it exists with a message to the log.
*/
func CheckFatal(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

/*
CheckFatal checks an error and notes a message to the log if it exists.
*/
func CheckLog(e error) {
	if e != nil {
		log.Error(e.Error())
	}
}

/*
RunCmd runs a command using exec, and returns what was printed to Stdout and Stderr as strings. If an error was
encountered, it is returned as well.
*/
func RunCmd(cmd string, args []string) (err error, outputStr, errorStr string) {
	var outputBuf bytes.Buffer
	var errorBuf bytes.Buffer
	argsStr := ""
	for i := range args {
		argsStr += args[i]
		if i != len(args)-1 {
			argsStr += " "
		}
	}
	fullCmd := cmd + " " + argsStr
	log.Infof("About to exec: [%s]", fullCmd)
	cmdVar := exec.Command(cmd, args...)
	cmdVar.Stdout = &outputBuf
	cmdVar.Stderr = &errorBuf
	errVar := cmdVar.Run()
	if errVar != nil {
		log.Infof("Got back err[%s], out[%s], err[%s]", errVar.Error(), outputBuf.String(), errorBuf.String())
	} else {
		log.Infof("Got back err[nil], out[%s], err[%s]", outputBuf.String(), errorBuf.String())
	}
	return errVar, outputBuf.String(), errorBuf.String()
}
