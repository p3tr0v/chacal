package antimem

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

		if utils.PList(utils.MemoryDumpListToCheck, pn) {
			//	fmt.Println("MEM_WATCHER:" + pn)
			return true
		}
	}
	return false
}

/**
	Look for programs that are able to inspect/dump memory.
	Example:
		if antimem.ByMemWatcher() {
		// exit or wait
	}
**/
func ByMemWatcher() bool {
	return processList()
}
