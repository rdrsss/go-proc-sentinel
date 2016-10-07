/*
 * @file 	: ProcessSentinel.go
 * @brief 	: Entry point for the service.
 * @author 	: Manuel A. Rodriguez (manuel.rdrs@gmail.com)
 */
package main

import (
	"log"
)

var (
	pa ProgramArbiter
)

// --------------------------------------------------------------
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pa.Init()
	pa.Start()
	// Main loop
	pa.Stop()
}
