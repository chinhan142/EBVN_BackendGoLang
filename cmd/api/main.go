package main

import "github.com/chinhan142/EBVN_BackendGoLang/internal/api"

func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
