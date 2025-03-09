package main

import (
	"context"
	app2 "github.com/Rasikrr/learning_platform_users/internal/app"
)

const (
	name = "users"
)

func main() {
	ctx := context.Background()
	app, err := app2.NewApp(ctx, name)
	if err != nil {
		panic(err)
	}
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
