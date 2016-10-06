/*
 * @file 	: ProgramArbiter.go
 * @brief 	: Manages the operation of programs and monitors their states.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/*
 * Process map, holds indexes for quick lookup to Program ptr
 */
type ProcessMap struct {
	pidMap   map[int]*Program  // Pid : Program ptr
	programs map[*Program]bool // Program ptr : Running status
	sync.Mutex
}

/*
 * Process monitor keeps track of running programs
 */
type ProgramArbiter struct {
	running    int32
	processMap ProcessMap
}

// Initialize the process monitor.
func (m *ProgramArbiter) Initialize() {
}

// Start monitoring processes.
func (m *ProgramArbiter) Start() {
	// Set running
	atomic.SwapInt32(&m.running, 1)
	// Spin off arbiter go routine
	go func() {
		for m.IsRunning() {
			// Check if pids are running
			// - return slice of pids that have exited
			// For pids that have exited check settings for restart

			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Stop monitoring processes
func (m *ProgramArbiter) Stop() {
	atomic.SwapInt32(&m.running, 0)
}

// Add a program to the process Monitor
func (m *ProgramArbiter) AddProgram(p *Program) error {
	if p == nil {
		return fmt.Errorf("Passed in nil program")
	}
	// Initialize Program
	if err := p.Init(); err != nil {
		log.Println("Failed to initialize program: ", err)
		return err
	}
	// Start Program
	if err := p.Start(); err != nil {
		log.Println("Failed to start program: ", err)
		return err
	}
	// Validate process has started
	if p.Cmd.Process == nil {
		return fmt.Errorf("Program with nil process")
	}
	// Set fields on program
	p.LastPid = p.Cmd.Process.Pid
	// Insert into Process Map
	m.processMap.AddProgram(p)
	return nil
}

// Check if Arbiter is running.
func (m *ProgramArbiter) IsRunning() bool {
	if atomic.LoadInt32(&m.running) > 0 {
		return true
	}
	return false
}

// Add program to process map
func (pm *ProcessMap) AddProgram(p *Program) {
	pm.Lock()
	pm.programs[p] = true
	pm.pidMap[p.Cmd.Process.Pid] = p
	pm.Unlock()
}

// Remove program from process map
func (pm *ProcessMap) RemoveProgram(p *Program) {
	pm.Lock()
	pm.Unlock()
}

// --------------------------------------------------------------
func checkPID(pid int) (bool, error) {
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
