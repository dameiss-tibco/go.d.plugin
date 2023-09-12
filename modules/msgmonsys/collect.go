// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	"github.com/netdata/go.d.plugin/pkg/prometheus"
	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (p *MsgmonSys) resetCurCache() {
	for sys := range p.curCache.systems {
		delete(p.curCache.systems, sys)
	}
}

func (p *MsgmonSys) collect() (map[string]int64, error) {
	pms, err := p.prom.ScrapeSeries()
	if err != nil {
		return nil, err
	}

	mx := p.collectMetrics(pms)
	p.updateCharts()
	p.resetCurCache()

	return stm.ToMap(mx), nil
}

func (p *MsgmonSys) collectMetrics(pms prometheus.Series) map[string]float64 {
	mx := make(map[string]float64)
	pms = findMsgmonSysMetrics(pms)
	for _, pm := range pms {
		sys := newSystemName(pm)
		if sys.name == "" {
			continue
		}

		value := pm.Value * multiplier(pm.Name())
		mx[pm.Name()+"_"+sys.name] += value

		p.curCache.systems[sys] = true
		mx["cpupct_rate_0_"+sys.name] = 0
		mx["cpupct_rate_100_"+sys.name] = 100
	}
	return mx
}

func newSystemName(pm prometheus.SeriesSample) systemName {
	return systemName{
		name: pm.Labels.Get("systemName"),
	}
}

func findMsgmonSysMetrics(pms prometheus.Series) prometheus.Series {
	var ms prometheus.Series
	pms = pms.FindByNames(
		metricBoottimeBootTimeSeconds,
		metricCPUCountCPUCount,
		metricCPUFreqCurrent,
		metricCPUPercentCPUPercent,
		metricCPUTimesGuest,
		metricCPUTimesGuestRate,
		metricCPUTimesGuestNice,
		metricCPUTimesGuestNiceRate,
		metricCPUTimesIdle,
		metricCPUTimesIdleRate,
		metricCPUTimeIOWait,
		metricCPUTimeIOWaitRate,
		metricCPUTimesIRQ,
		metricCPUTimesIRQRate,
		metricCPUTimesNice,
		metricCPUTimesNiceRate,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSoftIRQRate,
		metricCPUTimesSteal,
		metricCPUTimesStealRate,
		metricCPUTimesSystem,
		metricCPUTimesSystemRate,
		metricCPUTimesUser,
		metricCPUTimesUserRate,
		metricDiskIOCountersReadBytes,
		metricDiskIOCountersReadBytesRate,
		metricDiskIOCountersReadCount,
		metricDiskIOCountersReadCountRate,
		metricDiskIOCountersReadTime,
		metricDiskIOCountersReadTimeRate,
		metricDiskIOCountersWriteBytes,
		metricDiskIOCountersWriteBytesRate,
		metricDiskIOCountersWriteCount,
		metricDiskIOCountersWriteCountRate,
		metricDiskIOCountersWriteTime,
		metricDiskIOCountersWriteTimeRate,
		metricDiskUsageFree,
		metricDiskUsagePercentUsed,
		metricDiskUsageTotal,
		metricDiskUsageUsed,
		metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesRecvRate,
		metricNetworkIOCountersBytesSent,
		metricNetworkIOCountersBytesSentRate,
		metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropInRate,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersDropOutRate,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrInRate,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersErrOutRate,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsRecvRate,
		metricNetworkIOCountersPacketsSent,
		metricNetworkIOCountersPacketsSentRate,
		metricSwapMemoryPgFault,
		metricSwapMemoryPgFaultRate,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgInRate,
		metricSwapMemoryPgOut,
		metricSwapMemoryPgOutRate,
		metricSwapMemorySin,
		metricSwapMemorySinRate,
		metricSwapMemorySout,
		metricSwapMemorySoutRate,
		metricSwapMemoryFree,
		metricSwapMemoryTotal,
		metricSwapMemoryUsed,
		metricSwapMemoryUsedPercent,
		metricVirtualMemoryActive,
		metricVirtualMemoryAvailable,
		metricVirtualMemoryFree,
		metricVirtualMemoryInactive,
		metricVirtualMemoryTotal,
		metricVirtualMemoryUsed,
		metricVirtualMemoryUsedPercent,
		metricVirtualMemoryWired,
	)
	return append(ms, pms...)
}

func scale(metric string) int {
	switch metric {
	case metricBoottimeBootTimeSeconds:
		// Metric is in integer seconds
		return 1
	case metricCPUCountCPUCount:
		// Metric is a scalar
		return 1
	case metricCPUFreqCurrent:
		// Metric is in Mhz, convert to Ghz
		return 1000
	case metricCPUPercentCPUPercent:
		// Metric is a floating-point percentage multiplied by 100, scale back to percentage
		return 100
	case metricCPUTimesGuest,
		metricCPUTimesGuestNice,
		metricCPUTimesIdle,
		metricCPUTimeIOWait,
		metricCPUTimesIRQ,
		metricCPUTimesNice,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSteal,
		metricCPUTimesSystem,
		metricCPUTimesUser:
		// Metric is in seconds multiplied by 100, scale back to original value
		return 100
	case metricCPUTimesGuestRate,
		metricCPUTimesGuestNiceRate,
		metricCPUTimesIdleRate,
		metricCPUTimeIOWaitRate,
		metricCPUTimesIRQRate,
		metricCPUTimesNiceRate,
		metricCPUTimesSoftIRQRate,
		metricCPUTimesStealRate,
		metricCPUTimesSystemRate,
		metricCPUTimesUserRate:
		// Metric is in milliseconds per second, convert to seconds per second
		return 1000
	case metricDiskIOCountersReadBytes,
		metricDiskIOCountersWriteBytes:
		// Metric is in bytes
		return 1
	case metricDiskIOCountersReadCount,
		metricDiskIOCountersWriteCount:
		// Metric is in operations
		return 1
	case metricDiskIOCountersReadBytesRate,
		metricDiskIOCountersWriteBytesRate:
		// Metric is in bytes per second multiplied by 100 for additional precision, scale back to original
		return 100
	case metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteTime:
		// Metric is in milliseconds per second
		return 1
	case metricDiskIOCountersReadCountRate,
		metricDiskIOCountersWriteCountRate:
		// Metric is in operations per second multiplied by 100 for additional precision, scale back to original
		return 100
	case metricDiskIOCountersReadTimeRate,
		metricDiskIOCountersWriteTimeRate:
		// Metric is in milliseconds per second
		return 1
	case metricDiskUsageFree,
		metricDiskUsageTotal,
		metricDiskUsageUsed:
		// Metric is in bytes, scale to MB
		return 1000 * 1000
	case metricDiskUsagePercentUsed:
		// Metric is a floating-point percentage multiplied by 100 to get additional precision, scale back to percentage
		return 100
	case metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesSent:
		// Metric is in bytes, scale to kB
		return 1000
	case metricNetworkIOCountersBytesRecvRate,
		metricNetworkIOCountersBytesSentRate:
		// Metric is in bytes per second multiplied by 100 for additional precision, scale back to original
		return 100
	case metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsSent:
		// Metric is in packets
		return 1
	case metricNetworkIOCountersDropInRate,
		metricNetworkIOCountersDropOutRate,
		metricNetworkIOCountersErrInRate,
		metricNetworkIOCountersErrOutRate,
		metricNetworkIOCountersPacketsRecvRate,
		metricNetworkIOCountersPacketsSentRate:
		// Metric is in packets per second multiplied by 100 for additional precision, scale back to original
		return 100
	case metricSwapMemoryPgFault,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgOut,
		metricSwapMemorySin,
		metricSwapMemorySout:
		// Metric is in operations
		return 1
	case metricSwapMemoryPgFaultRate,
		metricSwapMemoryPgInRate,
		metricSwapMemoryPgOutRate,
		metricSwapMemorySinRate,
		metricSwapMemorySoutRate:
		// Metric is in operations per second multiplied by 100 for additional precision, scale back to original
		return 100
	case metricSwapMemoryFree,
		metricSwapMemoryTotal,
		metricSwapMemoryUsed:
		// Metric is in bytes
		return 1
	case metricSwapMemoryUsedPercent:
		// Metric is a percentage multiplied by 100 for additional precision, scale back to original
		return 100
	case metricVirtualMemoryActive,
		metricVirtualMemoryAvailable,
		metricVirtualMemoryFree,
		metricVirtualMemoryInactive,
		metricVirtualMemoryTotal,
		metricVirtualMemoryUsed,
		metricVirtualMemoryWired:
		// Metric is in bytes
		return 1
	case metricVirtualMemoryUsedPercent:
		// Metric is a percentage multiplied by 100 for additional precision, scale back to original
		return 100
	default:
		return 1
	}
}

func multiplier(metric string) float64 {
	switch metric {
	case metricBoottimeBootTimeSeconds:
		// Metric is in integer seconds
		return 1
	case metricCPUCountCPUCount:
		// Metric is a scalar
		return 1
	case metricCPUFreqCurrent:
		// Metric is in Mhz
		return 1
	case metricCPUPercentCPUPercent:
		// Metric is a floating-point percentage, multiply by 100 to get additional precision
		return 100
	case metricCPUTimesGuest,
		metricCPUTimesGuestNice,
		metricCPUTimesIdle,
		metricCPUTimeIOWait,
		metricCPUTimesIRQ,
		metricCPUTimesNice,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSteal,
		metricCPUTimesSystem,
		metricCPUTimesUser:
		// Metric is in seconds, multiply by 100 to get additional precision
		return 100
	case metricCPUTimesGuestRate,
		metricCPUTimesGuestNiceRate,
		metricCPUTimesIdleRate,
		metricCPUTimeIOWaitRate,
		metricCPUTimesIRQRate,
		metricCPUTimesNiceRate,
		metricCPUTimesSoftIRQRate,
		metricCPUTimesStealRate,
		metricCPUTimesSystemRate,
		metricCPUTimesUserRate:
		// Metric is in seconds per second, multiply by 1000 to get milliseconds per second
		return 1000
	case metricDiskIOCountersReadBytes,
		metricDiskIOCountersWriteBytes:
		// Metric is in bytes
		return 1
	case metricDiskIOCountersReadCount,
		metricDiskIOCountersWriteCount:
		// Metric is in operations
		return 1
	case metricDiskIOCountersReadBytesRate,
		metricDiskIOCountersWriteBytesRate:
		// Metric is in bytes per second, multiply by 100 for additional precision
		return 100
	case metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteTime:
		// Metric is in seconds per second, multiply by 1000 to get milliseconds per second
		return 1000
	case metricDiskIOCountersReadCountRate,
		metricDiskIOCountersWriteCountRate:
		// Metric is in operations per second, multiply by 100 for additional precision
		return 100
	case metricDiskIOCountersReadTimeRate,
		metricDiskIOCountersWriteTimeRate:
		// Metric is in seconds per second, multiply by 1000 to get milliseconds per second
		return 1000
	case metricDiskUsageFree,
		metricDiskUsageTotal,
		metricDiskUsageUsed:
		// Metric is in bytes
		return 1
	case metricDiskUsagePercentUsed:
		// Metric is a floating-point percentage, multiply by 100 to get additional precision
		return 100
	case metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesSent:
		// Metric is in bytes
		return 1
	case metricNetworkIOCountersBytesRecvRate,
		metricNetworkIOCountersBytesSentRate:
		// Metric is in bytes per second, multiply by 100 for additional precision
		return 100
	case metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsSent:
		// Metric is in packets
		return 1
	case metricNetworkIOCountersDropInRate,
		metricNetworkIOCountersDropOutRate,
		metricNetworkIOCountersErrInRate,
		metricNetworkIOCountersErrOutRate,
		metricNetworkIOCountersPacketsRecvRate,
		metricNetworkIOCountersPacketsSentRate:
		// Metric is in packets per second, multiply by 100 for additional precision
		return 100
	case metricSwapMemoryPgFault,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgOut,
		metricSwapMemorySin,
		metricSwapMemorySout:
		// Metric is in operations
		return 1
	case metricSwapMemoryPgFaultRate,
		metricSwapMemoryPgInRate,
		metricSwapMemoryPgOutRate,
		metricSwapMemorySinRate,
		metricSwapMemorySoutRate:
		// Metric is in operations per second, multiply by 100 for additional precision
		return 100
	case metricSwapMemoryFree,
		metricSwapMemoryTotal,
		metricSwapMemoryUsed:
		// Metric is in bytes
		return 1
	case metricSwapMemoryUsedPercent:
		// Metric is a percentage, multiply by 100 for additional precision
		return 100
	case metricVirtualMemoryActive,
		metricVirtualMemoryAvailable,
		metricVirtualMemoryFree,
		metricVirtualMemoryInactive,
		metricVirtualMemoryTotal,
		metricVirtualMemoryUsed,
		metricVirtualMemoryWired:
		// Metric is in bytes
		return 1
	case metricVirtualMemoryUsedPercent:
		// Metric is a percentage, multiply by 100 for additional precision
		return 100
	default:
		return 1
	}
}
