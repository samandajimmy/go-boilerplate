package cmdutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"
)

type EchoGroup struct {
	Api   *echo.Group
	Token *echo.Group
}

func LoadEnv() {
	envPath := ".env"

	if os.Getenv(`APP_PATH`) != "" {
		envPath = os.Getenv(`APP_PATH`) + "/" + envPath
	}

	_ = godotenv.Load(envPath)
}

func LoadTestData() {
	dataPath := "test/data"

	if os.Getenv(`APP_PATH`) != "" {
		dataPath = os.Getenv(`APP_PATH`) + "/" + dataPath
	}

	items, err := ioutil.ReadDir(dataPath)

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	if len(items) == 1 {
		return
	}

	viper.AddConfigPath(dataPath)
	_ = viper.ReadInConfig()

	mockDataFile := dataPath
	for _, item := range items {
		if item.IsDir() {
			// currently we does not expect any dir inside
			continue
		}

		// load all existed yaml file
		mockDataFile = mockDataFile + "/" + item.Name()
		viper.SetConfigName(strings.TrimSuffix(filepath.Base(item.Name()), filepath.Ext(item.Name())))
		viper.AddConfigPath(mockDataFile)
		err = viper.MergeInConfig()

		if err != nil {
			pgdlogger.Make().Fatal(err)
		}
	}
}
