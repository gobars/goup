package util

// Settings 表示初始化化MySQL集群所需要的参数结构
type Settings struct {
	Servers []Servers
	MySQLHa MySQLHa `toml:"mysql-ha"`
	KafkaHa KafkaHa `toml:"kafka-ha"`
	RedisHa RedisHa `toml:"redis-ha"`
}

type Servers struct {
	IP       string `toml:"ip"`
	Iface    string `toml:"iface" default:"eth0"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type KafkaHa struct {
	Brokers []string `toml:"Brokers"`
	Port    int      `toml:"Port" default:"9092"`
}

type RedisHa struct {
	MasterAddr string `toml:"MasterAddr"`
	SlaveAddr  string `toml:"SlaveAddr"`
	Port       int    `toml:"Port" default:"6379"`
	Password   string `toml:"Password"`
	VipAddr    string `toml:"VipAddr"`
	IfaceName  string `toml:"IfaceName"`
}

type MySQLHa struct {
	Master1Addr  string `validate:"empty=false"`      // Master1的地址(IP，域名)
	Master2Addr  string `validate:"empty=false"`      // Master2的地址(IP，域名)
	User         string `default:"root"`              // Root用户名
	Password     string `validate:"empty=false"`      // Root用户密码
	Host         string `default:"127.0.0.1"`         // MySQL 端口号
	Port         int    `default:"3306"`              // MySQL 端口号
	ReplUsr      string `default:"repl"`              // 复制用用户名
	ReplPassword string `default:"984d-CE5679F93918"` // 复制用户密码
	VipAddr      string `validate:"empty=false"`      // 负载均衡IP地址
	IfaceName    string `default:"eth0"`              // 网卡名称
}
