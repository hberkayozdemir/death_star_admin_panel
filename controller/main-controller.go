package controller

import (
	"github.com/hberkayozdemir/death_star_admin_panel/internal/apps"
	pb "github.com/hberkayozdemir/death_star_admin_panel/internal/protorender"
	"google.golang.org/grpc"
)

type Controller struct {
	App *apps.Server
}

func (c *Controller) NewControllerSet() {
	c.App = &apps.Server{}
	c.App.NewServer()
}

func (c *Controller) NewController() *grpc.Server {
	c.NewControllerSet()

	grpcServer := grpc.NewServer()

	// society.RegisterSocietyServiceServer(grpcServer, c.Society)
	pb.RegisterSocietyServiceServer(grpcServer, c.App)

	return grpcServer
}
