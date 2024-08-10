package options

import (
	"github.com/gin-gonic/gin"

	"demo/internal/config"
	"demo/internal/pkg/database"
	"demo/internal/pkg/middleware/log"
	"demo/internal/server"
)

const (
	_defaultConfigFile = "config/config.yaml"
)

type AppOptions struct {
	ConfFile string
	Config   *config.Config
}

func NewAppOptions() *AppOptions {
	o := &AppOptions{}
	return o
}

func (o *AppOptions) NewServer() (*server.Server, error) {
	s := server.NewServer()
	o.loadConfig(o.ConfFile)
	s.Config = o.Config

	gin.SetMode(s.Config.Server.Mode)
	s.Gin = gin.New()
	
	s.Log = log.NewClient(o.Config.Log.InfoFilePath, o.Config.Log.ErrorFilePath, o.Config.Log.Level,
		o.Config.Log.MaxSize, o.Config.Log.MaxBackups, o.Config.Log.MaxAge,
	)

	Connect := o.Config.DB.Mysql.Username + ":" + o.Config.DB.Mysql.Password + "@tcp(" +
		o.Config.DB.Mysql.Host + ":" + o.Config.DB.Mysql.Port + ")/" + o.Config.DB.Mysql.DbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	s.DB = database.New(o.Config.DB.Type, Connect, database.WithMigrate(true))
	return s, nil
}

func (o *AppOptions) loadConfig(configFile string) {
	if configFile == "" {
		configFile = _defaultConfigFile
	}
	o.Config = config.New(configFile)
	if o.Config.Log.Path == "" {
		o.Config.Log.Path = "./log"
	}
	o.Config.Log.ErrorFilePath = o.Config.Log.Path + "/" + o.Config.Log.ErrorFilename
	o.Config.Log.InfoFilePath = o.Config.Log.Path + "/" + o.Config.Log.InfoFilename
}
