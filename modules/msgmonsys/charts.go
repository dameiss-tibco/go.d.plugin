// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	"fmt"
	"github.com/netdata/go.d.plugin/agent/module"
)

type (
	Charts  = module.Charts
	Chart   = module.Chart
	Dims    = module.Dims
	Dim     = module.Dim
	Opts    = module.Opts
	DimOpts = module.DimOpts
)

var summaryCharts = Charts{
	summaryChart.Copy(),
}

var (
	summaryChart = Chart{
		ID:    "summary",
		Title: "CPU Percent Summary",
		Units: "percentage",
		Fam:   "msgmonsys.summary",
		Ctx:   "msgmonsys",
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
		ID:    "cputimes_%s",
		Title: "CPU Times",
		Units: "s/s",
		// Fam:   "msgmonsys_cputimes.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.cputimes_%s",
		Ctx:  "msgmonsys.cputimes",
		Type: module.Area,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystem + "_%s", Name: "System", Algo: module.Incremental, Div: scale(metricCPUTimesSystem)},
			{ID: metricCPUTimesUser + "_%s", Name: "User", Algo: module.Incremental, Div: scale(metricCPUTimesUser)},
			{ID: metricCPUTimesIdle + "_%s", Name: "Idle", Algo: module.Incremental, Div: scale(metricCPUTimesIdle)},
		},
	}
	cpuPercentChart = Chart{
		ID:    "cpupercent_%s",
		Title: "CPU Percent",
		Units: "percent",
		// Fam:   "msgmonsys_cpupercent.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.cpupercent_%s",
		Ctx:  "msgmonsys.cpupercent",
		Type: module.Line,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUPercentCPUPercent + "_%s", Name: "Percent", Div: scale(metricCPUPercentCPUPercent)},
		},
	}
	diskUsageChart = Chart{
		ID:    "diskusage_%s",
		Title: "Disk Usage",
		Units: "percent",
		// Fam:   "msgmonsys_diskusage.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.diskusage_percent_%s",
		Ctx:  "msgmonsys.diskusage",
		Type: module.Line,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskUsagePercentUsed + "_%s", Name: "Percent", Div: scale(metricDiskUsagePercentUsed)},
		},
	}
	virtualMemoryChart = Chart{
		ID:    "virtualmemory_%s",
		Title: "Virtual Memory",
		Units: "B",
		// Fam:   "msgmonsys_virtualmemory.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.virtualmemory_%s",
		Ctx:  "msgmonsys.virtualmemory",
		Type: module.Line,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricVirtualMemoryAvailable + "_%s", Name: "Available", Div: scale(metricVirtualMemoryAvailable)},
			{ID: metricVirtualMemoryUsed + "_%s", Name: "Used", Div: scale(metricVirtualMemoryUsed)},
		},
	}
	diskIODataChart = Chart{
		ID:    "diskdatarate_%s",
		Title: "Disk I/O Data Rates",
		Units: "B/s",
		// Fam:   "msgmonsys_diskio_datarate.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.disk_io_data_rate_incr_%s",
		Ctx:  "msgmonsys.diskdatarate",
		Type: module.Area,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadBytes + "_%s", Name: "Read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadBytes)},
			{ID: metricDiskIOCountersWriteBytes + "_%s", Name: "Write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteBytes)},
		},
	}
	diskIOOperationChart = Chart{
		ID:    "diskoprate_%s",
		Title: "Disk I/O Operation Rates",
		Units: "ops/s",
		// Fam:   "msgmonsys_diskio_oprate.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.diskio_operationrate_%s",
		Ctx:  "msgmonsys.diskoprate",
		Type: module.Area,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadCount + "_%s", Name: "Read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadCount)},
			{ID: metricDiskIOCountersWriteCount + "_%s", Name: "Write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteCount)},
		},
	}
	networkIODataChart = Chart{
		ID:    "networkdatarate_%s",
		Title: "Network I/O Data Rates",
		Units: "B/s",
		// Fam:   "msgmonsys_networkio_datarate.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.networkio_datarate_%s",
		Ctx:  "msgmonsys.networkdatarate",
		Type: module.Area,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersBytesRecv + "_%s", Name: "Receive", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesRecv)},
			{ID: metricNetworkIOCountersBytesSent + "_%s", Name: "Send", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesSent)},
		},
	}
	networkIOOperationChart = Chart{
		ID:    "networkoprate_%s",
		Title: "Network IO Operation Rates",
		Units: "packets/s",
		// Fam:   "msgmonsys_networkio_oprate.%s",
		Fam: "%s",
		// Ctx:   "msgmonsys.networkio_operationrate_%s",
		Ctx:  "msgmonsys.networkoprate",
		Type: module.Area,
		Opts: Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersPacketsRecv + "_%s", Name: "Receive", Algo: module.Incremental, Div: scale(metricNetworkIOCountersPacketsRecv)},
			{ID: metricNetworkIOCountersPacketsSent + "_%s", Name: "Send", Algo: module.Incremental, Div: scale(metricNetworkIOCountersPacketsSent)},
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
		// chart.Title = fmt.Sprintf(chart.Title, s.name)
		chart.ID = fmt.Sprintf(chart.ID, s.name)
		chart.Fam = fmt.Sprintf(chart.Fam, s.name)
		// chart.Ctx = fmt.Sprintf(chart.Ctx, s.name)
		for _, dim := range chart.Dims {
			dim.ID = fmt.Sprintf(dim.ID, s.name)
		}
	}
	if err := p.Charts().Add(*charts...); err != nil {
		p.Warning(err)
	}
}
