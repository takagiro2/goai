package main

import (
	"fmt"
	"math/rand"
)

var train_model = [][2]float64{
	{0, 0},
	{1, 2},
	{2, 4},
	{3, 6},
	{4, 8},
	{5, 10},
	{6, 12},
	{7, 14},
	{8, 16},
}

func rand_float(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func cost(w, b float64) float64 {
	var result float64 = 0.0
	for i := 0; i < len(train_model); i++ {
		var x float64 = float64(train_model[i][0])
		var y float64 = x*w + b
		var d float64 = y - train_model[i][1]
		result += d * d
	}
	result /= float64(len(train_model))
	return result
}

func main() {
	var eps float64 = 1e-3
	var rate float64 = 1e-3
	var w float64 = rand_float(0.0, 1.0) * 10.0
	var b float64 = rand_float(0.0, 1.0) * 5.0
	fmt.Printf("%f\n", cost(w, b))

	for i := 0; i < 100; i++ {
		var c float64 = cost(w, b)
		var dw float64 = (cost(w+eps, b) - c) / eps
		var db float64 = (cost(w, eps+b) - c) / eps

		w -= rate * dw
		b -= rate * db

		fmt.Printf("cost = %f, w = %f, b = %f\n", cost(w, b), w, b)
	}

	fmt.Println("------------------------------------")
	fmt.Printf("w = %f, b = %f\n", w, b)
}
