package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lirubao/ozone/cache"
	"github.com/lirubao/ozone/model"
)

// Init 初始化配置
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
