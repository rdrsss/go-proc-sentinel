/*
 * @file 	: ProcessSentinel_test.go
 * @brief 	: Tests surrounding the functionality of the Process Sentinel.
 *			  Tests require python 2.7.x, as python scripts are generated.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Define program names here, used to delete later.
var scripts = [...]string{
	"prog0.py",
	"prog1.py"}

/*
 * Create simple py program to test.
 */
var py_prog string = `
import time
if __name__ == '__main__':
	for x in range(0, 10):
		print(x)
		time.sleep(0.05)
`

/*
 * Create a simple python program to test a crash.
 */
var py_crash_prog string = `
import time, ctypes
if __name__ == '__main__':
	# Start Program
	for x in range(0, 30):
		print(x)
		time.sleep(0.05)
	# Crash Program
	i = ctypes.c_char('a')
	j = ctypes.pointer(i)
	c = 0
	while True:
		j[c] = 'a'
		c += 1
	j
`

// Check if a file exists.
func fileExists(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			// We have another error, but, file exists.
			return true, err
		}
	}
	return true, nil
}

// Create slew of test python programs and write them out to files.
func CreatePythonPrograms() error {
	if err := ioutil.WriteFile("prog0.py", []byte(py_prog), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile("prog1.py", []byte(py_crash_prog), 0664); err != nil {
		return err
	}
	return nil
}

// Delete pyhton programs created earlier for test.
func DeletePythonPrograms() {
	for _, v := range scripts {
		if exists, ferr := fileExists(v); exists == true {
			if ferr != nil {
				log.Println(ferr)
			}
			os.Remove(v)
		} else {
			log.Println(ferr)
		}
	}
}

// Quick program struct test with date program, standard to *nix.
func Test_SingleProgramStart(t *testing.T) {
	// Setup simple program
	prog := Program{
		Path: "date",
	}
	// Initialize program
	if init_err := prog.Init(); init_err != nil {
		t.Error(init_err)
	}

	// Pipe stdout
	stdout, stdout_err := prog.Cmd.StdoutPipe()
	if stdout_err != nil {
		t.Error(stdout_err)
	}

	if start_err := prog.Start(); start_err != nil {
		t.Error(start_err)
	}

	buf, buf_err := ioutil.ReadAll(stdout)
	if buf_err != nil {
		t.Error(buf_err)
	}

	log.Println(string(buf))

}

// Test a basic python program.
func Test_PythonProgram(t *testing.T) {
	// Define program.
	prog := Program{
		Path: "python",
		Args: []string{"prog0.py"},
	}
	// Initialize program.
	if init_err := prog.Init(); init_err != nil {
		t.Error(init_err)
	}
	// Pipe stdout and stderr.
	stdout, stdout_err := prog.Cmd.StdoutPipe()
	if stdout_err != nil {
		t.Error(stdout_err)
	}
	// Start running the program.
	if start_err := prog.Start(); start_err != nil {
		t.Error(start_err)
	}

	buf, buf_err := ioutil.ReadAll(stdout)
	if buf_err != nil {
		t.Error(buf_err)
	}

	log.Println(string(buf))
}

// Run script to purposefully crash, testing crash detection logic.
func Test_DetectCrash(t *testing.T) {
	// Define program.
	prog := Program{
		Path: "python",
		Args: []string{"prog1.py"},
	}

	var pm ProcessMonitor
	// Start monitor.
	pm.Initialize()
	// Insert program into monitor.
	// Start program.
	pm.Start()
	// Detect crash.
	// Stop program.
	pm.Stop()
}

// Setup scripts to be run.
func TestMain(m *testing.M) {
	// Create test programs to use with the process sentinel.
	CreatePythonPrograms()
	// Run tests.
	m.Run()
	// Delete test programs.
	DeletePythonPrograms()
}
