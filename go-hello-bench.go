package main

import (
   "github.com/cg-/go-hello-bench/common"
   "github.com/cg-/go-hello-bench/test"
   "fmt"
)

func main() {
   common.Logger.Printf("Starting program.")
   m := test.Filebench.CheckDeps()
   for k := range m{
      fmt.Printf("%s: %s", k, m[k])
   }
}
