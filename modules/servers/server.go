package servers

import (
	"bytes"
	"encoding/json"
	"integration-suspect-service/configs"
	"integration-suspect-service/pkg/loggers"
	"integration-suspect-service/pkg/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App      *fiber.App
	Cfg      *configs.Configs
	Db       *sqlx.DB
	C        *cache.Cache
	Log      *loggers.Logger
	Validate *validator.Validate
}

func NewServer(cfg *configs.Configs, db *sqlx.DB, c *cache.Cache, log *loggers.Logger, validate *validator.Validate) *Server {
	fiberConfig := fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
		JSONDecoder: func(data []byte, v interface{}) error {
			decoder := json.NewDecoder(bytes.NewReader(data))
			decoder.DisallowUnknownFields()
			return decoder.Decode(v)
		},
	}

	return &Server{
		App:      fiber.New(fiberConfig),
		Cfg:      cfg,
		Db:       db,
		C:        c,
		Log:      log,
		Validate: validate,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
