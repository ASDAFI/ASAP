package controllers

import (
	"context"
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_farm "farm/src/proto/messages/farm"
	"github.com/golang/protobuf/ptypes/empty"
)

func (f FarmServer) GetFarm(ctx context.Context, empty *empty.Empty) (*pb_farm.Farm, error) {
	userId := ctx.Value("user_id").(uint)

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQhandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}

	user, err := userQhandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)
	farmQhandler := farm.NewFarmQueryHandler(farmRepo)

	farmQuery := farm.GetFarmByIdQuery{FarmId: user.FarmID}

	fetchedFarm, err := farmQhandler.GetFarmById(ctx, farmQuery)
	if err != nil {
		return nil, err
	}

	return &pb_farm.Farm{
		Id:         uint32(fetchedFarm.ID),
		Name:       fetchedFarm.FarmName,
		Production: fetchedFarm.Production,
	}, nil
}
