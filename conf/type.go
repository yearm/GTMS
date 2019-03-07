package conf

type MySQlConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}
