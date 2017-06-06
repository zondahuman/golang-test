package test

import (
	"testing"
	"fmt"
	"container/list"
	"reflect"
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

func Test_Struct1(t *testing.T){
	list :=list.New()
	userBean1 := &UserBean{Id:1, Name:"lee1"}
	list.PushBack(userBean1)
	userBean2 := &UserBean{Id:2, Name:"lee2"}
	list.PushBack(userBean2)
	for elements := list.Front(); elements != nil;elements = elements.Next(){
		fmt.Println("elements=", elements.Value)
	}
}

func Test_Struct2(t *testing.T){
	userBean1 := &UserBean{}
	if isEmpty(userBean1){
		fmt.Println("userBean1 is null")
	}else{
		fmt.Println("userBean1 is  ", userBean1)
	}
	userBean2 := &UserBean{Id:2, Name:"lee2"}
	if userBean2 == nil{
		fmt.Println("userBean2 is null")
	}else{
		fmt.Println("userBean2 is  ", userBean2)
	}
}


type UserBean struct{
	Id int "id"
	Name string "name"
}


func isEmpty(a interface{}) bool {
    v := reflect.ValueOf(a)
    if v.Kind() == reflect.Ptr {
        v=v.Elem()
    }
    return v.Interface() == reflect.Zero(v.Type()).Interface()
}