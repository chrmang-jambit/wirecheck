package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := readConfig()
	e := echo.New()
	e.GET("/*", printHeader)

	addr := fmt.Sprintf("%s:%d", cfg.serverAddr, cfg.serverPort)
	log.Fatalln(e.Start(addr))
}

func printHeader(ctx echo.Context) error {
	headers := ctx.Request().Header
	var lines []string
	for k, v := range headers {
		v1 := strings.Join(v, ",")
		lines = append(lines, fmt.Sprintf("%s:%s", k, v1))
	}
	slices.Sort(lines)
	ret := strings.Join(lines, "\n") + "\n"
	ctx.String(http.StatusOK, ret)
	return nil
}
