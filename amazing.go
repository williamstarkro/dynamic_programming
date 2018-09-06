package main

import (
    "fmt"
    "strconv"
)

func main() {

    // fmt.Println("Enter number: ")
    // text := ""
    // fmt.Scanln(text)
    
    // val, err := strconv.ParseInt(text, 10, 64)
    // if err == nil {
    	var value int64 = 150
        res := amazingNumber(value)
        fmt.Println(strconv.FormatInt(res,10))
    // }
}

func amazingNumber(val int64) int64 {

	// for an amazing number, you have really 7 possible products of the original 3 primes
	// 2*2(n), 3*3(n), 7*7(n), 2*3(n), 2*7(n), 3*7(n), and 2*3*7(n) 

	var AN = make([]int64, val) // To store amazing numbers
	// we need index variables that can point where the next number is going to be
    i2 := 0
    i3 := 0
    i7 := 0
    // storing what iteration of multiple we are getting to
    var multiples_2 int64 = 2
    var multiples_3 int64 = 3
    var multiples_7 int64 = 7
    // tracker at the end
    var amazing int64 = 1
 	
 	// initializer
    AN[0] = 1
    // iterator from 
    var i int64
    for i = 1; i<val; i++ {
    	// to get the next value, we need to take all the multiples and see
    	// which one is next in line (min val)
    	// once we figure out which one in particular, we can then find the next
    	// multiple for that specific amazing number
    	// e.g. multiple_2 is min (so we add 1 to our tracker i2, and multiply multiple i2 by 2)
    	if multiples_7 < multiples_3 {
    		if multiples_7 < multiples_2 {
    			amazing = multiples_7
    		} else {
    			amazing = multiples_2
    		}
    	} else {
    		if multiples_3 < multiples_2 {
    			amazing = multiples_3
    		} else {
    			amazing = multiples_2
    		}
    	}
    	// let's update our list with the new amazing number
    	AN[i] = amazing;

    	if amazing == multiples_2 {
    		// next iteration of 2 multiples
    		i2 = i2+1
    		// get the next multiple of 2
           	multiples_2 = AN[i2]*2
    	}
    	if amazing == multiples_3 {
    		i3 = i3+1
           	multiples_3 = AN[i3]*3
    	}
    	if amazing == multiples_7 {
    		i7 = i7+1
           	multiples_7 = AN[i7]*7
    	}
    	//fmt.Println(strconv.FormatInt(AN[i],10))
    } 
    // at this point we should have calculated the next amazing number
    return amazing;

}