package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/hermanowiczpiotr/ola/user/Interface"
	"github.com/hermanowiczpiotr/ola/user/application"
	"github.com/hermanowiczpiotr/ola/user/application/command"
	"github.com/hermanowiczpiotr/ola/user/application/query"
	"github.com/hermanowiczpiotr/ola/user/infrastructure/persistence"
	server "github.com/hermanowiczpiotr/ola/user/infrastructure/server"
	"log"
	"net/http"
)

func main() {
	log.Print("starting server")

	services := persistence.NewRepositories()
	services.AutoMigrate()
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	server.StartHttpServer(
		func(router chi.Router) http.Handler {
			return server.HandlerFromMux(
				Interface.NewUsersHandler(
					application.UserApp{
						query.NewGetUserByIdQuery(services.User),
						query.NewGetUserByEmailQuery(services.User),
						command.NewAddUserCommand(services.User),
					},
					tokenAuth,
				),
				router,
			)
		})
}
