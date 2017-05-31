package http

import (
	"testing"
	"fmt"
	"golang-test/com/http/service"
	"golang-test/com/http/model"
	"encoding/json"
	 _ "github.com/go-sql-driver/mysql"
)


func Test_CallStudent(t *testing.T) {
	//{"uid":"149619833737477","name":"马克龙","id_card":"110101198606250113","mobile":"13645836311","school_name":"加州大学伯克利分校"}
	uid := "149619833737477"
	name := "马克龙"
	id_card := "110101198606250113"
	school_name := "加州大学伯克利分校"
	mobile := "13645836311"

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

func Test_CallWeather(t *testing.T) {
	region := "101010100"
	var result string
	result = service.CallWeather(region)
	fmt.Println("result=======", result)
}

func Test_CallMobileAttribution(t *testing.T) {
	var mobile string
	mobile = "13645836311"
	var result string
	result = service.CallMobileAttribution(mobile)
	fmt.Println("result=======", result)
}

func Test_CallIpAttribution(t *testing.T) {
	var ip string
	ip = "117.89.35.58"
	var result string
	result = service.CallIpAttribution(ip)
	fmt.Println("result=======", result)

	var ipModel model.IpModel
	if err := json.Unmarshal([]byte(result), &ipModel); err == nil{
		fmt.Println("ipModel====", ipModel)
	}else{
		fmt.Println("json error ;", err)
	}

}

