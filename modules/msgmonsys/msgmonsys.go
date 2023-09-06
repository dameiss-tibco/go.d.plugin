// SPDX-License-Identifier: GPL-3.0-or-later

package msgmonsys

import (
	"errors"
	"github.com/netdata/go.d.plugin/agent/module"
	"github.com/netdata/go.d.plugin/pkg/prometheus"
	"github.com/netdata/go.d.plugin/pkg/web"
	"sync"
	"time"
)

func init() {
	creator := module.Creator{
		Defaults: module.Defaults{
			UpdateEvery: 10,
		},
		Create: func() module.Module { return New() },
	}
	module.Register("msgmonsys", creator)
}

func New() *MsgmonSys {
	config := Config{
		HTTP: web.HTTP{
			Request: web.Request{
				URL: "http://127.0.0.1:8080/metrics",
			},
			Client: web.Client{
				Timeout: web.Duration{Duration: time.Second},
			},
		},
	}
	return &MsgmonSys{
		Config:   config,
		once:     &sync.Once{},
		charts:   summaryCharts.Copy(),
		cache:    newCache(),
		curCache: newCache(),
	}
}

type (
	Config struct {
		web.HTTP `yaml:",inline"`
	}

	MsgmonSys struct {
		module.Base
		Config `yaml:",inline"`

		prom     prometheus.Prometheus
		cache    *cache
		curCache *cache
		once     *sync.Once
		charts   *Charts
	}

	systemName struct{ name string }
	cache      struct {
		systems map[systemName]bool
	}
)

func newCache() *cache {
	return &cache{
		systems: make(map[systemName]bool),
	}
}

func (p MsgmonSys) validateConfig() error {
	if p.URL == "" {
		return errors.New("URL is not set")
	}
	return nil
}

func (p *MsgmonSys) initClient() error {
	client, err := web.NewHTTPClient(p.Client)
	if err != nil {
		return err
	}

	p.prom = prometheus.New(client, p.Request)
	return nil
}

func (p *MsgmonSys) Init() bool {
	if err := p.validateConfig(); err != nil {
		p.Errorf("config validation: %v", err)
		return false
	}
	if err := p.initClient(); err != nil {
		p.Errorf("client initializing: %v", err)
		return false
	}
	return true
}

func (p *MsgmonSys) Check() bool {
	return len(p.Collect()) > 0
}

func (p *MsgmonSys) Charts() *Charts {
	return p.charts
}

func (p *MsgmonSys) Collect() map[string]int64 {
	mx, err := p.collect()
	if err != nil {
		p.Error(err)
	}

	if len(mx) == 0 {
		return nil
	}
	return mx
}

func (MsgmonSys) Cleanup() {}
