/*
 * @file 	: ProcessMonitor.go
 * @brief 	: Monitors if processes exist via pid, writes to channel if doesn't
 *			  removes from map.
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

type ProcessMonitor struct {
	pidMap  map[int]chan int
	running int32
	sync.Mutex
}

// --------------------------------------------------------------
func (m *ProcessMonitor) Start() {
	atomic.SwapInt32(&m.running, 1)
	if m.pidMap == nil {
		m.pidMap = make(map[int]chan int)
	}

	go func() {
		for m.IsRunning() {
			m.Lock()
			// TODO ::
			// Check if each process is running
			// If not running
			// 	- write on associated channel
			//	- remove from pid map
			m.Unlock()
			// Sleep for 100 ms
			time.Sleep(100 * time.Millisecond)
		}
	}()

}

// --------------------------------------------------------------
func (m *ProcessMonitor) Stop() {
	atomic.SwapInt32(&m.running, 0)

}

// --------------------------------------------------------------
func (m *ProcessMonitor) AddPid(pid int, sig chan int) {
	m.Lock()
	m.pidMap[pid] = sig
	m.Unlock()
}

// --------------------------------------------------------------
func (m *ProcessMonitor) IsRunning() bool {
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
