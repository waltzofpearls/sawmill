package main

import (
	"log"

	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/logger"
	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

func main() {
	cf, err := config.New("config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	lg, err := logger.New(cf)
	if err != nil {
		log.Fatalln(err)
	}
	db, err := database.New(cf, lg)
	if err != nil {
		log.Fatalln(err)
	}
	c := db.Cluster

	rpo := repository.NewUrlInfoRepository(c)
	mdl := model.NewUrlInfoModel(
		"test.com/test/aaa/bbb.html?var=val&var2=val2",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		true,
	)

	if _, err := rpo.Save(mdl); err != nil {
		log.Fatalln(err)
	}
}
