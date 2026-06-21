package cronjob

import (
	"fmt"
	"forgelog/internal/state"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
)

func CollectStats() {
	var prevBytes uint64

	netStats, _ := psnet.IOCounters(false)
	if len(netStats) > 0 {
		prevBytes = netStats[0].BytesRecv + netStats[0].BytesSent
	}

	for {
		vm, _ := mem.VirtualMemory()
		du, _ := disk.Usage("/")

		cpuPercent, _ := cpu.Percent(0, false)

		netStats, _ := psnet.IOCounters(false)

		var bandwidth uint64

		if len(netStats) > 0 {
			currentBytes := netStats[0].BytesRecv + netStats[0].BytesSent

			if currentBytes >= prevBytes {
				bandwidth = currentBytes - prevBytes
			}

			prevBytes = currentBytes
		}

		state.SystemStats.Mu.Lock()
		state.SystemStats.Stat = state.Stats{
			Memory: fmt.Sprintf(
				"%s/%s",
				humanizeBytes(vm.Used),
				humanizeBytes(vm.Total),
			),
			Disk: fmt.Sprintf(
				"%s/%s",
				humanizeBytes(du.Used),
				humanizeBytes(du.Total),
			),
			CPU: fmt.Sprintf("%.0f%%", cpuPercent[0]),
			Bandwidth: fmt.Sprintf(
				"%s/s",
				humanizeBytes(bandwidth),
			),
		}
		state.SystemStats.Mu.Unlock()

		time.Sleep(2 * time.Second)
	}
}

func humanizeBytes(b uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case b >= GB:
		return fmt.Sprintf("%.1fGB", float64(b)/GB)
	case b >= MB:
		return fmt.Sprintf("%.1fMB", float64(b)/MB)
	case b >= KB:
		return fmt.Sprintf("%.1fKB", float64(b)/KB)
	default:
		return fmt.Sprintf("%dB", b)
	}
}
