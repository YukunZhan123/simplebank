package gapi

import (
	"fmt"

	db "github.com/yukunzhan/simplebank/db/sqlc"
	"github.com/yukunzhan/simplebank/pb"
	"github.com/yukunzhan/simplebank/token"
	"github.com/yukunzhan/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
