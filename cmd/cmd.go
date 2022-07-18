package cmd

import (
	"errors"
	"io"
	"os/exec"

	"github.com/shiningw/aria2go/aria2"
)

type Aria2Cmd struct {
	RunOptions
}

func NewAria2Cmd(opts RunOptions) *Aria2Cmd {
	a := &Aria2Cmd{opts}
	//a.Options = opts.Options
	return a
}
func (a *Aria2Cmd) Run() error {
	cmd := exec.Command("aria2c", a.Options...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	b, _ := io.ReadAll(stderr)
	if len(b) > 0 {
		cmd.Process.Release()
		return errors.New(string(b))
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func StartAria2(opts RunOptions) error {
	a := NewAria2Cmd(opts)
	return a.Run()
}

func StopAria2(a *aria2.Aria2) error {
	return a.Shutdown()
}

func Aria2IsRunning(a *aria2.Aria2) error {

	return a.IsRunning()
}
