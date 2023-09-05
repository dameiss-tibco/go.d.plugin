// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	// "errors"
	// "strings"

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

	// p.once.Do(func() {
	// 	p.adjustCharts(pms)
	// })

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

// func isPulsarHistogram(pm prometheus.SeriesSample) bool {
// 	s := pm.Name()
// 	return strings.HasPrefix(s, "pulsar_storage_write_latency") || strings.HasPrefix(s, "pulsar_entry_size")
// }

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
		metricCPUTimesUser:
		return 1000
	case metricCPUCountCPUCount,
		metricCPUFreqCurrent:
		return 1
	case metricCPUPercentCPUPercent,
		metricDiskUsagePercentUsed:
		return 100
	case metricDiskIOCountersReadBytes,
		metricDiskIOCountersReadCount,
		metricDiskIOCountersReadTime,
		metricDiskIOCountersWriteBytes,
		metricDiskIOCountersWriteCount,
		metricDiskIOCountersWriteTime,
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
		metricDiskUsageFree,
		metricDiskUsageTotal,
		metricDiskUsageUsed:
		return 1000000
	}
	return 1
}
