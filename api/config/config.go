package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost  string	`mapstructure:"DB_HOST"`
	DBName  string	`mapstructure:"DB_NAME"`
	DBUser  string	`mapstructure:"DB_USER"`
	DBPass  string	`mapstructure:"DB_PASSWORD"`
	LOGFile string	`mapstructure:"LOGFILE"` 
}

func LoadConfig() (*Config, error) {
	// 1. 環境変数を優先して設定
    viper.AutomaticEnv()
	viper.SetConfigName("config") //<--実行元が変わったときに対応できないので、相対Pathで指定すべき
	viper.SetConfigType("ini")
	viper.AddConfigPath(".") //ここまでで、config.iniというファイルを指定

    err := viper.ReadInConfig()
    if err != nil {
        // 設定ファイル(config.ini)が見つからない場合は無視
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("設定ファイル読込失敗 %w", err)
        }
		// config.iniが見つからない場合は環境変数のみ使用
    }

	//上記viperで読込んだ設定情報をconfig構造体にマッピング
	var config Config
	// 明示的に環境変数から取得（確実に取得するため）
    config.DBHost = viper.GetString("DB_HOST")
    config.DBName = viper.GetString("DB_NAME")
    config.DBUser = viper.GetString("DB_USER")
    config.DBPass = viper.GetString("DB_PASSWORD")
    config.LOGFile = viper.GetString("LOGFILE")

	//マッピング成功
	return &config, nil
}
