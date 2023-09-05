// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

/*
 */

const (
	// Boottime metrics
	metricBoottimeBootTimeSeconds = "system_boottime_boot_time_seconds "
	// CPU Count metrics
	metricCPUCountCPUCount = "system_cpucount_cpu_count"
	// CPU Frequency metrics
	metricCPUFreqCurrent = "system_cpufreq_current"
	// CPU Percent metrics
	metricCPUPercentCPUPercent = "system_cpupercent_cpu_percent"
	// CPU Times metrics
	metricCPUTimesGuest     = "system_cputimes_guest"
	metricCPUTimesGuestNice = "system_cputimes_guest_nice"
	metricCPUTimesIdle      = "system_cputimes_idle"
	metricCPUTimeIOWait     = "system_cputimes_iowait"
	metricCPUTimesIRQ       = "system_cputimes_irq"
	metricCPUTimesNice      = "system_cputimes_nice"
	metricCPUTimesSoftIRQ   = "system_cputimes_softirq"
	metricCPUTimesSteal     = "system_cputimes_steal"
	metricCPUTimesSystem    = "system_cputimes_system"
	metricCPUTimesUser      = "system_cputimes_user"
	// Disk I/O Counters metrics
	metricDiskIOCountersReadBytes  = "system_diskiocounters_read_bytes"
	metricDiskIOCountersReadCount  = "system_diskiocounters_read_count"
	metricDiskIOCountersReadTime   = "system_diskiocounters_read_time"
	metricDiskIOCountersWriteBytes = "system_diskiocounters_write_bytes"
	metricDiskIOCountersWriteCount = "system_diskiocounters_write_count"
	metricDiskIOCountersWriteTime  = "system_diskiocounters_write_time"
	// Disk Usage metrics
	metricDiskUsageFree        = "system_diskusage_free"
	metricDiskUsagePercentUsed = "system_diskusage_percent_used"
	metricDiskUsageTotal       = "system_diskusage_total"
	metricDiskUsageUsed        = "system_diskusage_used"
	// Network I/O Counters metrics
	metricNetworkIOCountersBytesRecv   = "system_networkiocounters_bytes_recv"
	metricNetworkIOCountersBytesSent   = "system_networkiocounters_bytes_sent"
	metricNetworkIOCountersDropIn      = "system_networkiocounters_dropin"
	metricNetworkIOCountersDropOut     = "system_networkiocounters_dropout"
	metricNetworkIOCountersErrIn       = "system_networkiocounters_errin"
	metricNetworkIOCountersErrOut      = "system_networkiocounters_errout"
	metricNetworkIOCountersPacketsRecv = "system_networkiocounters_packets_recv"
	metricNetworkIOCountersPacketsSent = "system_networkiocounters_packets_sent"
	// Swap Memory metrics
	metricSwapMemoryPgFault     = "system_swap_memory_pgfault"
	metricSwapMemoryPgIn        = "system_swap_memory_pgin"
	metricSwapMemoryPgOut       = "system_swap_memory_pgout"
	metricSwapMemorySin         = "system_swap_memory_sin"
	metricSwapMemorySout        = "system_swap_memory_sout"
	metricSwapMemoryFree        = "system_swapmemory_free"
	metricSwapMemoryTotal       = "system_swapmemory_total"
	metricSwapMemoryUsed        = "system_swapmemory_used"
	metricSwapMemoryUsedPercent = "system_swapmemory_used_percent"
	// Virtual Memory metrics
	metricVirtualMemoryActive      = "system_virtualmemory_active"
	metricVirtualMemoryAvailable   = "system_virtualmemory_available"
	metricVirtualMemoryFree        = "system_virtualmemory_free"
	metricVirtualMemoryInactive    = "system_virtualmemory_inactive"
	metricVirtualMemoryTotal       = "system_virtualmemory_total"
	metricVirtualMemoryUsed        = "system_virtualmemory_used"
	metricVirtualMemoryUsedPercent = "system_virtualmemory_used_percent"
	metricVirtualMemoryWired       = "system_virtualmemory_wired"
)
