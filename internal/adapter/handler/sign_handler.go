package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load(`../../../.env`)
}

func Attest(c echo.Context) error {
	rpcUrl := os.Getenv(`RPC_URL`)
	fmt.Printf("RPC " + rpcUrl)
	return c.JSON(http.StatusOK, map[string]bool{"Healthy": true})
}
