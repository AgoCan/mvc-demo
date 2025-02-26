package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

    "demo/internal/config"
	"demo/internal/pkg/database"
	"demo/internal/pkg/middleware/cors"
	"demo/internal/pkg/middleware/log"
)

type Server struct {
	Config *config.Config
	Gin    *gin.Engine
	// 导入日志
	Log *log.Client
	DB        database.DB
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	c := cors.New(cors.Config{
		AllowOrigins:     s.Config.Server.Origins,
		AllowMethods:     s.Config.Server.Methods,
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})

	s.Gin.Use(log.GinLogger(s.Log.Logger),
		log.GinRecovery(s.Log.Logger, true),
		c)

	s.migrate()
	
	s.SetupRouter()

	err := s.Gin.Run(s.Config.Server.Bind + ":" + s.Config.Server.Port)
	if err != nil {
		fmt.Println(err)
	}
}
