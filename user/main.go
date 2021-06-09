package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	db "github.com/micro/services/db/proto"
	"github.com/micro/services/user/handler"
	proto "github.com/micro/services/user/proto"
)

func main() {
	service := service.New(
		service.Name("user1"),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), handler.NewUser(db.NewDbService("db", service.Client())))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
