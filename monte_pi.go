// monte_pi use three functions for determine PI value using Monte Carlo method: http://en.wikipedia.org/wiki/Monte_Carlo_method#Introduction
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// GetPISimple use the classic iteration for determine PI
func GetPISimple(samples int) float64 {
	var inside int = 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	return ratio * 4
}

// GetPIMultiChannel given the samples and cpu input value, will use n-cpu routine to do the work
func GetPIMultiChannel(samples int, cpu int) float64 {

	results := make(chan float64, cpu)
	threadSamples := samples / cpu

	for j := 0; j < cpu; j++ {

		go func() {
			var inside int = 0

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()
				if (x*x + y*y) < 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var piTotal float64
	for t := 0; t < cpu; t++ {
		piTotal += <-results
	}
	return piTotal / float64(cpu)
}

// GetPIMultiCPU given the samples, will use max available threads to do the work
func GetPIMultiCPU(samples int) float64 {

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())

	results := make(chan float64, cpu)
	threadSamples := samples / cpu

	for j := 0; j < cpu; j++ {

		go func() {
			var inside int = 0

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()
				if (x*x + y*y) < 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var piTotal float64
	for t := 0; t < cpu; t++ {
		piTotal += <-results
	}
	return piTotal / float64(cpu)
}

// init set the seed for rand function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// sampleCount is the base for the test. More high is this value, more precise PI results.
const sampleCount = 10000000

func main() {

	fmt.Println("Let's find some PI:")
	fmt.Println("Simple -->", GetPISimple(sampleCount))
	fmt.Println("Using channel -->", GetPIMultiChannel(sampleCount, 4))
	fmt.Println("Using chanel and cpu --> ", GetPIMultiCPU(sampleCount))
}
