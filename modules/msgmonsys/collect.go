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

		value := pm.Value * precision(pm.Name())
		mx[pm.Name()+"_"+sys.name] += value

		p.curCache.systems[sys] = true
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
		metricCPUTimesGuestNice,
		metricCPUTimesIdle,
		metricCPUTimeIOWait,
		metricCPUTimesIRQ,
		metricCPUTimesNice,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSteal,
		metricCPUTimesSystem,
		metricCPUTimesUser,
		metricDiskIOCountersReadBytes,
		metricDiskIOCountersReadCount,
		metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteBytes,
		metricDiskIOCountersWriteCount,
		metricDiskIOCountersWriteTime,
		metricDiskUsageFree,
		metricDiskUsagePercentUsed,
		metricDiskUsageTotal,
		metricDiskUsageUsed,
		metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesSent,
		metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsSent,
		metricSwapMemoryPgFault,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgOut,
		metricSwapMemorySin,
		metricSwapMemorySout,
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
	case metricBoottimeBootTimeSeconds,
		metricCPUTimesGuest,
		metricCPUTimesGuestNice,
		metricCPUTimesIdle,
		metricCPUTimeIOWait,
		metricCPUTimesIRQ,
		metricCPUTimesNice,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSteal,
		metricCPUTimesSystem,
		metricCPUTimesUser,
		metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteTime:
		// Metric is in seconds
		return 1
	case metricCPUCountCPUCount:
		// Metric is a scalar
		return 1
	case metricCPUFreqCurrent:
		// Metric is in Mhz, convert to Ghz
		return 1000
	case metricCPUPercentCPUPercent,
		metricDiskUsagePercentUsed,
		metricSwapMemoryUsedPercent,
		metricVirtualMemoryUsedPercent:
		// Metric is a percent scaled up by 100, so divide by 100
		return 100
	case metricDiskIOCountersReadBytes,
		metricDiskIOCountersWriteBytes,
		metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesSent:
		// Metric is number of bytes, scale to KiB
		return 1000
	case metricDiskIOCountersReadCount,
		metricDiskIOCountersWriteCount,
		metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsSent,
		metricSwapMemoryPgFault,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgOut,
		metricSwapMemorySin,
		metricSwapMemorySout:
		// Metric is number of events
		return 1
	case metricSwapMemoryFree,
		metricSwapMemoryTotal,
		metricSwapMemoryUsed,
		metricVirtualMemoryActive,
		metricVirtualMemoryAvailable,
		metricVirtualMemoryFree,
		metricVirtualMemoryInactive,
		metricVirtualMemoryTotal,
		metricVirtualMemoryUsed,
		metricVirtualMemoryWired:
		// Metric is number of memory bytes, scale to MiB
		return 1024 * 1024
	case metricDiskUsageFree,
		metricDiskUsageTotal,
		metricDiskUsageUsed:
		// Metric is number of disk bytes, scale to GB
		return 1000 * 1000 * 1000
	default:
		return 1
	}
}

func precision(metric string) float64 {
	switch metric {
	case metricBoottimeBootTimeSeconds,
		metricCPUTimesGuest,
		metricCPUTimesGuestNice,
		metricCPUTimesIdle,
		metricCPUTimeIOWait,
		metricCPUTimesIRQ,
		metricCPUTimesNice,
		metricCPUTimesSoftIRQ,
		metricCPUTimesSteal,
		metricCPUTimesSystem,
		metricCPUTimesUser,
		metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteTime:
		// Metric is in seconds
		return 1
	case metricCPUCountCPUCount:
		// Metric is a scalar
		return 1
	case metricCPUFreqCurrent:
		// Metric is in Mhz, convert to Ghz
		return 1000
	case metricCPUPercentCPUPercent,
		metricDiskUsagePercentUsed,
		metricSwapMemoryUsedPercent,
		metricVirtualMemoryUsedPercent:
		// Metric is a percent, scale up by 100
		return 100
	case metricDiskIOCountersReadBytes,
		metricDiskIOCountersWriteBytes,
		metricNetworkIOCountersBytesRecv,
		metricNetworkIOCountersBytesSent:
		// Metric is number of bytes, scale to KiB
		return 1
	case metricDiskIOCountersReadCount,
		metricDiskIOCountersWriteCount,
		metricNetworkIOCountersDropIn,
		metricNetworkIOCountersDropOut,
		metricNetworkIOCountersErrIn,
		metricNetworkIOCountersErrOut,
		metricNetworkIOCountersPacketsRecv,
		metricNetworkIOCountersPacketsSent,
		metricSwapMemoryPgFault,
		metricSwapMemoryPgIn,
		metricSwapMemoryPgOut,
		metricSwapMemorySin,
		metricSwapMemorySout:
		// Metric is number of events
		return 1
	case metricSwapMemoryFree,
		metricSwapMemoryTotal,
		metricSwapMemoryUsed,
		metricVirtualMemoryActive,
		metricVirtualMemoryAvailable,
		metricVirtualMemoryFree,
		metricVirtualMemoryInactive,
		metricVirtualMemoryTotal,
		metricVirtualMemoryUsed,
		metricVirtualMemoryWired:
		// Metric is number of memory bytes
		return 1
	case metricDiskUsageFree,
		metricDiskUsageTotal,
		metricDiskUsageUsed:
		// Metric is number of disk bytes
		return 1
	default:
		return 1
	}
}
