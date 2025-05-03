package main

import (
	"log/slog"

	"oxid-gateway-admin-api/pkg/db"
	"oxid-gateway-admin-api/pkg/handlers"
	"oxid-gateway-admin-api/pkg/services"

	"github.com/go-fuego/fuego"
	"github.com/joho/godotenv"
)

func main() {
	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		slog.Warn("No .env file found")
	}

	db.Connect()
	defer db.Close()

	s := fuego.NewServer(
		fuego.WithoutAutoGroupTags(),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				DisableLocalSave: true,
				SpecURL:          "/docs/openapi.json",
				SwaggerURL:       "/docs",
			}),
		),
	)

	database, err := db.GetConection()

	if err != nil {
		panic(err)
	}

	query := db.New(database)

	upstreamService := services.UpstreamsService {
		Repository: query,
	}

	userService := services.UsersService {
		Repository: query,
	}

	userResources := handlers.UsersResources{
		UsersService: &userService,
	}

	upstreamResources := handlers.UpstreamsResources{
		UpstreamService: &upstreamService,
	}

	upstreamResources.Routes(s)
	userResources.Routes(s)

	s.OpenAPI.Description().Info.Title = "Oxid Gateway Admin API"
	s.OpenAPI.Description().Info.Description = "Oxid Gateway Admin API"
	s.OpenAPI.Description().Info.Version = "0.0.1"

	err = s.Run()

	if err != nil {
		panic(err)
	}
}
