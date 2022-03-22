package main

import (
	"server/infra"
	"server/middleware/cache"
	"server/middleware/errorhandler"
	"server/middleware/logger"
)

func main() {
	logger.Init()
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("caught panic: %v", err)
		}
		_ = logger.LoggerSync()
	}()
	conn, err := infra.Open()
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		if conn != nil {
			db, err := conn.DB()
			if err != nil {
				return
			}
			_ = db.Close()
		}
	}()
	// connect redis
	err = cache.InitClient()
	if err != nil {
		logger.Fatal(err)
	}
	r, err := infra.SetupServer(conn.Debug())
	if err != nil {
		logger.Error(err)
	}
	if err := r.Run(); err != nil {
		logger.Error(errorhandler.MsgErrServer, err)
	}
}
