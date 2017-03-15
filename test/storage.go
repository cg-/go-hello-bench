/*
Contains various storage benchmarks
 */

package test

var Filebench = Test{
   Name: "Filebench",
   Deps: []Dependency{
      Dependency{
         Name:           "build-essential",
         checkCommand:   "dpkg-query -W -f='${Status}' build-essential",
         isInstalledStr: "install ok installed",
         installCommand: "apt install build-essential",
      },
   },
}
