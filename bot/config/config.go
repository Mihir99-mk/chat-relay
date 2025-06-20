package config

import "bot/ent/entgen"

type IConfig interface {
	Env() IEnv
	DB() *entgen.Client
}

type config struct {
	env IEnv
	db  *entgen.Client
}

func NewConfig() IConfig {
	env := NewEnv()
	return &config{
		env: env,
		db:  GetDB(env),
	}
}

func (c *config) Env() IEnv {
	return c.env
}

func (c *config) DB() *entgen.Client {
	return c.db
}
