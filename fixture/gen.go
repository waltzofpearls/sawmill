package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/manveru/faker"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/logger"
	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

func main() {
	fake, err := faker.New("en")
	if err != nil {
		log.Fatalln(err)
	}

	db, lg, err := connectToRiak()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rpo := repository.NewUrlInfo(db.Cluster)
	num := parseFixtureNum()

	lg.Info(fmt.Sprintf("Generating [%d] data fixtures...", num))

	for i := 0; i < num; i++ {
		mdl := generateUrlInfo(fake)
		if _, err := rpo.Save(mdl); err != nil {
			log.Fatalln(err)
		}
	}
}

func connectToRiak() (*database.Database, *logger.Logger, error) {
	cf, err := config.New("config.yml")
	if err != nil {
		return nil, nil, err
	}
	lg, err := logger.New(cf)
	if err != nil {
		return nil, nil, err
	}
	adapter := &database.Riak{}
	db, err := database.New(adapter, cf, lg)
	if err != nil {
		return nil, nil, err
	}
	return db, lg, nil
}

func parseFixtureNum() int {
	num := 100
	if len(os.Args) > 1 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			num = i
		}
	}
	return num
}

func generateUrlInfo(fake *faker.Faker) *model.UrlInfo {
	fakeUrl := fmt.Sprintf(
		"%s?%s=%s",
		strings.TrimPrefix(fake.URL(), "http://"),
		fake.Characters(5),
		fake.Characters(8),
	)
	fakeDesc := fake.Paragraph(3, false)
	return model.NewUrlInfo(fakeUrl, fakeDesc, true)
}
