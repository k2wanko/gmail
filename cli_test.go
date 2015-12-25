package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	cli, o, e := newCLI()
	args := []string{"--version"}

	s := cli.Run(args)
	if s != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", s, ExitCodeOK)
	}

	t.Logf("Out=%s Err=%s", o, e)

	if !strings.Contains(e.String(), Version) {
		t.Errorf("Output=%q, want %q", e.String(), Version)
	}
}

func newCLI() (*CLI, *bytes.Buffer, *bytes.Buffer) {
	o, e := new(bytes.Buffer), new(bytes.Buffer)
	return &CLI{out: o, err: e}, o, e
}
