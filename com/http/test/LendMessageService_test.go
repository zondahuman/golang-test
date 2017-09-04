package http

import (
	"testing"
	"fmt"
	"golang-test/com/http/service"
	"golang-test/com/http/model"
	"encoding/json"
	 _ "github.com/go-sql-driver/mysql"
)


func Test_CallMessageConfirm(t *testing.T) {
	//{"uid":"306","name":"崔孝范","id_card":"230405197608040640","mobile":"18500804170","school_name":"把株式会社"}
	applicationId := "RRD_a068fcbddbbc54d_9"
	httpUrl := "http://172.16.2.145:9100/thirdparty/sendConfirm"
	httpLoginUrl := "http://172.16.2.145:9100/rrd-verify-web/login/login?username=admin@heika.com&password=d1xqqOOj23Dve6AskgwkDtl8yZqVEKrBmMgpVhWxSl9yqG/ZyTJ6O/XfQMvXWRJvdvshRAxGLYGYzjzVoS9/5pF2/YU4Ptmyx4YPlZSDKtl90T5PpRqj4UCPqpMle2vm8Z5sO2BOi0g1EJR6CAN4cOX/mpl1TbCzcNHTHuHcl6U="
	result := service.CallMessageConfirm(applicationId, httpUrl, httpLoginUrl)
	fmt.Println("result====", result)

	//obj := model.StudentModel{}
	var obj model.LendAplication
	if err := json.Unmarshal([]byte(result), &obj); err == nil{
		fmt.Println("obj====", obj)
	}else{
		fmt.Println("json error", err)
	}

}


