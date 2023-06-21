package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Network struct {
	ServerService string `yaml:"ServerService"`
	GamesvrUrl    string `yaml:"GamesvrUrl"`
}

type Common struct {
	ConnTimeOut      int64  `yaml:"ConnTimeOut"`
	StartService     string `yaml:"StartService"`
	ServiceId        int32  `yaml:"ServiceId"`
	NetLog           bool   `yaml:"NetLog"`
	PoolSize         int    `yaml:"PoolSize"`
	XXTeaKey         string `yaml:"XXTeaKey"`
	MaxPackPerMinute uint32 `yaml:"MaxPackPerMinute"`
}

type Config struct {
	Network Network `yaml:"Network"`
	Common  Common  `yaml:"Common"`
}

var Conf Config

func InitWithFile(configFile string) {
	//configFile := "../conf/service.yaml"

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("ReadFile %s failed. Err: %#v", configFile, err))
	}

	Conf = Config{}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(fmt.Sprintf("yaml.Unmarshal %s failed. Err: %#v", configFile, err))
	}

	/*if !snowflake.Init(Conf.Common.StartService, int(Conf.Common.ServiceId)) {
		panic(fmt.Sprintf("init snowflake failed! StartService: %s, ServiceId: %d", Conf.Common.StartService, Conf.Common.ServiceId))
	}*/

	log.Debug("init %s success, config: %#v", configFile, Conf)
}

func GetServiceId() int32 {
	return Conf.Common.ServiceId
}

func GetSelfConnIdx() uint64 {
	/*iType := ss.ServiceType_GatewaySvr
	idx := GetServiceId()
	// return (((uint64(iType) & 0xff) << 16) | (uint64(idx) & 0xffff))
	return utils.ServiceConnIdx(iType, idx)*/
	return 0

}
