package main

import (
	"io"
	"strings"
	"testing"
)

func TestAuthLogin(t *testing.T) {
	in_r, io_w := io.Pipe()
	defer in_r.Close()
	defer io_w.Close()

	cli, o, e := newCLI()
	cli.in = in_r
	args := []string{"auth", "login"}

	go io_w.Write([]byte("\n"))
	s := cli.Run(args)
	if s != ExitCodeError {
		t.Errorf("ExitStatus=%d, want %d", s, ExitCodeError)
	}

	t.Logf("Out=%s Err=%s", o, e)

	expect := "Require: Email"
	if !strings.Contains(e.String(), expect) {
		t.Errorf("Output=%q, want %q", e.String(), expect)
	}
}

// ToDo: Read password test
// func TestAuthLoginPassword(t *testing.T) {
// 	in_r, io_w := io.Pipe()
// 	defer in_r.Close()
// 	defer io_w.Close()

// 	cli, o, e := newCLI()
// 	cli.in = in_r
// 	args := []string{"auth", "login"}

// 	go io_w.Write([]byte("hogehoge\n\n"))
// 	s := cli.Run(args)
// 	if s != ExitCodeError {
// 		t.Errorf("ExitStatus=%d, want %d", s, ExitCodeError)
// 	}

// 	t.Logf("Out=%s Err=%s", o, e)

// 	expect := "Require: Password"
// 	if !strings.Contains(e.String(), expect) {
// 		t.Errorf("Output=%q, want %q", e.String(), expect)
// 	}
// }
