
CREATE TABLE `order_info` (
  `id` int(11) auto_increment primary key NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


go get github.com/go-sql-driver/mysql

go build golang-test/com/sql/mysql
go install golang-test/com/sql/mysql


go test -v golang-test/com/sql/mysql -run ^Test_insert$
go test -v golang-test/com/sql/mysql -run 'Test_insert'
go test -v golang-test/com/sql/mysql -run 'Test_findParam'





REDIS:
go get github.com/garyburd/redigo/redis

go build golang-test/com/cache/redis
go install golang-test/com/cache/redis




RABBITMQ:
go get github.com/streadway/amqp

go build golang-test/com/mq/rabbitmq
go install golang-test/com/mq/rabbitmq


Mongodb:
go get gopkg.in/mgo.v2


ORM:
https://github.com/xormplus/xorm
https://github.com/jinzhu/gorm


pycharm 2016 注册码
https://www.360kb.com/kb/2_24.html



create table weather_info (
id int auto_increment comment 'id',
region varchar(100) not null comment 'region',
call_status varchar(100)  not null comment 'call_status',
response_content text   not null comment 'response_content',
create_time datetime comment 'create_time',
update_time datetime comment 'update_time',
constraint pk_weather_info primary key(id)
)

库安装

go get github.com/go-xorm/xorm

Xorm工具

go get github.com/go-xorm/cmd/xorm


