package main

import (
	"log/slog"

	"oxid-gateway-admin-api/pkg/db"
	"oxid-gateway-admin-api/pkg/handlers"
	"oxid-gateway-admin-api/pkg/services"

	"github.com/getkin/kin-openapi/openapi3"
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
		fuego.WithSecurity(map[string]*openapi3.SecuritySchemeRef{
			"bearerAuth": {
				Value: openapi3.NewSecurityScheme().
					WithType("http").
					WithScheme("bearer").
					WithBearerFormat("JWT").
					WithDescription("Enter your JWT token in the format: Bearer <token>"),
			},
		}),
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

	upstreamService := services.UpstreamsService{
		Repository: query,
	}

	userService := services.UsersService{
		Repository: query,
	}

	requestContextService := handlers.RequestContextService{
		UsersService: &userService,
	}

	handlers.NewRequestContextService(requestContextService)

	userResources := handlers.UsersResources{
		UsersService: &userService,
	}

	upstreamResources := handlers.UpstreamsResources{
		UpstreamService: &upstreamService,
		UsersService: &userService,
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
