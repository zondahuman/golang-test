package http

import (
	"testing"
	"fmt"
	"golang-test/com/http/service"
	"golang-test/com/http/model"
	"encoding/json"
)


func Test_CallStudent(t *testing.T) {
	//{"uid":"306","name":"崔孝范","id_card":"230405197608040640","mobile":"18500804170","school_name":"把株式会社"}
	uid := "306"
	name := "崔孝范"
	id_card := "230405197608040640"
	school_name := "把株式会社"
	mobile := "18500804170"

	//call(uid,name,id_card,school_name,mobile)
	result := service.CallStudent(uid,name,id_card,school_name,mobile)
	fmt.Println("result====", result)

	//obj := model.StudentModel{}
	var obj model.StudentModel
	if err := json.Unmarshal([]byte(result), &obj); err == nil{
		fmt.Println("obj====", obj)
	}else{
		fmt.Println("json error", err)
	}

}



func Test_CallLendHistory(t *testing.T) {
	//{"uid":"306","name":"崔孝范","id_card":"230405197608040640","mobile":"18500804170","school_name":"把株式会社"}
	uid := "306"
	var phoneList []string
	phoneList = append(phoneList, "18513622928")
	fmt.Println("phoneList=======", len(phoneList))

	jsonPhoneList, err := json.Marshal(phoneList)
	if err != nil {
	    fmt.Println("error:", err)
	}
	fmt.Println("jsonPhoneList=======", string(jsonPhoneList))

	var idCardList []string
	idCardList = append(idCardList, "230405197608040640")

	//call(uid,name,id_card,school_name,mobile)
	result := service.CallLendHistory(uid, phoneList, idCardList)
	fmt.Println("result====", result)

	//obj := model.StudentModel{}
	var obj model.LendAplication
	if err := json.Unmarshal([]byte(result), &obj); err == nil{
		fmt.Println("obj====", obj)
	}else{
		fmt.Println("json error", err)
	}

}



func Test_JsonTtransfer(t *testing.T) {
	//{"uid":"306","name":"崔孝范","id_card":"230405197608040640","mobile":"18500804170","school_name":"把株式会社"}
	uid := "306"
	var phoneList []string
	phoneList = append(phoneList, "18513622928")
	fmt.Println("phoneList=======", len(phoneList))

	jsonPhoneList, err := json.Marshal(phoneList)
	if err != nil {
	    fmt.Println("error:", err)
	}
	fmt.Println("jsonPhoneList=======", string(jsonPhoneList))

	var idCardList []string
	idCardList = append(idCardList, "230405197608040640")

	request := make(map[string]interface{})
	request["uid"] = uid

	request["phones"] = phoneList

	request["id_cards"] = idCardList

	json, err := json.Marshal(request)
	if err != nil {
	    fmt.Println("error:", err)
	}

	fmt.Println("json=======", json)
	fmt.Println("json=======", string(json))
}
