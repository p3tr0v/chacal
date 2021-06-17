package antidebug

import (
	"github.com/p3tr0v/chacal/utils"
	//	"fmt"
	s "strings"

	ps "github.com/mitchellh/go-ps"
)

func processList() bool {
	processList, _ := ps.Processes()

	for x := range processList {
		var process ps.Process
		process = processList[x]
		pn := s.ToLower(process.Executable())

		if utils.PList(utils.ProcessListToCheck, pn) {
			//			fmt.Println("PROCESS_WATCHER:" + pn)
			return true
		}
	}
	return false
}

/**
	Look for suspicious process running
	Example:
	if antidebug.ByProcessWatcher() {
		// exit or wait
	}
**/
func ByProcessWatcher() bool {
	return processList()
}
