package command

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/pkg/errors"
)

type Command struct {
	cmd *exec.Cmd

	waitOnce      sync.Once
	waitError     error
	waitDone      chan struct{}
	sigsWait      sync.WaitGroup
	exitCode      int
	exitCodeValid bool
}

type Params struct {
	Path    string
	Args    []string
	EnvVars []string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
}

func New(p *Params) (*Command, error) {
	cmd := exec.Command(p.Path, p.Args...)
	// setting pgid allows to kill the child process when the parent killed.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmd.Env = append(os.Environ(), p.EnvVars...)

	cmd.Stderr = p.Stderr
	cmd.Stdout = p.Stdout

	if err := cmd.Start(); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Command{
		cmd:      cmd,
		waitDone: make(chan struct{}, 1),
	}, nil
}

func (c *Command) RelaySignals(sigs chan os.Signal) {
	process := c.cmd.Process
	if process == nil || process.Pid <= 0 {
		return
	}

	c.sigsWait.Add(1)
	go func() {
		defer c.sigsWait.Done()

		for {
			select {
			case s, ok := <-sigs:
				if !ok {
					return
				}
				if err := c.cmd.Process.Signal(s); err != nil {
					log.Printf("signalling failed: %s", err.Error())
				}
			case <-c.waitDone:
				return
			}
		}
	}()
}

func (c *Command) Sigkill() error {
	process := c.cmd.Process
	if process == nil || process.Pid <= 0 {
		return nil
	}

	if err := syscall.Kill(-process.Pid, syscall.SIGTERM); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *Command) Wait() error {
	c.waitOnce.Do(c.wait)

	return c.waitError
}

func (c *Command) wait() {
	defer close(c.waitDone)
	c.waitError = c.cmd.Wait()
	if c.waitError == nil {
		c.exitCode = 0
		c.exitCodeValid = true
		return
	}

	exitError, ok := c.waitError.(*exec.ExitError)
	if !ok {
		return
	}

	c.exitCode = exitError.ExitCode()
	c.exitCodeValid = true
}

func (c *Command) ExitCode() (int, bool) {
	return c.exitCode, c.exitCodeValid
}
