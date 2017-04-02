package test

/*
Test represents a testing program that can be run.
*/
type Test struct {
	Name   string
	Prereq *Test
	Deps   []Dependency
}

/*
SUPPORTED_TESTS is a list of benchmarks that go-hello-bench supports
*/
var SUPPORTED_TESTS = []string{
	"filebench",
	"testtest",
}

/*
ParseTestString returns the correct Test for a given string
*/
func ParseTestString(s string) Test {
	switch s {
	case "testtest":
		return TestTest
	case "filebench":
		return Filebench
	}
	log.Fatalf("invalid test name provided")
	return Test{} // so lint is happy. this should never be reached
}

/*
ParseTestStringSlice returns the correct Tests for a given string slice
*/
func ParseTestStringSlice(s []string) []Test {
	toReturn := []Test{}
	for i := range s {
		toReturn = append(toReturn, ParseTestString(s[i]))
	}
	return toReturn
}

/*
CheckDeps will run through the dependencies and return if
they are installed or not.
*/
func (t *Test) CheckDeps() map[string]bool {
	log.Infof("Checking dependencies for %s", t.Name)
	if t.Prereq != nil {
		t.Prereq.CheckDeps()
	}
	toRet := make(map[string]bool)
	for i := range t.Deps {
		val := t.Deps[i].Check()
		toRet[t.Deps[i].Name] = val
	}
	return toRet
}

/*
InstallDepsLocal will go through all the dependencies for a test
and install them if they are not installed.
*/
func (t *Test) InstallDepsLocal() {
	log.Infof("Installing dependencies for %s", t.Name)
	if t.Prereq != nil {
		t.Prereq.InstallDepsLocal()
	}
	for i := range t.Deps {
		val := t.Deps[i].Check()
		if val == false {
			t.Deps[i].InstallLocal()
		}
	}
}
