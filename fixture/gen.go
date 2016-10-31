package main

import (
	"fmt"
	"log"
	"strings"
	"time"

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

	rpo := repository.NewUrlInfo(c)

	for i := 0; i < 100; i++ {
		mdl := generateUrlInfo(fake)
		if _, err := rpo.Save(mdl); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
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
