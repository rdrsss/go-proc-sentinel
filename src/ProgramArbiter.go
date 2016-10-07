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

func (pm *ProcessMap) Init() {
	pm.pidMap = make(map[int]*Program)
	pm.programs = make(map[*Program]bool)
}

/*
 * Program arbiter periodically checks to make sure
 * all programs are running. Restarting them if they are
 * configured as such.
 */
type ProgramArbiter struct {
	running    int32
	processMap ProcessMap
}

// Check with the os if a pid is running
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

// Initialize the process monitor.
func (a *ProgramArbiter) Init() {
	a.processMap.Init()
}

// Start monitoring processes.
func (a *ProgramArbiter) Start() {
	// Set running
	atomic.SwapInt32(&a.running, 1)
	// Spin off arbiter go routine
	go func() {
		for a.IsRunning() {
			// Check if pids are running
			// - return slice of pids that have exited
			// For pids that have exited check settings for restart

			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Stop monitoring processes
func (a *ProgramArbiter) Stop() {
	// Set running
	atomic.SwapInt32(&a.running, 0)
}

// Add a program to the process Monitor.
// Makes a copy of a program, then initializes it internally and from here on out
// a ptr is used.
func (a *ProgramArbiter) AddProgram(p Program) error {
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
	p.State = ProgramRunning
	// Insert into Process Map
	a.processMap.AddProgram(&p)

	log.Printf("Started : program[%s] pid[%d]", p.Path, p.LastPid)
	return nil
}

// Check if Arbiter is running.
func (a *ProgramArbiter) IsRunning() bool {
	if atomic.LoadInt32(&a.running) > 0 {
		return true
	}
	return false
}

// Check if pids still exist, return a list of pids that don't.
func (a *ProgramArbiter) checkPids() []int {
	var pids []int

	return pids
}

// Add program to process map
func (pm *ProcessMap) AddProgram(p *Program) {
	pm.Lock()
	pm.programs[p] = true
	pm.pidMap[p.LastPid] = p
	pm.Unlock()
}

// Remove program from process map
func (pm *ProcessMap) RemoveProgram(p *Program) {
	pm.Lock()
	if _, ok := pm.programs[p]; ok {
		delete(pm.programs, p)
	}

	if _, ok := pm.pidMap[p.LastPid]; ok {
		delete(pm.pidMap, p.LastPid)
	}
	pm.Unlock()
}
