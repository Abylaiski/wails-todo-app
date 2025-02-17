package main

import (
	"context"
	"fmt"
)

<<<<<<< HEAD
// App хранит контекст приложения.
=======
// App struct
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
type App struct {
	ctx context.Context
}

<<<<<<< HEAD
// NewApp создаёт новый экземпляр App.
=======
// NewApp creates a new App application struct
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func NewApp() *App {
	return &App{}
}

<<<<<<< HEAD
// startup сохраняет переданный контекст для дальнейшего использования.
=======
// startup is called when the app starts. The context is saved
// so we can call the runtime methods
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

<<<<<<< HEAD
// Greet формирует приветственное сообщение с переданным именем.
=======
// Greet returns a greeting for the given name
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
