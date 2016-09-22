/*
 * @file 	: Program.go
 * @brief 	: Describes a program, how many processes it runs, it's settings, and metrics.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

type Program struct {
	ID   string
	Path string
	Args []string

	cmd       *exec.Cmd
	startTime time.Time
}

func (p *Program) InitProgram() error {
	// Validate Path
	path, err := exec.LookPath(p.Path)
	if err != nil {
		return err
	}
	// Setup Comamnd
	p.cmd = exec.Command(path, p.Args...)
	// Append to command map
	//p. = append(p.CmdMap, cmd)

	return nil
}

// --------------------------------------------------------------
func (p *Program) StartProgram() error {
	go func() {
		p.cmd.Wait()
		fmt.Println("Program exited ")
	}()

	if err := p.cmd.Start(); err != nil {
		return err
	}
	/*
		for _, cmd := range p.CmdMap {
			cmd.Start()
		}
	*/
	return nil
}

// --------------------------------------------------------------
func (p *Program) StopProgram() error {
	return nil
}

// --------------------------------------------------------------
func (p *Program) IsProgramRunning() (bool, error) {
	// Get Program pid
	p.checkPID(p.cmd.Process.Pid)

	return true, nil
}

// --------------------------------------------------------------
func (p Program) checkPID(pid int) (bool, error) {
	// *nix only
	buf, err := exec.Command("kill", "-s", "0", strconv.Itoa(pid)).CombinedOutput()
	if err != nil {
		log.Println(err)
		return false, err
	}

	if string(buf) != "" {
		log.Println(string(buf))
	}

	return true, nil
}
