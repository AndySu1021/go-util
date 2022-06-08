package etcd

type Config struct {
	Endpoints   []string `mapstructure:"endpoints"`
	DialTimeout int      `mapstructure:"dial_timeout"`
	Username    string   `mapstructure:"username"`
	Password    string   `mapstructure:"password"`
}
