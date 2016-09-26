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
	pm ProcessMonitor
)

// --------------------------------------------------------------
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pm.Initialize()

}
