/*
 * @file 	: Program.go
 * @brief 	: Describes a program, how many processes it runs, it's settings, and metrics.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"os/exec"
)

type Program struct {
	ID         string
	BinPath    string
	Args       []string
	ProcessMap []exec.Cmd
}

// --------------------------------------------------------------
func (p *Program) StartProgram() error {

	return nil
}

// --------------------------------------------------------------
func (p *Program) StopProgram() error {

	return nil
}

// --------------------------------------------------------------
func (p *Program) startProcess() error {

	return nil
}

// --------------------------------------------------------------
func (p *Program) stopProcess() error {

	return nil
}
