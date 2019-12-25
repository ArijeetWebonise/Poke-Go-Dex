package app

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {
	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	// app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/sagar/:name", app.renderView(app.SagarPage))

	app.Router.Get("/pokemon/generate", app.renderView(app.GenerateNewPokemon))
	app.Router.Get("/pokemon", app.renderView(app.DisplayPokemons))
	app.Router.Get("/pokemon/:id", app.renderView(app.GetPokemon))
}
