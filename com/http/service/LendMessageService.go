package service

import (
	"golang-test/com/basic/util/http"
	"fmt"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"

)




func CallMessageConfirm(applicationId string, httpUrl string, httpLoginUrl string) string{
	request := make(map[string]string)
	request["applicationId"] = applicationId

	json, err := json.Marshal(request)
	if err != nil {
	    fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["source"] = "LOAN"

	fmt.Println("json=======", string(json))
	result := common.HttpPostFormByCookie(request, header, httpUrl, httpLoginUrl)

	fmt.Println("result=", result)
	return result
}















