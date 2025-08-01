package main

import (
	"flag"
	"gblog/app"
	"gblog/app/cronjob"
	"gblog/app/handler"
	"gblog/app/views"
	"log"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/naipad/loach"
)

var (
	addr   = flag.String("addr", ":8080", "TCP address to listen to")
	dbpath = flag.String("dbpath", "localdb", "Directory to serve sdb from")
)

func main() {
	flag.Parse()

	server := app.New(*addr, *dbpath)

	db, err := loach.OpenDefault(*dbpath)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	server.DB = db

	// get CachedSize from app config
	CachedSize := 5
	server.MemCache = fastcache.New(CachedSize * 1024 * 1024)

	cronManager := cronjob.NewCronJobManager(server)
	cronManager.RegisterCronJob("Clean up expired articles", "*/1 * * * *", cleanUpExpiredArticles)

	go cronManager.Start()

	log.Println("Registered Cron Jobs:", cronManager.GetCronJobList())

	if err := cronManager.StopCronJob("Clean up expired articles"); err != nil {
		log.Printf("Error stopping cron job: %v", err)
	}

	handler.MakeRoutes(server, views.AssetFiles, views.TemplateFiles)

	server.Start()
}

func cleanUpExpiredArticles() {
	log.Println("Running cron job: Clean up expired articles...")
}
