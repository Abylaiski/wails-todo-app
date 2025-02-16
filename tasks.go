package main

import (
	"context"
	"fmt"
)

// App хранит контекст приложения.
type App struct {
	ctx context.Context
}

// NewApp создаёт новый экземпляр App.
func NewApp() *App {
	return &App{}
}

// startup сохраняет переданный контекст для дальнейшего использования.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet формирует приветственное сообщение с переданным именем.
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
