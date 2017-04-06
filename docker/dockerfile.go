package docker

import logging "github.com/op/go-logging"
import "fmt"

var log = logging.MustGetLogger("go-hello-bench")

/*
Dockerfile represents something that will create a Dockerfile to run a
container.
*/
type Dockerfile struct {
	From       string
	Maintainer string
	Tests      []Dockertest
}

/*
Dockertest represents a test that should be set up within a Dockerfile
*/
type Dockertest struct {
	Commands []string
}

func (d *Dockertest) String() {
	out := ""
	for i := range d.Commands {
		out = fmt.Sprintf("%s\n%s", out, d.Commands[i])
	}
}

/*
WriteOut writes the Dockerfile to the given location
*/
func (d *Dockerfile) WriteOut(path string) error {
	return nil
}
