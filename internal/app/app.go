package app

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
	"github.com/Rasikrr/learning_platform_users/internal/ports/grpc"
)

type App struct {
	*application.App
}

func NewApp(ctx context.Context, name string) (*App, error) {
	app := &App{
		App: application.NewApp(ctx, name),
	}
	if err := app.Init(ctx); err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Init(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		a.initGRPCServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	grpc.NewServer(a.GrpcServer().Srv())
	return nil
}
