package main

import (
	"flag"
	"github.com/google/go-cmdtest"
	"testing"
)

var update = flag.Bool("update", true, "replace test file contents with output")

func Test(t *testing.T) {
	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}
	ts.Commands["docker"] = cmdtest.Program("/usr/bin/docker")
	ts.Run(t, *update)
}

func TestDocker(t *testing.T) {
	ts, err := cmdtest.Read("docker19")
	if err != nil {
		t.Fatal(err)
	}
	ts.Commands["docker"] = cmdtest.Program("/usr/bin/docker")
	ts.Run(t, *update)
}

func TestKonvoy2(t *testing.T) {
	ts, err := cmdtest.Read("konvoy2-doctests")
	if err != nil {
		t.Fatal(err)
	}
	ts.Commands["konvoy2"] = cmdtest.Program("./konvoy2")
	ts.Run(t, *update)
}
