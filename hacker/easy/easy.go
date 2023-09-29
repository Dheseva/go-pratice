package hacker

import "fmt"

func Matrix() {
    // Define a two-dimensional slice (matrix)
    matrix := [][]int32{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // Get the diagonal values
    var diagonal []int32

    for i := 0; i < len(matrix); i++ {
        diagonal = append(diagonal, matrix[i][i])
		diagonal = append(diagonal, matrix[i][len(matrix)-1-i])
    }

    // Print the diagonal values
    fmt.Println("Diagonal values:", diagonal)
}

func Staircase() {
	height := 6
    print := "#"
    for i := 1; i <= int(height); i++{
        for sp := 1; sp <= int(height)-i; sp++{
            fmt.Print(" ")
        }
        for p := 1; p <= i; p++{
            fmt.Print(print)
        }
		fmt.Println()
    }
        
}

func miniMaxSum(arr []int32) {
    // Write your code here
    
    vaeri := arr
    var allresult []int32
    var minmax []int32
    
    for i := 0; i < len(vaeri); i++{
            var sum int32
            for in, number := range vaeri {
                if in != i {
                sum += number
                }
            }
            allresult = append(allresult, sum)
    }
    min := allresult[0] // slices.min(arr)
    max := allresult[0] // slices.max(arr)
     for _, num := range allresult {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    minmax = append(minmax, min)
    minmax = append(minmax, max)
    
    fmt.Println(min,max)

}

func birthdayCakeCandles(candles []int32) int32 {
    // Jumlah lilin paling tinggi
    lilin := candles
    maks := lilin[0]
    var total int32
    for _, v := range lilin{
        if v > maks{
            maks = v
        }
    }
    for _, v2 := range lilin{
        if v2 == maks{
            total = total + 1
        }
    }
    return total
}