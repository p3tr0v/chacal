package antidebug

import (
	"time"
)

/**
	Compare whether the difference between initial and ended time is bigger than time allowed.
	You should get the time in the begging of your function and compare before the end. If the duration would be bigger than alloed, you must be debugging.

	Example:
	func myFuncHere(){
		initTime := time.Now()
		// do your actions here
		if antidebug.ByTimmingDiff(timeInit, 2){ // if your function takes 2 seconds or more, you must be debbuging. Exit your programm.
			exit(0)
		}
	}
**/
func ByTimmingDiff(initTime time.Time, maxTimeAllowed int) bool {
	return (time.Since(initTime) > time.Duration(maxTimeAllowed)*time.Second)

}
