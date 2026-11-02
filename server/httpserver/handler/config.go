package handler

import (
	"log/slog"

	"github.com/Chad-Glazier/controlward/riot"
)

type Config struct {
	Riot   *riot.Client
	Logger *slog.Logger
}
