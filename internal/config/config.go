package config

import (
	"github.com/huerni/gmitex/pkg/config"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Prefix string          `json:"prefix"`
	Grpc   config.RpcConf  `json:"grpc"`
	Http   config.HttpConf `json:"http"`

	Etcd       config.EtcdConf    `json:"etcd,optional"`
	Mysql      config.MysqlConf   `json:"mysql,optional"`
	MysqlSlave config.MysqlConf   `json:"mysql-slave,optional"`
	Traefik    config.TraefikConf `json:"traefik,optional"`
}

var (
	Cfg *Config
)

func InitConfig(filePath string) (*Config, error) {
	if Cfg == nil {
		Cfg = &Config{}
		conf.MustLoad(filePath, Cfg)
		err := FigureConf(Cfg)
		if err != nil {
			return nil, err
		}
	}

	return Cfg, nil
}

func GetConfig() *Config {
	return Cfg
}

func FigureConf(c *Config) error {
	err := c.Grpc.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Http.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Etcd.FigureConfig()
	if err != nil {
		return err
	}

	err = c.Mysql.FigureConfig()
	if err != nil {
		return err
	}
	if c.Mysql.DSN == "" && c.Mysql.HasConfig() {
		err := config.GetFigureFromEtcd(c.Prefix, c.Etcd.Hosts, &c.Mysql)
		if err != nil {
			return err
		}
	}

	err = c.MysqlSlave.FigureConfig()
	if err != nil {
		return err
	}
	if c.MysqlSlave.DSN == "" && c.MysqlSlave.HasConfig() {
		err := config.GetFigureFromEtcd(c.Prefix, c.Etcd.Hosts, &c.MysqlSlave)
		if err != nil {
			return err
		}
	}

	err = c.Traefik.FigureConfig()
	if err != nil {
		return err
	}

	return nil
}
