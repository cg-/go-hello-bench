/*
Contains various storage benchmarks
*/

package test

var TestTest = Test{
	Name:   "Test Test",
	Prereq: nil,
	Deps: []Dependency{
		NewFileDependency("/tmp/asdasdasd", "http://www.example.com/index.html"),
	},
}

var Common = Test{
	Name:   "Common Test Files",
	Prereq: nil,
	Deps: []Dependency{
		NewAptDependency("build-essential"),
		NewAptDependency("automake"),
		NewAptDependency("cmake"),
		NewAptDependency("wget"),
		NewAptDependency("tar"),
		NewAptDependency("libtool"),
		NewAptDependency("bison"),
		NewAptDependency("flex"),
		NewAptDependency("git"),
	},
}

var Filebench = Test{
	Name:   "Filebench",
	Prereq: &Common,
	Deps: []Dependency{
		Dependency{
			Name:               "build-essential",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "build-essential"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "build-essential"},
		},
	},
}
