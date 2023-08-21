package api

import (
	"os"
	"testing"
	"user_api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	collect, err := db.NewDB()
	if err != nil {
		panic("Error new db")
	}
	r = NewApi(collect).NewRouter()

	os.Exit(m.Run())
}
