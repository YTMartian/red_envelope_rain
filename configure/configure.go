/*
配置文件
*/
package configure

const (
	MaxSnatch          = 10        //最多抢红包数
	TotalMoney         = 100000000 //红包总金额
	TotalAmount        = 1000      //红包总个数
	MinMoney           = 1
	MaxMoney           = 2*TotalMoney/TotalAmount - MinMoney //这样就尽量做到把钱用完
	MysqlAddr          = "127.0.0.1:3306"
	MysqlUser          = "root"
	MysqlPass          = "root"
	MysqlDatabase      = "red_envelope"
	MachineId          = 1 //分布式机器编号
	RedisAddr          = "127.0.0.1:6379"
	RedisPass          = ""
	JSMAXNUM           = 1 << 53 //js整数的取值范围在[-2^53,2^53]
	ReadTimeout        = 60
	WriteTimeout       = 60
	LISTEN_PORT        = ":8080"
	GIN_MODE           = "debug" //debug or release
	RMQ_NAMESEVER_ADDR = "127.0.0.1:9876"
	RMQ_BROKER_ADDR    = "127.0.0.1:10911"
)
