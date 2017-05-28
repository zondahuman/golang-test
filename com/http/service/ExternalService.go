package service

import (
	"golang-test/com/basic/util/http"
	"fmt"

	"encoding/json"
)
func CallStudent(uid string, name string, id_card string, school_name string, mobile string) string{
	httpUrl := "http://python.loan.com/rules/is_student"
	request := make(map[string]string)
	request["uid"] = uid
	request["name"] = name
	request["id_card"] = id_card
	request["school_name"] = school_name
	request["mobile"] = mobile

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"
	result := common.HttpPost(request, header, httpUrl)

	fmt.Println("result=", result)
	return result
}

func CallLendHistory(uid string, phoneList []string, idCardList []string) string{
	httpUrl := "http://python.loan.com/rules/loan_history/v2"
	request := make(map[string]interface{})
	request["uid"] = uid

	request["phones"] = phoneList

	request["id_cards"] = idCardList

	json, err := json.Marshal(request)
	if err != nil {
	    fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"

	fmt.Println("json=======", json)
	fmt.Println("json=======", string(json))
	//result := common.HttpPostJson(json, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)

	fmt.Println("result=", result)
	return result
}
