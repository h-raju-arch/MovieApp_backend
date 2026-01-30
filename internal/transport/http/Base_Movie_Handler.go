package httptransport

import "github.com/h-raju-arch/MoiveApp_backend/internal/service"

type Movie_handler struct {
	svc service.Movie_Service
}

func New_Movie_Handler(svc service.Movie_Service) *Movie_handler {
	return &Movie_handler{svc: svc}
}
