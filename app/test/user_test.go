package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hosod/fridge_server/app/internal/database"
	"github.com/hosod/fridge_server/app/internal/service/user"
)

func TestNameList(t *testing.T) {
	database.TestInit()
	// db := database.GetDB()
	var service user.Service
	names, err := service.GetWholeNameList()
	if err!=nil {
		t.Errorf("got Error: %v\n", err)
	}
	fmt.Println(names)

	expected := []string{"yamada", "tanaka", "sato"}
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("got: %v\nwant: %v", names, expected)
	}
}
