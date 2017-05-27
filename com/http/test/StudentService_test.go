package http

import (
	"testing"
	"fmt"
	"golang-test/com/http/service"
	"golang-test/com/http/model"
	"encoding/json"
)


func Test_sendRabbitmq(t *testing.T) {
	//{"uid":"306","name":"崔孝范","id_card":"230405197608040640","mobile":"18500804170","school_name":"把株式会社"}
	uid := "306"
	name := "崔孝范"
	id_card := "230405197608040640"
	school_name := "把株式会社"
	mobile := "18500804170"

	//call(uid,name,id_card,school_name,mobile)
	result := service.Call(uid,name,id_card,school_name,mobile)
	fmt.Println("result====", result)

	//obj := model.StudentModel{}
	var obj model.StudentModel
	if err := json.Unmarshal([]byte(result), &obj); err == nil{
		fmt.Println("obj====", obj)
	}else{
		fmt.Println("json error", err)
	}

}