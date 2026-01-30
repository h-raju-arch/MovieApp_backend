package main

import (
	httptransport "ggithub.com/h-raju-arch/MoiveApp_backend/internal/transport/http"
	"github.com/h-raju-arch/MoiveApp_backend/internal/db"
	movierepo "github.com/h-raju-arch/MoiveApp_backend/internal/repo/movie_repo"
	"github.com/h-raju-arch/MoiveApp_backend/internal/service"
)

func main() {
	database := db.Open()
	defer database.Close()
	repo := movierepo.New_Movie_Repo(database)
	svc := service.New_Movie_Service(*repo)
	router := httptransport.NewRouter(svc)

	router.Run(":3000")
}
