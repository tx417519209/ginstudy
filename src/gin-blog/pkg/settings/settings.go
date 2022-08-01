package settings

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

var (
	Cfg           *viper.Viper
	RunMode       string
	HTTPPort      int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	PageSize      int
	JwtSecret     string
	DBType        string
	DBName        string
	DBUser        string
	DBPassWord    string
	DBHost        string
	DBTablePrefix string
)

func LoadBase() {
	RunMode = Cfg.GetString("system.RUN_MODE")
}
func LoadDb() {
	DBType = Cfg.GetString("database.TYPE")
	DBName = Cfg.GetString("database.NAME")
	DBUser = Cfg.GetString("database.USER")
	DBPassWord = Cfg.GetString("database.PASSWORD")
	DBHost = Cfg.GetString("database.HOST")
	DBTablePrefix = Cfg.GetString("database.TABLE_PREFIX")
}
func LoadServer() {
	HTTPPort = Cfg.GetInt("server.HTTP_PORT")
	ReadTimeout = time.Duration(Cfg.GetInt("server.READ_TIMEOUT"))
	WriteTimeout = time.Duration(Cfg.GetInt("server.WRITE_TIMEOUT"))
}
func LoadApp() {
	JwtSecret = Cfg.GetString("app.JWT_SECRET")
	PageSize = Cfg.GetInt("app.  PAGE_SIZE")
}

func init() {
	Cfg = viper.New()
	curDir, _ := os.Getwd()
	configPath := path.Join(curDir, "src", "gin-blog", "config", "config.yaml")
	Cfg.SetConfigFile(configPath)
	if err := Cfg.ReadInConfig(); err != nil {
		panic(err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadDb()
}
