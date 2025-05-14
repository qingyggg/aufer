package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func EnvInit() {
	mode := os.Getenv("FOO_ENV")

	// 根据环境模式加载对应的环境变量文件
	envFile := filepath.Join("./envs", "local.env")
	if "production" == mode {
		envFile = filepath.Join("./envs", "prod.env")
	}

	// 加载环境变量文件
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Warning: Could not load env file %s: %v", envFile, err)

		// 尝试使用绝对路径作为备选方案
		fallbackFile := filepath.Join("/app/envs", filepath.Base(envFile))
		err = godotenv.Load(fallbackFile)
		if err != nil {
			log.Printf("Warning: Could not load fallback env file %s: %v", fallbackFile, err)
		} else {
			log.Printf("Successfully loaded environment from fallback path %s", fallbackFile)
		}
	} else {
		log.Printf("Successfully loaded environment from %s", envFile)
	}
}
