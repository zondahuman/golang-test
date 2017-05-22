
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











