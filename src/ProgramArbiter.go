/*
 * @file 	: ProgramArbiter.go
 * @brief 	: Manages the operation of programs and monitors their states.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"log"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

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
	//sync.Mutex
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
			// TODO ::
			// Check if each process is running
			// If not running
			// 	- write on associated channel
			//	- remove from pid map
			// Sleep for 100 ms
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Stop monitoring processes
func (m *ProgramArbiter) Stop() {
	atomic.SwapInt32(&m.running, 0)
}

// Add a program to the process Monitor
func (m *ProgramArbiter) AddProgram(p *Program) {

}

// Check if Arbiter is running.
func (m *ProgramArbiter) IsRunning() bool {
	if atomic.LoadInt32(&m.running) > 0 {
		return true
	}
	return false
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
