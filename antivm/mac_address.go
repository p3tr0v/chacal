package antivm

import (
	"net"

	"github.com/p3tr0v/chacal/utils"
)

var macList = []string{"00:0c:29", "00:50:56", "08:00:27", "52:54:00 ", "00:21:F6", "00:14:4F", "00:0F:4B", "00:10:E0", "00:00:7D", "00:21:28", "00:01:5D", "00:21:F6", "00:A0:A4",
	"00:07:82", "00:03:BA", "08:00:20", "2C:C2:60", "00:10:4F", "00:0F:4B", "00:13:97", "00:20:F2", "00:14:4F"}

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
