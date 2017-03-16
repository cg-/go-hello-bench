package test

import (
	"fmt"
	"github.com/cg-/go-hello-bench/common"
)

/*
Dependency is a type that represents a program that must be installed on a system in order for a benchmark to function.
*/

type Dependency struct {
	Name               string
	checkCommand       string
	checkCommandArgs   []string
	isInstalledStr     string
	installCommand     string
	installCommandArgs []string
}

/*
Check returns true if a dependency is satisfied, or false if not.
*/
func (d *Dependency) Check() bool {
	log.Infof("Checking dependency %s", d.Name)
	err, outputStr, _ := common.RunCmd(d.checkCommand, d.checkCommandArgs)
	common.CheckFatal(err)
	log.Infof("Comparing %s to %s.", outputStr, d.isInstalledStr)
	if outputStr == d.isInstalledStr {
		return true
	}
	return false
}

/*
InstallLocal will set up a dependency on a target local system.
*/
func (d *Dependency) InstallLocal() {
	log.Infof("Installing dependency %s", d.Name)
	err, outputStr, errorStr := common.RunCmd(d.installCommand, d.installCommandArgs)
	common.CheckFatal(err)
	log.Infof("output: %s error: %s", outputStr, errorStr)
}

/*
InstallDocker will add the required lines to the specified container's Dockerfile.
todo: Implement Docker
*/
func (d *Dependency) InstallDocker() {
	log.Infof("Installing dependency %s in Docker", d.Name)
	common.CheckFatal(fmt.Errorf("Not implemented."))
}
