package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost  string
	DBName  string
	DBUser  string
	DBPass  string
	LOGFile string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") //<--実行元が変わったときに対応できないので、相対Pathで指定すべき
	viper.SetConfigType("ini")
	viper.AddConfigPath(".") //ここまでで、config.iniというファイルを指定
	viper.AutomaticEnv()     //環境変数を優先するように指示

    err := viper.ReadInConfig()
    if err != nil {
        // 設定ファイル(config.ini)が見つからない場合は無視
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("設定ファイル読込失敗 %w", err)
        }
    }

	//上記viperで読込んだ設定情報をconfig構造体にマッピング

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルマッピング失敗 %w", err)
	}

	//マッピング成功
	return &config, nil
}
