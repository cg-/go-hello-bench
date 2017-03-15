package test

import (
   "os/exec"
   "bytes"
   "github.com/cg-/go-hello-bench/common"
   "fmt"
)

/*
Dependency is a type that represents a program that must be installed on a system in order for a benchmark to function.
 */
type Dependency struct {
   Name string
   checkCommand   string
   isInstalledStr string
   installCommand string
}

/*
Check returns true if a dependency is satisfied, or false if not.
 */
func (d *Dependency) Check() bool {
   err, outputStr, _ := common.RunCmd(d.checkCommand)
   common.CheckFatal(err)
   if outputStr != d.isInstalledStr {
      return true
   }
   return false
}

/*
InstallLocal will set up a dependency on a target local system.
 */
func (d *Dependency) InstallLocal() error {
}

/*
InstallDocker will add the required lines to the specified container's Dockerfile.
todo: Implement Docker
 */
func (d *Dependency) InstallDocker() {
   common.CheckFatal(fmt.Errorf("Not implemented."))
}
