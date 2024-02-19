package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", printHeader)
	addr := "0.0.0.0:8080"
	log.Fatalln(e.Start(addr))
}

func printHeader(ctx echo.Context) error {
	headers := ctx.Request().Header
	var lines []string
	for k, v := range headers {
		v1 := strings.Join(v, ",")
		lines = append(lines, fmt.Sprintf("%s:%s", k, v1))
	}
	ret := strings.Join(lines, "\n") + "\n"
	ctx.String(http.StatusOK, ret)
	return nil
}
