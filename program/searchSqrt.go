package main

import "fmt"

func main(){
	x:=6
	result := searchSqrt(x)
	fmt.Println(result)
}

func searchSqrt(x int) int{
	left,right:=1,x
	for left<=right {
		mid:=(left+right)/2
		if mid==x/mid && x%mid==0{
			return mid
		}else if mid>x/mid{
			right=mid-1
		}else if mid<x/mid{
			left=mid+1
		}
	}
	fmt.Println(left,right)
	return right
}

