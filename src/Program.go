/*
 * @file 	: Program.go
 * @brief 	: Describes a program, how many processes it runs, it's settings, and metrics.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

type Program struct {
	ID   string
	Path string
	Args []string

	Cmd       *exec.Cmd
	startTime time.Time
}

// --------------------------------------------------------------
func (p *Program) InitProgram() error {
	// Validate Path
	path, err := exec.LookPath(p.Path)
	if err != nil {
		return err
	}
	// Setup Comamnd
	p.Cmd = exec.Command(path, p.Args...)
	// Set to different process group than parent.
	pattr := syscall.SysProcAttr{
		Setpgid: true,
	}
	p.Cmd.SysProcAttr = &pattr

	return nil
}

// --------------------------------------------------------------
func (p *Program) StartProgram() error {
	go func() {
		p.Cmd.Wait()
		fmt.Println("Program exited ")
	}()

	if err := p.Cmd.Start(); err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------
func (p *Program) StopProgram() error {
	return nil
}
