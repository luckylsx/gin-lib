package conf

import (
	"errors"

	"gin-lib/conf/setting"
)

type MyMap map[string]map[string]string

var mysqlConf = map[string]MyMap{
	"dsp-business-pool": {
		"read": {
			"host":     setting.Conf().DspBusinessPoolRead.Host,
			"password": setting.Conf().DspBusinessPoolRead.Password,
			"port":     setting.Conf().DspBusinessPoolRead.Port,
			"user":     setting.Conf().DspBusinessPoolRead.User,
			"dbname":   setting.Conf().DspBusinessPoolRead.DBName,
		},
	},
	"dsp-pool": {
		"read": {
			"host":     setting.Conf().DspPoolRead.Host,
			"password": setting.Conf().DspPoolRead.Password,
			"port":     setting.Conf().DspPoolRead.Port,
			"user":     setting.Conf().DspPoolRead.User,
			"dbname":   setting.Conf().DspPoolRead.DBName,
		},
		"write": {
			"host":     setting.Conf().DspPoolWrite.Host,
			"password": setting.Conf().DspPoolWrite.Password,
			"port":     setting.Conf().DspPoolWrite.Port,
			"user":     setting.Conf().DspPoolWrite.User,
			"dbname":   setting.Conf().DspPoolWrite.DBName,
		},
	},
	"dsp-pool-record": {
		"read": {
			"host":     setting.Conf().DspRecordRead.Host,
			"password": setting.Conf().DspRecordRead.Password,
			"port":     setting.Conf().DspRecordRead.Port,
			"user":     setting.Conf().DspRecordRead.User,
			"dbname":   setting.Conf().DspRecordRead.DBName,
		},
		"write": {
			"host":     setting.Conf().DspRecordWrite.Host,
			"password": setting.Conf().DspRecordWrite.Password,
			"port":     setting.Conf().DspRecordWrite.Port,
			"user":     setting.Conf().DspRecordWrite.User,
			"dbname":   setting.Conf().DspRecordWrite.DBName,
		},
	},
	"default": {
		"read": {
			"host":     "192.168.1.201",
			"password": "devuser",
			"port":     "3306",
			"user":     "devuser",
			"dbname":   "dsp_manager",
		},
	},
}

// GetDbConfig get the config message of mysql
func GetAllDbConfig() map[string]MyMap {
	return mysqlConf
}

// GetDbConfig 获取某个连接配置
func GetDbConfig(key string) (MyMap, error) {
	if key == "" {
		key = "default"
	}
	if conf, ok := mysqlConf[key]; ok {
		return conf, nil
	} else {
		return nil, errors.New("not exists")
	}
}
