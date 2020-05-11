package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/service/fridge"
)


// Fridge is local test
func TestFridge(t *testing.T) {
	database.TestInit()
	// db := database.GetDB()
	var service fridge.Service
	fridge, err := service.GetByID("1")
	log.Println(fridge.Name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(fridge.Name)
	if fridge.Name != "山田家" {
		t.Errorf("got: %v\nwant: %v", fridge.Name, "山田家")
	}

}
