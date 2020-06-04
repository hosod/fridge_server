package content

import (
	"github.com/gin-gonic/gin"
	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/entity"
)

type Service struct {}


func (s *Service) GetByID(id string) () {
	db := database.GetDB()
	var content entity.
	
}