package service

import (
	"golang-test/com/basic/util/http"
	"fmt"

	"encoding/json"
	"github.com/go-xorm/xorm"
	"golang-test/com/persist/model"
	_ "github.com/go-sql-driver/mysql"

	"time"
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


func CallWeather(regionCode string) string{
	httpUrl := "http://www.weather.com.cn/data/sk/"+regionCode+".html"
	var engine *xorm.Engine
    	engine, err := xorm.NewEngine("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		fmt.Println("err==", err)
	}
	var key int
	weatherInfo := &model.WeatherInfo{Id:key,CallStatus:"INIT",Region:regionCode,ResponseContent:"",CreateTime:time.Now(),UpdateTime:time.Now()}
	//affected, err := engine.Insert(&model.WeatherInfo{CallStatus:"INIT",Region:regionCode,ResponseContent:"",CreateTime:time.Now(),UpdateTime:time.Now()})
	affected, err := engine.Insert(weatherInfo)
	fmt.Println("affected=", affected, "id=",weatherInfo.Id)
	result := common.HttpGet(httpUrl)
	affectedUpdate, err := engine.Update(&model.WeatherInfo{CallStatus:"SUCCESS",Region:regionCode,ResponseContent:result,CreateTime:time.Now(),UpdateTime:time.Now()})
	fmt.Println("affectedUpdate=", affectedUpdate)

	fmt.Println("result=", result)
	return result
}


func CallMobileAttribution(mobile string) string {
	httpUrl := "http://www.apifree.net/mobile/"+mobile+".xml"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
	return result
}

func CallIpAttribution(ip string) string {
	httpUrl := "http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip="+ip
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)

	return string(result)
}


func CallSchedule(applicationId string, column string, status string, httpUrl string, httpLoginUrl string) string{
	request := make(map[string]string)
	request["applicationId"] = applicationId
	request["column"] = column
	request["status"] = status

	json, err := json.Marshal(request)
	if err != nil {
	    fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"

	fmt.Println("json=======", string(json))
	result := common.HttpPostFormByCookie(request, httpUrl, httpLoginUrl)

	fmt.Println("result=", result)
	return result
}