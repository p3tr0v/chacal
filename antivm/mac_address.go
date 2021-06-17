package antivm

import (
	"net"

	"github.com/p3tr0v/chacal/utils"
)

var macList = []string{"00:0c:29", "00:50:56"}

func getMacAddr() []string {
	ifas, _ := net.Interfaces()

	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as
}

func ByMacAddress() bool {
	return utils.ContainsPrefix(macList, getMacAddr()[0])

}
