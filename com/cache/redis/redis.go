package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golang-test/com/cache/enums"
)

//func getClient() interface{}{
func getClient() redis.Conn {
	// 从配置文件获取redis配置并连接
	host := "172.16.2.145:6379"
	db := 0
	// tcp连接redis
	rs, err := redis.Dial("tcp", host)
	// 操作完后自动关闭
	//defer rs.Close()
	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		//return
	}
	// 选择db
	rs.Do("SELECT", db)
	return rs
}

func operateSetnx(key string, value string) int {
	// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
	// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	client := getClient()
	n, err := client.Do("SETNX", key, value)
	// 若操作失败则返回
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println("n=====", n)
	success := enums.INIT
	// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
	if n == int64(1) {
		// 设置过期时间为24小时
		n, _ := client.Do("EXPIRE", key, 24 * 3600)
		if n == int64(1) {
			fmt.Println("success")
			success = enums.SUCCESS

		}
	} else if n == int64(0) {
		fmt.Println("the key has already existed")
		success = enums.EXISTS
	}else {
		success = enums.UNKNOWN
	}
	return success
}

func operateSet(key string, value string) {
	client := getClient()
	fmt.Println("client=", client)
	value, err := client.Do("SET", key, value)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("value=", value)

}

func operateGet(key string) {
	client := getClient()
	fmt.Println("client=", client)

	var value, err = redis.String(client.Do("GET", key))
	if err != nil {
		fmt.Println("err=", err)
		//return
	}
	fmt.Println("value=", value)
}