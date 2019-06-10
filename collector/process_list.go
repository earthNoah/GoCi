package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/process"
)

type processListCollector struct {
	metric []typedDesc
}

func init() {
	registerCollector("process", defaultEnabled, NewProcessListCollector)
}

// NewLoadavgCollector returns a new Collector exposing load average stats.
func NewProcessListCollector() (Collector, error) {
	return &loadavgCollector{
		metric: []typedDesc{
			{prometheus.NewDesc(namespace+"_load1", "1m load average.", nil, nil), prometheus.GaugeValue},
			{prometheus.NewDesc(namespace+"_load5", "5m load average.", nil, nil), prometheus.GaugeValue},
			{prometheus.NewDesc(namespace+"_load15", "15m load average.", nil, nil), prometheus.GaugeValue},
		},
	}, nil
}

type ProcessEntity struct {
	pid  int32
	name string
	user string
	CP   float64
	MP   float32
	RM   uint64
}

func main() {
	fmt.Println("进程统计:")
	pi, _ := process.Pids()

	for i := int32(0); i < int32(len(pi)); i++ {
		var pid = pi[i]
		p := &process.Process{Pid: pid}
		n, err1 := p.Name()
		c, err2 := p.CPUPercent()
		m, err3 := p.MemoryPercent()
		u, err4 := p.Username()
		ccc, err4 := p.MemoryInfo()
		if err4 == nil {
			var realMB = ccc.RSS / 1024 / 1024
			if err1 == nil && err2 == nil && err3 == nil {
				var processEntity = ProcessEntity{pid, n, u, c, m, realMB}
				fmt.Println(processEntity)
			}
		}
	}

}
