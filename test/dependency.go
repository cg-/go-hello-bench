package test

import (
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
   common.Logger.Println("Checking dependency %s", d.Name)
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
func (d *Dependency) InstallLocal() {
   common.Logger.Println("Installing dependency %s", d.Name)
   common.CheckFatal(fmt.Errorf("Not implemented."))
}

/*
InstallDocker will add the required lines to the specified container's Dockerfile.
todo: Implement Docker
 */
func (d *Dependency) InstallDocker() {
   common.Logger.Println("Installing dependency %s in Docker", d.Name)
   common.CheckFatal(fmt.Errorf("Not implemented."))
}
