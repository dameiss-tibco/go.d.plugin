// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	"fmt"
	// "strings"

	// "github.com/netdata/go.d.plugin/pkg/prometheus"

	"github.com/netdata/go.d.plugin/agent/module"
)

type (
	Charts = module.Charts
	Chart  = module.Chart
	Dims   = module.Dims
	Dim    = module.Dim
	Opts   = module.Opts
)

// var initialCharts = Charts{}

var systemCharts = Charts{
	cpuTimesChart.Copy(),
}

var (
	cpuTimesChart = Chart{
		ID:    "cpu_times_%s",
		Title: "CPU Times",
		Units: "increase/s",
		Fam:   "%s cputimes",
		Ctx:   "msgmonsys.cpu_times_%s",
		Type:  module.Stacked,
		Opts:  Opts{StoreFirst: true},
		Dims: Dims{
			{ID: metricCPUTimesSystem + "_%s", Name: "system", Algo: module.Incremental, Div: 1000},
			{ID: metricCPUTimesUser + "_%s", Name: "user", Algo: module.Incremental, Div: 1000},
			{ID: metricCPUTimesIdle + "_%s", Name: "idle", Algo: module.Incremental, Div: 1000},
		},
	}
)

// func (p *Pulsar) adjustCharts(pms prometheus.Series) {
// 	if pms := pms.FindByName(metricPulsarStorageReadRate); pms.Len() == 0 || pms[0].Labels.Get("namespace") == "" {
// 		p.removeSummaryChart(sumStorageOperationsRateChart.ID)
// 		p.removeNamespaceChart(nsStorageOperationsChart.ID)
// 		p.removeNamespaceChart(topicStorageReadRateChart.ID)
// 		p.removeNamespaceChart(topicStorageWriteRateChart.ID)
// 		delete(p.topicChartsMapping, topicStorageReadRateChart.ID)
// 		delete(p.topicChartsMapping, topicStorageWriteRateChart.ID)
// 	}
// 	if pms.FindByName(metricPulsarSubscriptionMsgRateRedeliver).Len() == 0 {
// 		p.removeSummaryChart(sumSubsMsgRateRedeliverChart.ID)
// 		p.removeSummaryChart(sumSubsBlockedOnUnackedMsgChart.ID)
// 		p.removeNamespaceChart(nsSubsMsgRateRedeliverChart.ID)
// 		p.removeNamespaceChart(nsSubsBlockedOnUnackedMsgChart.ID)
// 		p.removeNamespaceChart(topicSubsMsgRateRedeliverChart.ID)
// 		p.removeNamespaceChart(topicSubsBlockedOnUnackedMsgChart.ID)
// 		delete(p.topicChartsMapping, topicSubsMsgRateRedeliverChart.ID)
// 		delete(p.topicChartsMapping, topicSubsBlockedOnUnackedMsgChart.ID)
// 	}
// 	if pms.FindByName(metricPulsarReplicationBacklog).Len() == 0 {
// 		p.removeSummaryChart(sumReplicationRateChart.ID)
// 		p.removeSummaryChart(sumReplicationThroughputRateChart.ID)
// 		p.removeSummaryChart(sumReplicationBacklogChart.ID)
// 		p.removeNamespaceChart(nsReplicationRateChart.ID)
// 		p.removeNamespaceChart(nsReplicationThroughputChart.ID)
// 		p.removeNamespaceChart(nsReplicationBacklogChart.ID)
// 		p.removeNamespaceChart(topicReplicationRateInChart.ID)
// 		p.removeNamespaceChart(topicReplicationRateOutChart.ID)
// 		p.removeNamespaceChart(topicReplicationThroughputRateInChart.ID)
// 		p.removeNamespaceChart(topicReplicationThroughputRateOutChart.ID)
// 		p.removeNamespaceChart(topicReplicationBacklogChart.ID)
// 		delete(p.topicChartsMapping, topicReplicationRateInChart.ID)
// 		delete(p.topicChartsMapping, topicReplicationRateOutChart.ID)
// 		delete(p.topicChartsMapping, topicReplicationThroughputRateInChart.ID)
// 		delete(p.topicChartsMapping, topicReplicationThroughputRateOutChart.ID)
// 		delete(p.topicChartsMapping, topicReplicationBacklogChart.ID)
// 	}
// }

// func (p *Pulsar) removeSummaryChart(chartID string) {
// 	if err := p.Charts().Remove(chartID); err != nil {
// 		p.Warning(err)
// 	}
// }

// func (p *Pulsar) removeNamespaceChart(chartID string) {
// 	if err := p.nsCharts.Remove(chartID); err != nil {
// 		p.Warning(err)
// 	}
// }

func (p *MsgmonSys) updateCharts() {
	for s := range p.curCache.systems {
		if !p.cache.systems[s] {
			p.cache.systems[s] = true
			p.addSystemCharts(s)
		}
	}
	for s := range p.cache.systems {
		if p.curCache.systems[s] {
			continue
		}
		delete(p.cache.systems, s)
		// p.removeSystemFromCharts(s)
	}
}

// func (p *Pulsar) updateCharts() {
// 	// NOTE: order is important
// 	for ns := range p.curCache.namespaces {
// 		if !p.cache.namespaces[ns] {
// 			p.cache.namespaces[ns] = true
// 			p.addNamespaceCharts(ns)
// 		}
// 	}
// 	for top := range p.curCache.topics {
// 		if !p.cache.topics[top] {
// 			p.cache.topics[top] = true
// 			p.addTopicToCharts(top)
// 		}
// 	}
// 	for top := range p.cache.topics {
// 		if p.curCache.topics[top] {
// 			continue
// 		}
// 		delete(p.cache.topics, top)
// 		p.removeTopicFromCharts(top)
// 	}
// 	for ns := range p.cache.namespaces {
// 		if p.curCache.namespaces[ns] {
// 			continue
// 		}
// 		delete(p.cache.namespaces, ns)
// 		p.removeNamespaceFromCharts(ns)
// 	}
// }

func (p *MsgmonSys) addSystemCharts(s systemName) {
	charts := p.charts.Copy()
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

// func (p *Pulsar) addNamespaceCharts(ns namespace) {
// 	charts := p.nsCharts.Copy()
// 	for _, chart := range *charts {
// 		chart.ID = fmt.Sprintf(chart.ID, ns.name)
// 		chart.Fam = fmt.Sprintf(chart.Fam, ns.name)
// 		for _, dim := range chart.Dims {
// 			dim.ID = fmt.Sprintf(dim.ID, ns.name)
// 		}
// 	}
// 	if err := p.Charts().Add(*charts...); err != nil {
// 		p.Warning(err)
// 	}
// }

// func (p *Pulsar) removeNamespaceFromCharts(ns namespace) {
// 	for _, chart := range *p.nsCharts {
// 		id := fmt.Sprintf(chart.ID, ns.name)
// 		if chart = p.Charts().Get(id); chart != nil {
// 			chart.MarkRemove()
// 		} else {
// 			p.Warningf("could not remove namespace chart '%s'", id)
// 		}
// 	}
// }

// func (p *Pulsar) addTopicToCharts(top topic) {
// 	for id, metric := range p.topicChartsMapping {
// 		id = fmt.Sprintf(id, top.namespace)
// 		chart := p.Charts().Get(id)
// 		if chart == nil {
// 			p.Warningf("could not add topic '%s' to chart '%s': chart not found", top.name, id)
// 			continue
// 		}

// 		dim := Dim{ID: metric + "_" + top.name, Name: extractTopicName(top)}
// 		switch metric {
// 		case metricPulsarThroughputIn,
// 			metricPulsarThroughputOut,
// 			metricPulsarReplicationThroughputIn,
// 			metricPulsarReplicationThroughputOut:
// 			dim.Div = 1024 * 1000
// 		case metricPulsarRateIn,
// 			metricPulsarRateOut,
// 			metricPulsarStorageWriteRate,
// 			metricPulsarStorageReadRate,
// 			metricPulsarSubscriptionMsgRateRedeliver,
// 			metricPulsarReplicationRateIn,
// 			metricPulsarReplicationRateOut:
// 			dim.Div = 1000
// 		case metricPulsarStorageSize:
// 			dim.Div = 1024
// 		}

// 		if err := chart.AddDim(&dim); err != nil {
// 			p.Warning(err)
// 		}
// 		chart.MarkNotCreated()
// 	}
// }

// func (p *Pulsar) removeTopicFromCharts(top topic) {
// 	for id, metric := range p.topicChartsMapping {
// 		id = fmt.Sprintf(id, top.namespace)
// 		chart := p.Charts().Get(id)
// 		if chart == nil {
// 			p.Warningf("could not remove topic '%s' from chart '%s': chart not found", top.name, id)
// 			continue
// 		}

// 		if err := chart.MarkDimRemove(metric+"_"+top.name, true); err != nil {
// 			p.Warning(err)
// 		}
// 		chart.MarkNotCreated()
// 	}
// }

// func topicChartsMapping() map[string]string {
// 	return map[string]string{
// 		topicSubscriptionsChart.ID:                metricPulsarSubscriptionsCount,
// 		topicProducersChart.ID:                    metricPulsarProducersCount,
// 		topicConsumersChart.ID:                    metricPulsarConsumersCount,
// 		topicMessagesRateInChart.ID:               metricPulsarRateIn,
// 		topicMessagesRateOutChart.ID:              metricPulsarRateOut,
// 		topicThroughputRateInChart.ID:             metricPulsarThroughputIn,
// 		topicThroughputRateOutChart.ID:            metricPulsarThroughputOut,
// 		topicStorageSizeChart.ID:                  metricPulsarStorageSize,
// 		topicStorageReadRateChart.ID:              metricPulsarStorageReadRate,
// 		topicStorageWriteRateChart.ID:             metricPulsarStorageWriteRate,
// 		topicMsgBacklogSizeChart.ID:               metricPulsarMsgBacklog,
// 		topicSubsDelayedChart.ID:                  metricPulsarSubscriptionDelayed,
// 		topicSubsMsgRateRedeliverChart.ID:         metricPulsarSubscriptionMsgRateRedeliver,
// 		topicSubsBlockedOnUnackedMsgChart.ID:      metricPulsarSubscriptionBlockedOnUnackedMessages,
// 		topicReplicationRateInChart.ID:            metricPulsarReplicationRateIn,
// 		topicReplicationRateOutChart.ID:           metricPulsarReplicationRateOut,
// 		topicReplicationThroughputRateInChart.ID:  metricPulsarReplicationThroughputIn,
// 		topicReplicationThroughputRateOutChart.ID: metricPulsarReplicationThroughputOut,
// 		topicReplicationBacklogChart.ID:           metricPulsarReplicationBacklog,
// 	}
// }

// func extractTopicName(top topic) string {
// 	// persistent://sample/ns1/demo-1 => p:demo-1
// 	if idx := strings.LastIndexByte(top.name, '/'); idx > 0 {
// 		return top.name[:1] + ":" + top.name[idx+1:]
// 	}
// 	return top.name
// }
