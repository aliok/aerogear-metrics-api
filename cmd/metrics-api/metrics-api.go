package main

import (
	"net/http"

	"github.com/aerogear/aerogear-metrics-api/pkg/dao"
	"github.com/aerogear/aerogear-metrics-api/pkg/mobile"
	"github.com/aerogear/aerogear-metrics-api/pkg/web"
	"log"
)

const ListenAddress = ":3000"

func main() {
	router := web.NewRouter()
	db, err := dao.Connect()
	if err != nil {
		panic("failed to connect to sql database : " + err.Error())
	}
	defer db.Close()
	metricDao := dao.NewMetricsDAO(db)

	//metrics route
	{
		metricsService := mobile.NewMetricsService(metricDao)
		metricsHandler := web.NewMetricsHandler(metricsService)
		web.MetricsRoute(router, metricsHandler)
	}
	//health route
	{
		healthHandler := web.NewHealthHandler()
		web.HealthzRoute(router, healthHandler)
	}

	log.Printf("Starting application... going to listen on %v", ListenAddress)

	//start
	if err := http.ListenAndServe(":3000", router); err != nil {
		panic("failed to start " + err.Error())
	}
}
