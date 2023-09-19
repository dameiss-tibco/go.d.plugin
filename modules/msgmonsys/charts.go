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
		Fam:   "CPU Percent",
		Ctx:   "msgmonsys.cpu_percent",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
	}
)

var systemCharts = Charts{
	// cpuTimesChart.Copy(),
	cpuTimesIncrChart.Copy(),
	// cpuTimesRawChart.Copy(),
	cpuPercentChart.Copy(),
	diskUsageChart.Copy(),
	virtualMemoryChart.Copy(),
	// diskIODataChart.Copy(),
	diskIODataIncrChart.Copy(),
	// diskIOOperationChart.Copy(),
	diskIOOperationIncrChart.Copy(),
	// networkIODataChart.Copy(),
	networkIODataIncrChart.Copy(),
	// networkIOOperationChart.Copy(),
	networkIOOperationIncrChart.Copy(),
}

var (
	cpuTimesChart = Chart{
		ID:    "cpu_times_%s",
		Title: "CPU Times",
		Units: "s/s",
		Fam:   "%s CPU Times",
		Ctx:   "msgmonsys.cpu_times_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystemRate + "_%s", Name: "System", Div: scale(metricCPUTimesSystemRate)},
			{ID: metricCPUTimesUserRate + "_%s", Name: "User", Div: scale(metricCPUTimesUserRate)},
			{ID: metricCPUTimesIdleRate + "_%s", Name: "Idle", Div: scale(metricCPUTimesIdleRate)},
		},
	}
	cpuTimesIncrChart = Chart{
		ID:    "cpu_times_incr_%s",
		Title: "CPU Times (Incr)",
		Units: "s/s",
		Fam:   "%s CPU Times",
		Ctx:   "msgmonsys.cpu_times_incr_%s",
		Type:  module.Area,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystem + "_%s", Name: "System", Algo: module.Incremental, Div: scale(metricCPUTimesSystem)},
			{ID: metricCPUTimesUser + "_%s", Name: "User", Algo: module.Incremental, Div: scale(metricCPUTimesUser)},
			{ID: metricCPUTimesIdle + "_%s", Name: "Idle", Algo: module.Incremental, Div: scale(metricCPUTimesIdle)},
		},
	}
	cpuTimesRawChart = Chart{
		ID:    "cpu_times_raw_%s",
		Title: "CPU Times (Raw)",
		Units: "seconds",
		Fam:   "%s CPU Times",
		Ctx:   "msgmonsys.cpu_times_raw_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystem + "_%s", Name: "System", Div: scale(metricCPUTimesSystem)},
			{ID: metricCPUTimesUser + "_%s", Name: "User", Div: scale(metricCPUTimesUser)},
			{ID: metricCPUTimesIdle + "_%s", Name: "Idle", Div: scale(metricCPUTimesIdle)},
		},
	}
	cpuPercentChart = Chart{
		ID:    "cpu_percent_%s",
		Title: "CPU Percent",
		Units: "percentage",
		Fam:   "%s CPU Percent",
		Ctx:   "msgmonsys.cpu_percent_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUPercentCPUPercent + "_%s", Name: "Percent", Div: scale(metricCPUPercentCPUPercent)},
			// {ID: "cpupct_rate_0_%s", Name: "bottom", DimOpts: DimOpts{Hidden: true}},
			// {ID: "cpupct_rate_100_%s", Name: "top", DimOpts: DimOpts{Hidden: true}},
		},
	}
	diskUsageChart = Chart{
		ID:    "disk_usage_%s",
		Title: "Disk Usage Percent",
		Units: "percent",
		Fam:   "%s Disk Usage",
		Ctx:   "msgmonsys.disk_usage_percent_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskUsagePercentUsed + "_%s", Name: "Percent", Div: scale(metricDiskUsagePercentUsed)},
		},
	}
	virtualMemoryChart = Chart{
		ID:    "virtual_memory_%s",
		Title: "Virtual Memory",
		Units: "B",
		Fam:   "%s Virtual Memory",
		Ctx:   "msgmonsys.virtual_memory_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricVirtualMemoryAvailable + "_%s", Name: "Available", Div: scale(metricVirtualMemoryAvailable)},
			{ID: metricVirtualMemoryUsed + "_%s", Name: "Used", Div: scale(metricVirtualMemoryUsed)},
		},
	}
	diskIODataChart = Chart{
		ID:    "disk_io_data_rate_%s",
		Title: "Disk I/O Data Rates",
		Units: "B/s",
		Fam:   "%s Disk I/O",
		Ctx:   "msgmonsys.disk_io_data_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadBytesRate + "_%s", Name: "Read", Div: scale(metricDiskIOCountersReadBytesRate)},
			{ID: metricDiskIOCountersWriteBytesRate + "_%s", Name: "Write", Div: scale(metricDiskIOCountersWriteBytesRate)},
		},
	}
	diskIODataIncrChart = Chart{
		ID:    "disk_io_data_rate_incr_%s",
		Title: "Disk I/O Data Rates Incr",
		Units: "B/s",
		Fam:   "%s Disk I/O",
		Ctx:   "msgmonsys.disk_io_data_rate_incr_%s",
		Type:  module.Area,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadBytes + "_%s", Name: "Read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadBytes)},
			{ID: metricDiskIOCountersWriteBytes + "_%s", Name: "Write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteBytes)},
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
			{ID: metricDiskIOCountersReadCountRate + "_%s", Name: "Read", Div: scale(metricDiskIOCountersReadCountRate)},
			{ID: metricDiskIOCountersWriteCountRate + "_%s", Name: "Write", Div: scale(metricDiskIOCountersWriteCountRate)},
		},
	}
	diskIOOperationIncrChart = Chart{
		ID:    "disk_io_operation_rate_incr_%s",
		Title: "Disk I/O Operation Rates",
		Units: "ops/s",
		Fam:   "%s Disk I/O",
		Ctx:   "msgmonsys.disk_io_operation_rate_incr_%s",
		Type:  module.Area,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricDiskIOCountersReadCount + "_%s", Name: "Read", Algo: module.Incremental, Div: scale(metricDiskIOCountersReadCount)},
			{ID: metricDiskIOCountersWriteCount + "_%s", Name: "Write", Algo: module.Incremental, Div: scale(metricDiskIOCountersWriteCount)},
		},
	}
	networkIODataChart = Chart{
		ID:    "network_io_data_rate_%s",
		Title: "Network I/O Data Rates",
		Units: "B/s",
		Fam:   "%s Network I/O",
		Ctx:   "msgmonsys.network_io_data_rate_%s",
		Type:  module.Line,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersBytesRecvRate + "_%s", Name: "Receive", Div: scale(metricNetworkIOCountersBytesRecvRate)},
			{ID: metricNetworkIOCountersBytesSentRate + "_%s", Name: "Send", Div: scale(metricNetworkIOCountersBytesSentRate)},
		},
	}
	networkIODataIncrChart = Chart{
		ID:    "network_io_data_rate_incr_%s",
		Title: "Network I/O Data Rates",
		Units: "B/s",
		Fam:   "%s Network I/O",
		Ctx:   "msgmonsys.network_io_data_rate_incr_%s",
		Type:  module.Area,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricNetworkIOCountersBytesRecv + "_%s", Name: "Receive", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesRecv)},
			{ID: metricNetworkIOCountersBytesSent + "_%s", Name: "Send", Algo: module.Incremental, Div: scale(metricNetworkIOCountersBytesSent)},
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
			{ID: metricNetworkIOCountersPacketsRecvRate + "_%s", Name: "Receive", Div: scale(metricNetworkIOCountersPacketsRecvRate)},
			{ID: metricNetworkIOCountersPacketsSentRate + "_%s", Name: "Send", Div: scale(metricNetworkIOCountersPacketsSentRate)},
		},
	}
	networkIOOperationIncrChart = Chart{
		ID:    "network_io_operation_rate_incr_%s",
		Title: "Network IO Operation Rates",
		Units: "packets/s",
		Fam:   "%s Network I/O",
		Ctx:   "msgmonsys.network_io_operation_rate_incr_%s",
		Type:  module.Area,
		Opts:  Opts{StoreFirst: true},
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
