package mpdice

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// DicePlugin mackerel plugin
type DicePlugin struct {
	Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (u DicePlugin) MetricKeyPrefix() string {
	if u.Prefix == "" {
		u.Prefix = "dice"
	}
	return u.Prefix
}

// GraphDefinition interface for mackerelplugin
func (u DicePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(u.Prefix)
	return map[string]mp.Graphs{
		"": {
			Label: labelPrefix,
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "dice", Label: "dice"},
			},
		},
	}
}

// FetchMetrics interface for mackerelplugin
func (u DicePlugin) FetchMetrics() (map[string]float64, error) {
	var err error
	var result float64

	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	result = float64(rand.Intn(6) + 1)

	return map[string]float64{"dice": result}, nil
}

// Do the plugin
func Do() {
	optPrefix := flag.String("metric-key-prefix", "dice", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	u := DicePlugin{
		Prefix: *optPrefix,
	}
	helper := mp.NewMackerelPlugin(u)
	helper.Tempfile = *optTempfile
	helper.Run()
}
