package test

import (
	"fmt"
	"github.com/cg-/go-hello-bench/common"
)

/*
Dependency is a type that represents a program that must be installed on a system in order for a benchmark to function.
*/

type Dependency struct {
	Name           string
	check          Command
	isInstalledStr string
	install        Command
	statusCode     int
	additionalDeps []Dependency
}

type Command struct {
	cmd  string
	args []string
}

var SKIPCMD = Command{cmd: "", args: []string{}}
var SKIPSTR = ""
var SKIPEXIT = -1

func NewCommandDependency(name, command string, args []string) Dependency {
	return NewCommandDependencyExtra(name, command, args, []Dependency{})
}

func NewCommandDependencyExtra(name, command string, args []string, cmds []Dependency) Dependency {
	return Dependency{
		Name:           name,
		check:          SKIPCMD,
		isInstalledStr: SKIPSTR,
		install: Command{
			cmd:  command,
			args: args,
		},
		statusCode:     SKIPEXIT,
		additionalDeps: cmds,
	}
}

func NewAptDependency(name string) Dependency {
	return NewAptDependencyExtra(name, []Dependency{})
}

func NewAptDependencyExtra(name string, cmds []Dependency) Dependency {
	return Dependency{
		Name: name,
		check: Command{
			cmd:  "dpkg-query",
			args: []string{"-W", "-f=${Status}", name},
		},
		isInstalledStr: "install ok installed",
		install: Command{
			cmd:  "apt",
			args: []string{"install", "-y", name},
		},
		statusCode:     SKIPEXIT,
		additionalDeps: cmds,
	}
}

func NewFileDependency(path, url string) Dependency {
	return NewFileDependencyExtra(path, url, []Dependency{})
}

func NewFileDependencyExtra(path, url string, cmds []Dependency) Dependency {
	return Dependency{
		Name: path,
		check: Command{
			cmd:  "test",
			args: []string{"-f", path},
		},
		isInstalledStr: "",
		install: Command{
			cmd:  "wget",
			args: []string{"-q", url, "-O", path},
		},
		statusCode:     0,
		additionalDeps: cmds,
	}
}

func NewDirDependency(path string) Dependency {
	return NewDirDependencyExtra(path, []Dependency{})
}

func NewDirDependencyExtra(path string, cmds []Dependency) Dependency {
	return Dependency{
		Name: path,
		check: Command{
			cmd:  "test",
			args: []string{"-d", path},
		},
		isInstalledStr: "",
		install: Command{
			cmd:  "mkdir",
			args: []string{"-p", path},
		},
		statusCode:     0,
		additionalDeps: cmds,
	}
}

/*
Check returns true if a dependency is satisfied, or false if not.
*/
func (d *Dependency) Check() bool {
	// Check if we want to skip this command
	if d.check.cmd == "" {
		return true
	}

	testedSomething := false
	log.Infof("Checking dependency %s", d.Name)

	// Run check command
	err, outputStr, _ := common.RunCmd(d.check.cmd, d.check.args)

	// Check status code (if specified-- -1 means skip this test)
	if d.statusCode != SKIPEXIT {
		testedSomething = true

		// Exit code 0 will return nil...
		if d.statusCode == 0 && err == nil {
			log.Infof("Status code check [0 vs 0].")
			return true
		}

		log.Infof("Status code check [%s vs %d].", err.Error(), d.statusCode)
		if fmt.Sprintf("exit status %d", d.statusCode) != err.Error() {
			return false
		}
	}

	// Check output text (if specified-- empty string means skip this test)
	log.Infof("Comparing %s to %s.", outputStr, d.isInstalledStr)
	if d.isInstalledStr != SKIPSTR {
		testedSomething = true
		if outputStr == d.isInstalledStr {
			return true
		}
		return false
	}

	if testedSomething {
		return true
	} else {
		log.Fatal("No dependency tests to check! Crashing.")
		return false // This will never be reached
	}
}

/*
InstallLocal will set up a dependency on a target local system.
*/
func (d *Dependency) InstallLocal() {
	log.Infof("Installing dependency %s", d.Name)
	err, outputStr, errorStr := common.RunCmd(d.install.cmd, d.install.args)
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
