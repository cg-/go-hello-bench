package test

import "github.com/cg-/go-hello-bench/common"

/* Test represents a testing program that can be run.
 */
type Test struct {
   Deps []Dependency
}

/* CheckDeps will run through the dependencies and return if
 * they are installed or not.
 */
func (t *Test) CheckDeps() (error, map[string]bool) {
   toRet := make(map[string]bool)
   for i := range t.Deps{
      val := t.Deps[i].Check()
      toRet[t.Deps[i].Name] = val
   }
}

/* InstallDepsLocal will go through all the dependencies for a test
 * and install them if they are not installed.
 */
func (t *Test) InstallDepsLocal() error {
   for i := range t.Deps{
      val := t.Deps[i].Check()
      if val == false{
         err := t.Deps[i].InstallLocal()
         common.CheckFatal(err)
      }
   }
}


