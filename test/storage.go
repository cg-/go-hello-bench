/*
Contains various storage benchmarks
*/

package test

/*
TestTest is a test
*/
var TestTest = Test{
	Name:   "Test Test",
	Prereq: nil,
	Deps: []Dependency{
		NewFileDependency("/tmp/asdasdasd", "http://www.example.com/index.html"),
	},
}

/*
Common sets up common files for all tests.
*/
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

/*
Filebench is the filebench benchmarking software.
*/
var Filebench = Test{
	Name:   "Filebench",
	Prereq: &Common,
	Deps: []Dependency{
		NewAptDependency("build-essential"),
	},
}
