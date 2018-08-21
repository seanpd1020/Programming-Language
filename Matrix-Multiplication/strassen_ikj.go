//strassen + ikj
package main

import (
	"fmt"
	"os"
	"time"
)

var matrix1, matrix2 [4096][4096]int
var ch chan int = make(chan int, 7)
var p1, p2, p3, p4, p5, p6, p7 [2048][2048]int
var add1, add2, add3, add4, add5, add6, add7, add8, add9, add10 [2048][2048]int
var a11, a12, a21, a22, b11, b12, b21, b22 [2048][2048]int

func main() {
	file, err := os.Open("test2")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var n1, n2, n3, n4 int
	var matrix3 [4096][4096]int
	var c11, c12, c21, c22 [2048][2048]int
	fmt.Fscanln(file, &n1, &n2)

	//read matrix data from file and partition them
	//matrixA
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			fmt.Fscanf(file, "%d", &matrix1[i][j])
			if i < n1/2 && j < n2/2 {
				a11[i][j] = matrix1[i][j]
			}
			if i < n1/2 && j >= n2/2 {
				a12[i][j-n2/2] = matrix1[i][j]
			}
			if i >= n1/2 && j < n2/2 {
				a21[i-n1/2][j] = matrix1[i][j]
			}
			if i >= n1/2 && j >= n2/2 {
				a22[i-n1/2][j-n2/2] = matrix1[i][j]
			}
		}
	}
	//matrixB
	fmt.Fscanln(file, &n3, &n4)
	for i := 0; i < n3; i++ {
		for j := 0; j < n4; j++ {
			fmt.Fscanf(file, "%d", &matrix2[i][j])
			if i < n1/2 && j < n2/2 {
				b11[i][j] = matrix2[i][j]
			}
			if i < n1/2 && j >= n2/2 {
				b12[i][j-n2/2] = matrix2[i][j]
			}
			if i >= n1/2 && j < n2/2 {
				b21[i-n1/2][j] = matrix2[i][j]
			}
			if i >= n1/2 && j >= n2/2 {
				b22[i-n1/2][j-n2/2] = matrix2[i][j]
			}
		}
	}

	new_n1 := n1 / 2
	new_n2 := n2 / 2
	new_n4 := n4 / 2

	//prepare for compute p1 - p7
	for i := 0; i < new_n1; i++ {
		for j := 0; j < new_n2; j++ {
			add1[i][j] = a11[i][j] + a22[i][j]
			add2[i][j] = b11[i][j] + b22[i][j]
			add3[i][j] = a21[i][j] + a22[i][j]
			add4[i][j] = b12[i][j] - b22[i][j]
			add5[i][j] = b21[i][j] - b11[i][j]
			add6[i][j] = a11[i][j] + a12[i][j]
			add7[i][j] = a21[i][j] - a11[i][j]
			add8[i][j] = b11[i][j] + b12[i][j]
			add9[i][j] = a12[i][j] - a22[i][j]
			add10[i][j] = b21[i][j] + b22[i][j]
		}
	}

	fmt.Println("Start Computing...")
	s := time.Now()
	//computing p1 - p7
	//matrix multiple using goroutine for concurrency
	go m_m1(new_n1, new_n4, new_n2)
	go m_m2(new_n1, new_n4, new_n2)
	go m_m3(new_n1, new_n4, new_n2)
	go m_m4(new_n1, new_n4, new_n2)
	go m_m5(new_n1, new_n4, new_n2)
	go m_m6(new_n1, new_n4, new_n2)
	go m_m7(new_n1, new_n4, new_n2)

	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch

	//fmt.Println(time.Since(s))
	//calculate c1 -- c4
	for row := 0; row < new_n1; row++ {
		for col := 0; col < new_n2; col++ {
			c11[row][col] = p1[row][col] + p4[row][col] - p5[row][col] + p7[row][col]
			c12[row][col] = p3[row][col] + p5[row][col]
			c21[row][col] = p2[row][col] + p4[row][col]
			c22[row][col] = p1[row][col] - p2[row][col] + p3[row][col] + p6[row][col]
		}
	}

	//move computing result into matrix3
	for row := 0; row < new_n1; row++ {
		for col := 0; col < new_n4; col++ {
			matrix3[row][col] = c11[row][col]
			matrix3[row][col+new_n4] = c12[row][col]
			matrix3[row+new_n1][col] = c21[row][col]
			matrix3[row+new_n1][col+new_n4] = c22[row][col]
		}
	}

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

//矩陣相乘 (p1 ~ p7)
func m_m1(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p1[row][col] += add1[row][i] * add2[i][col]
			}
		}
	}
	ch <- 1
}

func m_m2(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p2[row][col] += add3[row][i] * b11[i][col]
			}
		}
	}
	ch <- 1
}
func m_m3(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p3[row][col] += a11[row][i] * add4[i][col]
			}
		}
	}
	ch <- 1
}
func m_m4(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p4[row][col] += a22[row][i] * add5[i][col]
			}
		}
	}
	ch <- 1
}
func m_m5(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p5[row][col] += add6[row][i] * b22[i][col]
			}
		}
	}
	ch <- 1
}
func m_m6(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p6[row][col] += add7[row][i] * add8[i][col]
			}
		}
	}
	ch <- 1
}
func m_m7(n1 int, n4 int, n2 int) {
	for row := 0; row < n1; row++ {
		for i := 0; i < n2; i++ {
			for col := 0; col < n4; col++ {
				p7[row][col] += add9[row][i] * add10[i][col]
			}
		}
	}
	ch <- 1
}
