// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	"fmt"
	"github.com/netdata/go.d.plugin/agent/module"
)

type (
	Charts = module.Charts
	Chart  = module.Chart
	Dims   = module.Dims
	Dim    = module.Dim
	Opts   = module.Opts
)

var summaryCharts = Charts{
	summaryChart.Copy(),
}

var (
	summaryChart = Chart{
		ID:    "summary",
		Title: "CPU Percent Summary",
		Units: "percentage",
		Fam:   "CPU Percent",
		Ctx:   "msgmonsys.cpu_percent",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
	}
)

var systemCharts = Charts{
	cpuTimesChart.Copy(),
	cpuPercentChart.Copy(),
	diskUsageChart.Copy(),
	virtualMemoryChart.Copy(),
	diskIODataChart.Copy(),
	diskIOOperationChart.Copy(),
	networkIODataChart.Copy(),
	networkIOOperationChart.Copy(),
}

var (
	cpuTimesChart = Chart{
		ID:    "cpu_times_%s",
		Title: "CPU Times",
		Units: "increase/s",
		Fam:   "%s CPU",
		Ctx:   "msgmonsys.cpu_times_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystem + "_%s", Name: "system", Algo: module.Incremental, Div: scale(metricCPUTimesSystem)},
			{ID: metricCPUTimesUser + "_%s", Name: "user", Algo: module.Incremental, Div: scale(metricCPUTimesUser)},
			{ID: metricCPUTimesIdle + "_%s", Name: "idle", Algo: module.Incremental, Div: scale(metricCPUTimesIdle)},
		},
	}
	cpuPercentChart = Chart{
		ID:    "cpu_percent_%s",
		Title: "CPU Percent",
		Units: "percentage",
		Fam:   "%s CPU",
		Ctx:   "msgmonsys.cpu_percent_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUPercentCPUPercent + "_%s", Name: "percent", Div: scale(metricCPUPercentCPUPercent)},
		},
	}
	diskUsageChart = Chart{
		ID:    "disk_usage_%s",
		Title: "Disk Usage Percent",
		Units: "percentage",
		Fam:   "%s Disk Usage",
		Ctx:   "msgmonsys.disk_usage_percent_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskUsagePercentUsed + "_%s", Name: "percent", Div: scale(metricDiskUsagePercentUsed)},
		},
	}
	virtualMemoryChart = Chart{
		ID:    "virtual_memory_%s",
		Title: "Virtual Memory",
		Units: "MiB",
		Fam:   "%s Virtual Memory",
		Ctx:   "msgmonsys.virtual_memory_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricVirtualMemoryAvailable + "_%s", Name: "available", Div: scale(metricVirtualMemoryAvailable)},
			{ID: metricVirtualMemoryUsed + "_%s", Name: "used", Div: scale(metricVirtualMemoryUsed)},
		},
	}
	diskIODataChart = Chart{
		ID:    "disk_io_data_rate_%s",
		Title: "Disk I/O Data Rates",
		Units: "kiB/s",
		Fam:   "%s Disk I/O",
		Ctx:   "msgmonsys.disk_io_data_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadBytes + "_%s", Name: "read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadBytes)},
			{ID: metricDiskIOCountersWriteBytes + "_%s", Name: "write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteBytes)},
		},
	}
	diskIOOperationChart = Chart{
		ID:    "disk_io_operation_rate_%s",
		Title: "Disk I/O Operation Rates",
		Units: "ops/s",
		Fam:   "%s Disk I/O",
		Ctx:   "msgmonsys.disk_io_operation_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadCount + "_%s", Name: "read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadCount)},
			{ID: metricDiskIOCountersWriteCount + "_%s", Name: "write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteCount)},
		},
	}
	networkIODataChart = Chart{
		ID:    "network_io_data_rate_%s",
		Title: "Network I/O Data Rates",
		Units: "kiB/s",
		Fam:   "%s Network I/O",
		Ctx:   "msgmonsys.network_io_data_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersBytesRecv + "_%s", Name: "recv", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesRecv)},
			{ID: metricNetworkIOCountersBytesSent + "_%s", Name: "send", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesSent)},
		},
	}
	networkIOOperationChart = Chart{
		ID:    "network_io_operation_rate_%s",
		Title: "Network IO Operation Rates",
		Units: "packets/s",
		Fam:   "%s Network I/O",
		Ctx:   "msgmonsys.network_io_operation_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersPacketsRecv + "_%s", Name: "recv", Algo: module.Incremental, Div: scale(metricNetworkIOCountersPacketsRecv)},
			{ID: metricNetworkIOCountersPacketsSent + "_%s", Name: "send", Algo: module.Incremental, Div: scale(metricNetworkIOCountersPacketsSent)},
		},
	}
)

func (p *MsgmonSys) updateCharts() {
	for s := range p.curCache.systems {
		if !p.cache.systems[s] {
			p.cache.systems[s] = true
			p.addSystemCharts(s)
			dim := &Dim{ID: metricCPUPercentCPUPercent + "_" + s.name, Name: s.name, Div: 100}
			if err := p.charts.Get("summary").AddDim(dim); err != nil {
				p.Warning(fmt.Sprintf("Error adding dimension %s to summary chart: %s", s.name, err.Error()))
			}
		}
	}
	for s := range p.cache.systems {
		if p.curCache.systems[s] {
			continue
		}
		delete(p.cache.systems, s)
		p.charts.Get("summary").MarkDimRemove(metricCPUPercentCPUPercent+"_"+s.name, true)
	}
}

func (p *MsgmonSys) addSystemCharts(s systemName) {
	charts := systemCharts.Copy()
	for _, chart := range *charts {
		chart.ID = fmt.Sprintf(chart.ID, s.name)
		chart.Fam = fmt.Sprintf(chart.Fam, s.name)
		chart.Ctx = fmt.Sprintf(chart.Ctx, s.name)
		for _, dim := range chart.Dims {
			dim.ID = fmt.Sprintf(dim.ID, s.name)
		}
	}
	if err := p.Charts().Add(*charts...); err != nil {
		p.Warning(err)
	}
}
