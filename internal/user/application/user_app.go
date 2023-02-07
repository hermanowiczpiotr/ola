package application

import (
	"github.com/hermanowiczpiotr/ola/user/application/command"
	query2 "github.com/hermanowiczpiotr/ola/user/application/query"
)

type UserApp struct {
	GetUserByIdQueryHandler    query2.GetUserByIdQueryHandler
	GetUserByEmailQueryHandler query2.GetUserByEmailQueryHandler
	AddUserCommandHandler      command.AddUserCommandHandler
}
