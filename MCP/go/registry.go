package main

import (
	"github.com/spell-check-client/mcp-server/config"
	"github.com/spell-check-client/mcp-server/models"
	tools_spellcheck "github.com/spell-check-client/mcp-server/tools/spellcheck"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_spellcheck.CreateSpellcheckerTool(cfg),
	}
}
