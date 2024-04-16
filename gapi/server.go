package gapi

import (
	"fmt"

	db "github.com/yukunzhan/simplebank/db/sqlc"
	"github.com/yukunzhan/simplebank/pb"
	"github.com/yukunzhan/simplebank/token"
	"github.com/yukunzhan/simplebank/util"
	"github.com/yukunzhan/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
