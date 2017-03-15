package main

import (
   "github.com/op/go-logging"

   "github.com/cg-/go-hello-bench/test"
   "fmt"
   "os"
)

var log *logging.Logger

func init() {
   log = logging.MustGetLogger("go-hello-bench")
   logging.MustStringFormatter(
      `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
   )
   logging.SetBackend(logging.NewLogBackend(os.Stderr, "", 0))
}

func main() {
   log.Info("Starting program.")

   // Check if superuser.
   /*
   if os.Getegid() != 0 {
      log.Fatal("Please run this as superuser.")
   }
   */

   m := test.Filebench.CheckDeps()
   for k := range m{
      fmt.Printf("%s: %s", k, m[k])
   }
}
