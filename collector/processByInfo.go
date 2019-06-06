package collector

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

type ProcessEntity struct {
	pid  int32
	name string
	user string
	CP   float64
	MP   float32
	RM   uint64
	URM  float32
}

func main() {
	fmt.Println("进程统计:")
	pi, _ := process.Pids()
	totalMem, err0 := mem.VirtualMemory()

	if err0 == nil {
		var totalMemMB = float32(totalMem.Total) / 1024 / 1024
		for i := int32(0); i < int32(len(pi)); i++ {
			var pid = pi[i]
			p := &process.Process{Pid: pid}
			n, err1 := p.Name()
			c, err2 := p.CPUPercent()
			m, err3 := p.MemoryPercent()
			u, err4 := p.Username()
			var mMB = float32(totalMemMB) * m
			ccc, err4 := p.MemoryInfo()
			if err4 == nil {
				var realMB = ccc.RSS / 1024 / 1024
				if err1 == nil && err2 == nil && err3 == nil {
					var processEntity = ProcessEntity{pid, n, u, c, m, realMB, mMB}
					fmt.Println(processEntity)
				}
			}
		}
	}

}
