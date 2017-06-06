package test

import (
	"testing"
	"fmt"
)

func Test_Array1(t *testing.T){
	var array1 =[5] int{9,5,2}
	fmt.Println("array1=", array1)
	for index, value := range array1{
		fmt.Println("index=", index,", value=", value)
	}
	var array2 = [...]int{4,5,6}
	fmt.Println("array2=", array2[:])
	for i:=0;i<len(array2);i++{
		fmt.Println("array2[",i,"]=", array2[i])
	}
}


func Test_Splice1(t *testing.T){
	var splice1 =[] int{9,5,2}
	fmt.Println("splice1=", splice1)
	for index, value := range splice1{
		fmt.Println("index=", index,", value=", value)
	}
	var splice2 = make([]int, 2, 10)
	splice2 = append(splice2,4,5,6,7,8,9,1)
	fmt.Println("splice2=", splice2[:])
	for i:=0;i<len(splice2);i++{
		fmt.Println("splice2[",i,"]=", splice2[i])
	}
	fmt.Println("len(splice2)=", len(splice2))

}

