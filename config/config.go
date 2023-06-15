package config

import (
	"fmt"
	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		RMQ  `yaml:"rabbitmq"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"url"      env:"PG_URL"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
		URL            string `env-required:"true" yaml:"url"                 env:"RMQ_URL"`
	}
)

func NewConfig() *Config {
	cfg := &Config{}
	cwd := projectRoot()
	envFilePath := cwd + ".env"
	envFilePath, errRunfile := bazel.Runfile(".env")
	if errRunfile != nil {
		panic(errRunfile)
	}

	fmt.Printf("###READ ENV %s ####\n", envFilePath)

	contents, errReadFile := ioutil.ReadFile(envFilePath)
	if errReadFile != nil {
		panic(errReadFile)
	}
	fmt.Printf("Contents: %s", contents)

	err := readEnv(envFilePath, cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func readEnv(envFilePath string, cfg *Config) error {
	envFileExists := checkFileExists(envFilePath)

	if envFileExists {
		err := cleanenv.ReadConfig(envFilePath, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	} else {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {

			if _, statErr := os.Stat(envFilePath + ".example"); statErr == nil {
				return fmt.Errorf("missing environmentvariables: %w\n\nprovide all required environment variables or rename and update .env.example to .env for convinience", err)
			}

			return err
		}
	}

	fmt.Printf("Read env: %v\n", cfg)

	return nil
}

func checkFileExists(fileName string) bool {
	envFileExists := false
	if _, err := os.Stat(fileName); err == nil {
		envFileExists = true
	}

	folder := filepath.Dir(fileName)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Files in folder %s:", folder)
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	folder = folder + "/__main__"

	fmt.Printf("Files in parent folder %s:", folder)
	files, err = ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	return envFileExists
}

func projectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(b)

	return projectRoot + "/../"
}
