package antivm

import (
	"github.com/p3tr0v/chacal/utils"

	"github.com/StackExchange/wmi"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Win32_DiskDrive struct {
	PNPDeviceID string
	Size        uint64
}

// wmic diskdrive list full
func queryWMI() []Win32_DiskDrive {
	var diskdrive []Win32_DiskDrive
	q := wmi.CreateQuery(&diskdrive, "")
	wmi.Query(q, &diskdrive)
	return diskdrive
}

func diskTotalSize(maxSize uint64) bool {
	diskdrive := queryWMI()

	for _, diskInfo := range diskdrive {
		return (diskInfo.Size / GB) < maxSize
	}
	return false
}

/**
	AntiVM by total disk size in GB
	Example:
	if antivm.BySizeDisk(100) { // whether total disk size is less than 100 GB, exit the program
		// exit or wait
	}
**/
func BySizeDisk(maxSize uint64) bool {
	return diskTotalSize(maxSize)
}

var deviceIDList = []string{"virtual", "vmware", "vbox"}

/**
	Check whether may be a virtual disk
	Example:
	if antivm.IsVirtualDisk() {
		// exit or wait
	}
**/
func IsVirtualDisk() bool {
	diskdrive := queryWMI()

	for _, diskInfo := range diskdrive {
		return utils.ContainsInList(deviceIDList, diskInfo.PNPDeviceID)
		/*return strings.Contains(strings.ToLower(diskInfo.PNPDeviceID), "virtual") ||
		strings.Contains(strings.ToLower(diskInfo.PNPDeviceID), "vmware") ||
		strings.Contains(strings.ToLower(diskInfo.PNPDeviceID), "vbox")*/
	}
	return false
}
