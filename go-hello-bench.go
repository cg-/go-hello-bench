package main

import (
	"flag"

	"github.com/op/go-logging"

	"github.com/cg-/go-hello-bench/common"
	"github.com/cg-/go-hello-bench/test"
)

var log *logging.Logger
var tests = []test.Test{}
var operation = ""

// SUPPORTED_OPERATIONS is the operations that our tests support
var SUPPORTED_OPERATIONS = []string{"install", "check"}

// Flags
var verboseFlag = flag.Bool("v", false, "verbose output")
var targetFlag = flag.String("target", "local", "target ('local' or 'docker'")

func init() {
	log = logging.MustGetLogger("go-hello-bench")
	logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	flag.Parse()
	checkArgs()
	if *verboseFlag {
		logging.SetLevel(logging.DEBUG, "go-hello-bench")
	} else {
		logging.SetLevel(logging.CRITICAL, "go-hello-bench")
	}
}

func checkArgs() {
	log.Info("Parsing arguments.")
	argsToCheck := flag.Args()
	for i := range argsToCheck {
		arg := argsToCheck[i]
		log.Infof("Parsing %s.", arg)

		// check if its a test that we support
		if common.IsInStringSlice(arg, test.SUPPORTED_TESTS) {
			tests = append(tests, test.ParseTestString(arg))
			continue

			// check if its a test that we support
		} else if common.IsInStringSlice(arg, SUPPORTED_OPERATIONS) {
			log.Infof("found operation %s", arg)
			if operation != "" {
				log.Fatalf("multiple operations specified (had %s, got %s)", operation, arg)
			}
			operation = arg
			continue

			// otherwise
		} else {
			log.Errorf("No match found: [%s]", arg)
		}
	}

	// state check
	if operation == "" {
		log.Fatalf("no operation specified")
	}
}

func main() {
	log.Info("Starting program.")
	for i := range tests {
		log.Info(tests[i].Name)
	}
}
