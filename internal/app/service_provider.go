package app

import (
	"context"
	"log"

	"github.com/kirillmc/platform_common/pkg/closer"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/platform_common/pkg/db/pg"
	"github.com/kirillmc/trainings-server/internal/api/training"
	"github.com/kirillmc/trainings-server/internal/config"
	"github.com/kirillmc/trainings-server/internal/config/env"
	"github.com/kirillmc/trainings-server/internal/repository"
	trainingRepo "github.com/kirillmc/trainings-server/internal/repository/training"
	"github.com/kirillmc/trainings-server/internal/service"
	trainingService "github.com/kirillmc/trainings-server/internal/service/training"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient db.Client

	trainingRepository     repository.TrainingRepository
	trainingService        service.TrainingService
	trainingImplementation *training.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = pgConfig
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = grpcConfig
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		httpConfig, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = httpConfig
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		swaggerConfig, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = swaggerConfig
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.pgConfig.DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TrainingRepository(ctx context.Context) repository.TrainingRepository {
	if s.trainingRepository == nil {
		s.trainingRepository = trainingRepo.NewRepository(s.DBClient(ctx))
	}

	return s.trainingRepository
}

func (s *serviceProvider) TrainingService(ctx context.Context) service.TrainingService {
	if s.trainingService == nil {
		s.trainingService = trainingService.NewService(s.TrainingRepository(ctx))
	}

	return s.trainingService
}

func (s *serviceProvider) TrainingImplementation(ctx context.Context) *training.Implementation {
	if s.trainingImplementation == nil {
		s.trainingImplementation = training.NewImplementation(s.TrainingService(ctx))
	}

	return s.trainingImplementation
}
