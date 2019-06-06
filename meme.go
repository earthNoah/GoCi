package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func main() {
	fmt.Println("进程统计:")
	pi, _ := process.Pids()

	for i := int32(0); i < int32(len(pi)); i++ {
		p := &process.Process{Pid: i}
		v, e := p.IsRunning()
		if e != nil {
			fmt.Println(e)
		}
		if v == true {
			pName,_ := p.Name()
			fmt.Print(pName)
			fmt.Print(p.CPUPercent())
			fmt.Print(p.MemoryPercent())
			fmt.Println("")
		}
	}
}