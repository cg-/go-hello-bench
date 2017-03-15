package test


/* Test represents a testing program that can be run.
 */
type Test struct {
   Name string
   Deps []Dependency
}

/* CheckDeps will run through the dependencies and return if
 * they are installed or not.
 */
func (t *Test) CheckDeps() map[string]bool {
   log.Infof("Checking dependencies for %s", t.Name)
   toRet := make(map[string]bool)
   for i := range t.Deps{
      val := t.Deps[i].Check()
      toRet[t.Deps[i].Name] = val
   }
   return toRet
}

/* InstallDepsLocal will go through all the dependencies for a test
 * and install them if they are not installed.
 */
func (t *Test) InstallDepsLocal() {
   log.Infof("Installing dependencies for %s", t.Name)
   for i := range t.Deps{
      val := t.Deps[i].Check()
      if val == false{
         t.Deps[i].InstallLocal()
      }
   }
}


