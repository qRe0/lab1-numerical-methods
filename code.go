package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

// Исходная функция f1(x)
func f1(x float64) float64 {
	return x * math.Cos(x+5)
}

// Исходная функция f2(x)
func f2(x float64) float64 {
	return 1 / (1 + 25*x*x)
}

// Равностоящие узлы: xi = a + i * ((b-a) / n)
func equdistant(a, b float64, n int) []float64 {
	eqnodes := make([]float64, n+1)
	for i := 0; i < n+1; i++ {
		eqnodes[i] = a + float64(i)*(b-a)/float64(n)
	}
	return eqnodes
}

// Чебышевские узлы: xi = (a+b)/2 + (b-a)/2 * cos((2i+1)pi / 2n+2)
func chebshev(a, b float64, n int) []float64 {
	chebnodes := make([]float64, n+1)
	for i := 0; i < n+1; i++ {
		chebnodes[i] = (a+b)/2 + (b-a)/2*math.Cos(float64(2*i+1)*math.Pi/(2*float64(n)+2))
	}
	return chebnodes
}

// Таблица разделенных разностей
func divDiff(x, y []float64) [][]float64 {
	n := len(x)
	ddt := make([][]float64, n)
	for i := range ddt {
		ddt[i] = make([]float64, n)
		ddt[i][0] = y[i]
	}
	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			ddt[i][j] = (ddt[i+1][j-1] - ddt[i][j-1]) / (x[i+j] - x[i])
		}
	}
	return ddt
}

// Функция для вычисления значения интерполяционного многочлена Ньютона
func newtonInterp(x, y []float64, value float64) float64 {
	ddt := divDiff(x, y)
	n := len(x)
	result := ddt[0][0]
	prod := 1.0

	for i := 1; i < n; i++ {
		prod *= (value - x[i-1])
		result += ddt[0][i] * prod
	}

	return result
}

func arrMax(arr []float64) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[len(arr)-1]
}

func main() {
	a, b := -5.0, 5.0
	n := 10

	eqnodes := equdistant(a, b, n)
	chebnodes := chebshev(a, b, n)

	eq1 := make([]float64, n+1)
	eq2 := make([]float64, n+1)
	cheb1 := make([]float64, n+1)
	cheb2 := make([]float64, n+1)

	for i := 0; i <= n; i++ {
		eq1[i] = f1(eqnodes[i])
		eq2[i] = f2(eqnodes[i])
		cheb1[i] = f1(chebnodes[i])
		cheb2[i] = f2(chebnodes[i])
	}

	fmt.Printf("Number of nodes: %d\n", n)
	fmt.Println("Equidistant nodes:")
	for i := 0; i <= n; i++ {
		fmt.Printf("f1(%f) = %f, f2(%f) = %f\n", eqnodes[i], eq1[i], eqnodes[i], eq2[i])
	}

	fmt.Println("Chebyshev nodes:")
	for i := 0; i <= n; i++ {
		fmt.Printf("f1(%f) = %f, f2(%f) = %f\n", chebnodes[i], cheb1[i], chebnodes[i], cheb2[i])
	}

	d := make([]float64, 101)
	for i := 0; i <= 100; i++ {
		x := a + float64(i)*(b-a)/100
		di := math.Abs(newtonInterp(eqnodes, eq1, x) - f1(x))
		d[i] = di
	}

	fmt.Println("-------------------------------")

	maxEl := arrMax(d)
	fmt.Printf("|d| = %.8f\n", maxEl)

	fmt.Println("-------------------------------")

	// Выводим результат в соответсвующие файлы
	filenames := []string{"eq1.txt", "eq2.txt", "cheb1.txt", "cheb2.txt"}
	nodes := [][]float64{eqnodes, eqnodes, chebnodes, chebnodes}
	vals := [][]float64{eq1, eq2, cheb1, cheb2}
	for i, name := range filenames {
		file, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			file, err = os.Create(name)
			if err != nil {
				fmt.Println("Cannot create file: ", err)
				return
			}
		} else {
			file.Truncate(0)
		}
		defer file.Close()

		fmt.Fprintln(file, n+1)

		for j := 0; j <= n; j++ {
			fmt.Fprintf(file, "%f %f\n", nodes[i][j], vals[i][j])
		}

		for j := 0; j <= 100; j++ {
			x := a + float64(j)*(b-a)/100
			fmt.Fprintf(file, "%f %f\n", x, newtonInterp(nodes[i], vals[i], x))
		}
	}
}
