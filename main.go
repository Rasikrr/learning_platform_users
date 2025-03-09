package main

import (
	"context"
	"github.com/Rasikrr/learning_platform_users/internal/app"
)

const (
	appName = "users"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx, appName)
	if err != nil {
		panic(err)
	}
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
