package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app     *fiber.App
	options ServerOptions
}

type ServerOptions struct {
	IsHttps      bool
	CertFilePath string
	KeyFilePath  string
	Port         string
}

func New(options *ServerOptions) *Server {
	serverInst := new(Server)
	serverInst.app = fiber.New()
	serverInst.options = getServerOptions(options)
	return serverInst
}

func (s *Server) AddRoute(method, path string, handler func(c *fiber.Ctx) error) {
	s.app.Add(method, path, handler)
}

func (s *Server) Start() {
	s.app.Hooks().OnListen(onStart)
	s.app.Hooks().OnShutdown(onShutdown)

	if s.options.IsHttps {
		panic(s.app.ListenTLS(s.options.Port, s.options.CertFilePath, s.options.KeyFilePath))
	} else {
		panic(s.app.Listen(s.options.Port))
	}
}

func onStart() error {
	fmt.Println("Starting server ...")
	return nil
}

func onShutdown() error {
	fmt.Println("Shutting down server.")
	return nil
}

func getServerOptions(sOptions *ServerOptions) ServerOptions {
	options := ServerOptions{
		IsHttps:      false,
		CertFilePath: "",
		KeyFilePath:  "",
		Port:         ":80",
	}

	if sOptions == nil {
		return options
	}

	if sOptions.IsHttps {
		options.IsHttps = true
		options.CertFilePath = sOptions.CertFilePath
		options.KeyFilePath = sOptions.KeyFilePath
		options.Port = ":443"
	}

	if sOptions.Port != "" {
		options.Port = sOptions.Port
	}

	return options
}

