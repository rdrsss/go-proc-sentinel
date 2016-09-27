/*
 * @file 	: ProcessSentinel_test.go
 * @brief 	: Tests
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"io/ioutil"
	"log"
	"testing"
)

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

// --------------------------------------------------------------
func TestMain(m *testing.M) {
	m.Run()
}
