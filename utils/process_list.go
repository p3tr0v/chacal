package utils

import (
	"strings"
)

var ProcessListToCheck = []string{"processhacker.exe", "procmon.exe", "pestudio.exe", "procmon64.exe", "x32dbg.exe", "x64dbg.exe", "CFF Explorer.exe",
	"procexp64.exe", "procexp.exe", "pslist.exe", "tcpview.exe", "tcpvcon.exe", "dbgview.exe", "wireshark.exe", "RAMMap.exe", "RAMMap64.exe", "vmmap.exe", "ollydbg.exe", "agent.py"}

var MemoryDumpListToCheck = []string{"DumpIt.exe", "RAMMap.exe", "RAMMap64.exe", "vmmap.exe"}

/**
Check if a program is into an array.
Example:
	s := []string{"James", "Wagner", "Christene", "Mike"}
	fmt.Println(PList(s, "James")) // true
	fmt.Println(PList(s, "Jack")) // false
**/
func PList(s []string, str string) bool {
	for _, v := range s {
		if strings.ToLower(v) == str {
			return true
		}
	}

	return false
}

/**
Check if a program is into an array.
Example:
	s := []string{"00:55:00", "55:55:00", "11:02:BB"}
	fmt.Println(ContainsPrefix(s, "00:55:00:AA:BB:CC")) // true
	fmt.Println(ContainsPrefix(s, "00:BB:00:AA:BB:CC")) // false
**/
func ContainsPrefix(s []string, macAdress string) bool {
	for _, v := range s {
		return strings.HasPrefix(macAdress, v)
	}

	return false
}

/**
Check if a program is into an array.
**/
func ContainsInList(s []string, l string) bool {
	for _, v := range s {

		return strings.Contains(strings.ToLower(l), v)
	}

	return false
}
