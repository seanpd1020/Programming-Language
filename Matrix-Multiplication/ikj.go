// ikj
package main

import (
	"fmt"
	"os"
	"time"
)

var matrix1, matrix2, matrix3 [4096][4096]int
var ch chan int = make(chan int, 3000)

func main() {
	file, err := os.Open("test3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var n1, n2, n3, n4 int

	fmt.Fscanln(file, &n1, &n2)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			fmt.Fscanf(file, "%d", &matrix1[i][j])
		}
	}

	fmt.Fscanln(file, &n3, &n4)
	for i := 0; i < n3; i++ {
		for j := 0; j < n4; j++ {
			fmt.Fscanf(file, "%d", &matrix2[i][j])
		}
	}

	fmt.Println("Start Computing...")
	s := time.Now()

	go m_m1(n1, n4, n2)
	go m_m2(n1, n4, n2)
	go m_m3(n1, n4, n2)
	go m_m4(n1, n4, n2)

	<-ch
	<-ch
	<-ch
	<-ch

	fmt.Println(time.Since(s))
	//output
	file2, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	for row := 0; row < n1; row++ {
		for col := 0; col < n4; col++ {
			fmt.Fprintf(file2, "%d ", matrix3[row][col])
		}
		fmt.Fprintf(file2, "\n")
	}

}

func m_m1(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			if row < n1/4 {

				for col := 0; col < n4; col++ {
					matrix3[row][col] += matrix1[row][i] * matrix2[i][col]
				}
			}
		}
	}
	ch <- 1
}

func m_m2(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			if row >= n1/4 && row < n1*2/4 {

				for col := 0; col < n4; col++ {
					matrix3[row][col] += matrix1[row][i] * matrix2[i][col]
				}
			}
		}
	}
	ch <- 1
}

func m_m3(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			if row >= n1*2/4 && row < n1*3/4 {

				for col := 0; col < n4; col++ {
					matrix3[row][col] += matrix1[row][i] * matrix2[i][col]
				}
			}
		}
	}
	ch <- 1
}

func m_m4(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			if row >= n1*3/4 {

				for col := 0; col < n4; col++ {
					matrix3[row][col] += matrix1[row][i] * matrix2[i][col]
				}
			}
		}
	}
	ch <- 1
}
