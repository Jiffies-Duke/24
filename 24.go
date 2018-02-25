package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var oper = [4]string{
	"+", "-", "*", "/",
}

func main() {
	var number [4]float64

	number = getInput()
	//fmt.Println("这 4 个数字是：=", number)
	for i := 0; i < 4; i++ {
		if number[i] <= 0 {
			fmt.Println("其中出现零或负数，退出")
			os.Exit(0)
		}
	}
	answer, _ := calc(24, number[0:4])
	if answer == "N" {
		answer = "无解"
	}
	fmt.Println("答案： " + answer)
}

func getInput() [4]float64 {
	var temp [4]float64
	fmt.Println("请输入 4 个 1 - 13 的正整数，空格分开：")
	fmt.Scanf("%f %f %f %f ", &temp[0], &temp[1], &temp[2], &temp[3])
	return temp
}

func calc(result float64, candidate []float64) (string, int) {
	//fmt.Println(result, candidate, len(candidate))
	if result < 0 {
		return "N", 0
	}
	if result-math.Floor(result) > 0.01 || math.Ceil(result)-result > 0.01 {
		return "N", 0
	}
	s := len(candidate)
	if s <= 0 {
		return "N", 0
	}
	if s == 1 {
		x := candidate[0]
		if x-result < 0.0001 && result-x < 0.0001 {
			return strconv.FormatFloat(x, 'f', 0, 64), 6
		}
		return "N", 0
	}

	for n := 0; n < s; n++ {
		used := []bool{false, false, false, false}
		priority := []int{0, 1, 2, 3, 1, 3, 6}
		x := candidate[n]
		used[n] = true
		a := initA(result, x)

		b := initB(candidate, used) //candidate[1:s]
		for i, v := range a {
			try, p := calc(v, b)
			//fmt.Println("try, priority, x, i = ", try, p, x, i)
			if try != "N" {
				left := ""
				right := ""
				if priority[i] > priority[p] || (priority[i]==priority[p] &&(i==1 || i==3)){
					left = "("
					right = ")"
				}
				if i < 4 {
					return strconv.FormatFloat(x, 'f', 0, 64) + " " + oper[i] + " " + left + try + right, i
				}
				i = i*2 - 7
				return left + try + right + " " + oper[i] + " " + strconv.FormatFloat(x, 'f', 0, 64), i
			}
		}
	}
	return "N", 0
}
func initB(candidate []float64, used []bool) []float64 {
	var b []float64
	for i := 0; i < len(candidate); i++ {
		if !used[i] {
			b = append(b, candidate[i])
		}
	}
	return b
}

func initA(result float64, x float64) []float64 {
	var a []float64
	a = append(a, result-x, x-result)
	if x != 0 {
		a = append(a, result/x)
	} else {
		a = append(a, -1)
	}

	if result != 0 {
		a = append(a, x/result)
	} else {
		a = append(a, -1)
	}
	a = append(a, result+x, x*result)
	return a
}
