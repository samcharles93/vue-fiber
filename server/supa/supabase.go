package supa

import (
	"vue-fiber/config"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func Init(cfg *config.Config) {
	Client = supabase.CreateClient(cfg.Supabase.URL, cfg.Supabase.Key)
}
