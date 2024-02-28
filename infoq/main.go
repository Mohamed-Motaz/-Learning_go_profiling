package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

func main() {
	// Create a CPU profile file
	cpuF, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer cpuF.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuF); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Create a MEM profile file
	memF, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer memF.Close()

	// Start MEM profiling
	if err := pprof.WriteHeapProfile(memF); err != nil {
		panic(err)
	}

	// Start tracing
	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()

	// Generate some random numbers and calculate their squares
	for i := 0; i < 1000000; i++ {
		n := rand.Intn(100)
		s := square(n)
		fmt.Printf("%d^2 = %d\n", n, s)
	}
}

func square(n int) int {
	return n * n
}
