package app

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
	"github.com/Rasikrr/learning_platform_users/internal/ports/grpc"
	enrollmentsR "github.com/Rasikrr/learning_platform_users/internal/repositories/enrollments"
	usersR "github.com/Rasikrr/learning_platform_users/internal/repositories/users"
	enrollmentsS "github.com/Rasikrr/learning_platform_users/internal/services/enrollments"
	usersS "github.com/Rasikrr/learning_platform_users/internal/services/users"
)

type App struct {
	*application.App
	usersRepository       usersR.Repository
	usersService          usersS.Service
	enrollmentsService    enrollmentsS.Service
	enrollmentsRepository enrollmentsR.Repository
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
		a.initRepositories,
		a.initServices,
		a.initClients,
		a.initGRPCServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initRepositories(_ context.Context) error {
	a.usersRepository = usersR.NewRepository(a.Postgres())
	a.enrollmentsRepository = enrollmentsR.NewRepository(a.Postgres())
	return nil
}

func (a *App) initServices(_ context.Context) error {
	a.usersService = usersS.NewService(a.usersRepository)
	a.enrollmentsService = enrollmentsS.NewService(a.enrollmentsRepository)
	return nil
}

func (a *App) initClients(ctx context.Context) error {
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	grpc.NewServer(
		a.GrpcServer().Srv(),
		a.usersService,
		a.enrollmentsService,
	)
	return nil
}
