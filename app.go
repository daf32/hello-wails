package main

import (
	"context"
	"fmt"
	"hello-wails/internal/domain"
	"hello-wails/internal/service"
)

// App struct
type App struct {
	ctx     context.Context
	userSvc *service.UserService
	calcSvc *service.CalculatorService
}

// NewApp creates a new App application struct
func NewApp(
	userSvc *service.UserService,
	calcSvc *service.CalculatorService,
) *App {
	return &App{
		userSvc: userSvc,
		calcSvc: calcSvc,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s from wails", name)
}

func (a *App) GetUser(id int) (domain.User, error) {
	return a.userSvc.GetUser(a.ctx, id)
}

func (a *App) ListUsers(limit int) ([]domain.User, error) {
	return a.userSvc.ListUsers(a.ctx, limit)
}

func (a *App) Divide(dividend, divisor float64) (float64, error) {
	return a.calcSvc.Divide(dividend, divisor)
}
