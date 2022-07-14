package delivery

import (
	"fmt"

	"livecode-lopei-grpc-client/config"
	"livecode-lopei-grpc-client/manager"
)

type cli struct {
	useCaseManager manager.UseCaseManager
}

func (c *cli) Run() {
	balance, err := c.useCaseManager.CheckBalanceUseCase().GetBalance(int32(3))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}
func Cli() *cli {
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	return &cli{
		useCaseManager: useCaseManager,
	}
}
