/*
 * @file 	: ProcessSentinel_test.go
 * @brief 	: Tests
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

/*
 * Create simple py program to test
 */
var py_prog string = `
import time
if __name__ == '__main__':
	for x in range(0, 10):
		print(x)
		time.sleep(0.05)
`

func CreatePyProgram() error {
	return ioutil.WriteFile("prog0.py", []byte(py_prog), 0644)
}

func DeletePyProgram() error {
	return os.Remove("prog0.py")

}

// --------------------------------------------------------------
func Test_SingleProgramStart(t *testing.T) {
	// Setup simple program
	prog := Program{
		Path: "date",
	}
	// Initialize program
	if init_err := prog.InitProgram(); init_err != nil {
		t.Error(init_err)
	}

	// Pipe stdout
	stdout, stdout_err := prog.Cmd.StdoutPipe()
	if stdout_err != nil {
		t.Error(stdout_err)
	}

	if start_err := prog.StartProgram(); start_err != nil {
		t.Error(start_err)
	}

	buf, buf_err := ioutil.ReadAll(stdout)
	if buf_err != nil {
		t.Error(buf_err)
	}

	log.Println(string(buf))

}

func Test_PythonProgram(t *testing.T) {
	// Create Progra
	if create_err := CreatePyProgram(); create_err != nil {
		t.Error(create_err)
	}
	// Run Program

	// Delete Program
	if delete_err := DeletePyProgram(); delete_err != nil {
		t.Error(delete_err)
	}
}

// --------------------------------------------------------------
func TestMain(m *testing.M) {
	m.Run()
}
