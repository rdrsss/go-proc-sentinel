/*
 * @file 	: ProcessSentinel_test.go
 * @brief 	: Tests
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"testing"
)

// --------------------------------------------------------------
func Test_SingleProgramStart(t *testing.T) {
	// Setup simple program
	prog := Program{
		Path: "date",
	}

	if init_err := prog.InitProgram(); init_err != nil {
		t.Error(init_err)
	}

	if start_err := prog.StartProgram(); start_err != nil {
		t.Error(start_err)
	}

	// Pipe stdout

}

// --------------------------------------------------------------
func TestMain(m *testing.M) {
	m.Run()
}
