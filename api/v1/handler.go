package v1

import (
	"github.com/dentist/config"
	"github.com/dentist/pkg/logger"
	"github.com/dentist/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
	logger logger.Logger
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
	Logger logger.Logger
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg: options.Cfg,
		storage: options.Storage,
		logger: options.Logger,
	}
}