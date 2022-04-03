package bootstrap

import (
	"alidrive_uploader/conf"
	"github.com/sirupsen/logrus"
	"math"
)

func InitConfig() {
	//vipConfig := viper.New()
	conf.VipConfig.SetConfigFile(conf.Opt.Config) // 指定配置文件路径
	// 查找并读取配置文件
	if err := conf.VipConfig.ReadInConfig(); err != nil {
		logrus.Fatalf("读取配置出错: %s \n", err)
	}
	if err := conf.VipConfig.Unmarshal(&conf.Conf); err != nil {
		logrus.Fatalf("解析配置出错: %s \n", err)
	}
	if conf.Opt.Debug == nil {
		conf.Opt.Debug = &conf.Conf.Debug
	}

	if conf.Opt.Transfers == nil {

		conf.Opt.Transfers = &conf.Conf.Transfers
	}
	//最小任务数 1
	*conf.Opt.Transfers = uint64(math.Max(float64(*conf.Opt.Transfers), 1))
}