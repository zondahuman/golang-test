package service

import (
	"golang-test/com/basic/util/http"
	"fmt"
)
func Call(uid string, name string, id_card string, school_name string, mobile string) string{
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

func main() {
	uid := "2011"
	name := "许丽菲"
	id_card := "640121199805251120"
	school_name := "北京科技大学"
	mobile := "15226201503"

	Call(uid,name,id_card,school_name,mobile)
	//result := call(uid,name,id_card,school_name,mobile)
	//fmt.Println("result=", result)
}
