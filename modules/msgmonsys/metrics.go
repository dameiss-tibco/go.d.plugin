// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

/*
 */

const (
	// Boottime metrics
	metricBoottimeBootTimeSeconds = "system_boottime_boot_time_seconds"
	// CPU Count metrics
	metricCPUCountCPUCount = "system_cpucount_cpu_count"
	// CPU Frequency metrics
	metricCPUFreqCurrent = "system_cpufreq_current"
	// CPU Percent metrics
	metricCPUPercentCPUPercent = "system_cpupercent_cpu_percent"
	// CPU Times metrics
	metricCPUTimesGuest         = "system_cputimes_guest"
	metricCPUTimesGuestRate     = "system_cputimes_guest_rate"
	metricCPUTimesGuestNice     = "system_cputimes_guest_nice"
	metricCPUTimesGuestNiceRate = "system_cputimes_guest_nice_rate"
	metricCPUTimesIdle          = "system_cputimes_idle"
	metricCPUTimesIdleRate      = "system_cputimes_idle_rate"
	metricCPUTimeIOWait         = "system_cputimes_iowait"
	metricCPUTimeIOWaitRate     = "system_cputimes_iowait_rate"
	metricCPUTimesIRQ           = "system_cputimes_irq"
	metricCPUTimesIRQRate       = "system_cputimes_irq_rate"
	metricCPUTimesNice          = "system_cputimes_nice"
	metricCPUTimesNiceRate      = "system_cputimes_nice_rate"
	metricCPUTimesSoftIRQ       = "system_cputimes_softirq"
	metricCPUTimesSoftIRQRate   = "system_cputimes_softirq_rate"
	metricCPUTimesSteal         = "system_cputimes_steal"
	metricCPUTimesStealRate     = "system_cputimes_steal_rate"
	metricCPUTimesSystem        = "system_cputimes_system"
	metricCPUTimesSystemRate    = "system_cputimes_system_rate"
	metricCPUTimesUser          = "system_cputimes_user"
	metricCPUTimesUserRate      = "system_cputimes_user_rate"
	// Disk I/O Counters metrics
	metricDiskIOCountersReadBytes      = "system_diskiocounters_read_bytes"
	metricDiskIOCountersReadBytesRate  = "system_diskiocounters_read_bytes_rate"
	metricDiskIOCountersReadCount      = "system_diskiocounters_read_count"
	metricDiskIOCountersReadCountRate  = "system_diskiocounters_read_count_rate"
	metricDiskIOCountersReadTime       = "system_diskiocounters_read_time"
	metricDiskIOCountersReadTimeRate   = "system_diskiocounters_read_time_rate"
	metricDiskIOCountersWriteBytes     = "system_diskiocounters_write_bytes"
	metricDiskIOCountersWriteBytesRate = "system_diskiocounters_write_bytes_rate"
	metricDiskIOCountersWriteCount     = "system_diskiocounters_write_count"
	metricDiskIOCountersWriteCountRate = "system_diskiocounters_write_count_rate"
	metricDiskIOCountersWriteTime      = "system_diskiocounters_write_time"
	metricDiskIOCountersWriteTimeRate  = "system_diskiocounters_write_time_rate"
	// Disk Usage metrics
	metricDiskUsageFree        = "system_diskusage_free"
	metricDiskUsagePercentUsed = "system_diskusage_percent_used"
	metricDiskUsageTotal       = "system_diskusage_total"
	metricDiskUsageUsed        = "system_diskusage_used"
	// Network I/O Counters metrics
	metricNetworkIOCountersBytesRecv       = "system_networkiocounters_bytes_recv"
	metricNetworkIOCountersBytesRecvRate   = "system_networkiocounters_bytes_recv_rate"
	metricNetworkIOCountersBytesSent       = "system_networkiocounters_bytes_sent"
	metricNetworkIOCountersBytesSentRate   = "system_networkiocounters_bytes_sent_rate"
	metricNetworkIOCountersDropIn          = "system_networkiocounters_dropin"
	metricNetworkIOCountersDropInRate      = "system_networkiocounters_dropin_rate"
	metricNetworkIOCountersDropOut         = "system_networkiocounters_dropout"
	metricNetworkIOCountersDropOutRate     = "system_networkiocounters_dropout_rate"
	metricNetworkIOCountersErrIn           = "system_networkiocounters_errin"
	metricNetworkIOCountersErrInRate       = "system_networkiocounters_errin_rate"
	metricNetworkIOCountersErrOut          = "system_networkiocounters_errout"
	metricNetworkIOCountersErrOutRate      = "system_networkiocounters_errout_rate"
	metricNetworkIOCountersPacketsRecv     = "system_networkiocounters_packets_recv"
	metricNetworkIOCountersPacketsRecvRate = "system_networkiocounters_packets_recv_rate"
	metricNetworkIOCountersPacketsSent     = "system_networkiocounters_packets_sent"
	metricNetworkIOCountersPacketsSentRate = "system_networkiocounters_packets_sent_rate"
	// Swap Memory metrics
	metricSwapMemoryPgFault     = "system_swap_memory_pgfault"
	metricSwapMemoryPgFaultRate = "system_swap_memory_pgfault_rate"
	metricSwapMemoryPgIn        = "system_swap_memory_pgin"
	metricSwapMemoryPgInRate    = "system_swap_memory_pgin_rate"
	metricSwapMemoryPgOut       = "system_swap_memory_pgout"
	metricSwapMemoryPgOutRate   = "system_swap_memory_pgout_rate"
	metricSwapMemorySin         = "system_swap_memory_sin"
	metricSwapMemorySinRate     = "system_swap_memory_sin_rate"
	metricSwapMemorySout        = "system_swap_memory_sout"
	metricSwapMemorySoutRate    = "system_swap_memory_sout_rate"
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
