package main

import (
	"os"
	"strconv"
	"fmt"
)

import "runtime"

var step float64
var num_steps int
var _numCPUs = runtime.NumCPU()

func _init_numCPUs() {
	runtime.GOMAXPROCS(_numCPUs)
}
func main() {
	_init_numCPUs()
	num_steps, _ = strconv.Atoi(os.Args[1])
	var x, sum float64
	step = 1.0 / float64(num_steps)
	var _barrier_0_float64 = make(chan float64)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int) {
			var ()
			var sum float64
			for i := _routine_num + 0; i < (num_steps+0)/1; i += _numCPUs {
				x = (float64(i) + 0.5) * step
				sum = sum + 4.0/(1.0+x*x)
			}
			_barrier_0_float64 <- sum
		}(_i)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		sum += <-_barrier_0_float64
	}

	sum = step * sum
	fmt.Println("Pi=", sum)
}
