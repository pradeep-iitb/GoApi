package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pradeep-iitb/GoApi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println(`   ______  ____        ___    ____  ____
  / ____/ / __ \      /   |  / __ \/  _/
 / / __  / / / /     / /| | / /_/ // /
/ /_/ / / /_/ /     / ___ |/ ____// /
\____/  \____/     /_/  |_/_/   /___/`)


	fmt.Println("Starting GO API service ...")


	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
