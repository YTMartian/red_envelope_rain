/*
配置文件
*/
package configure

const (
	MaxSnatch     = 10        //最多抢红包数
	TotalMoney    = 100000000 //红包总金额
	TotalAmount   = 100  //红包总个数
	MinMoney      = 1
	MaxMoney      = 100
	MysqlAddr     = "localhost"
	MysqlPort     = 3306
	MysqlUser     = "root"
	MysqlPass     = "root"
	MysqlDatabase = "red_envelope"
	MachineId     = 1 //分布式机器编号
	RedisAdder    = "localhost"
	RedisPort     = 6379
	RedisPass     = ""
)
