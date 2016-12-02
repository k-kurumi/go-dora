package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var e *echo.Echo = echo.New()

func main() {
	e.Logger.SetLevel(log.INFO)

	e.GET("/", getJSON)

	// 遅延して応答する
	e.GET("/delay/:second", getDelay)

	// pingする
	e.GET("/ping/:address", getPing)

	// 指定のHTTPステータスを返す
	e.GET("/status/:code", getStatus)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func getJSON(c echo.Context) error {
	jsonMap := map[string]string{
		"foo": "FOO",
		"bar": "BAR",
	}
	return c.JSON(http.StatusOK, jsonMap)
}

// e.GET("/delay/:second", getDelay)
func getDelay(c echo.Context) error {
	second, _ := strconv.Atoi(c.Param("second"))
	time.Sleep(time.Duration(second) * time.Second)
	return c.String(http.StatusOK, fmt.Sprintf("delay: %d sec", second))
}

// 外部コマンド実行
func getPing(c echo.Context) error {
	address := c.Param("address")
	out, err := exec.Command("ping", "-c4", address).Output()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, string(out))
}

// 指定されたHTTPステータスを返す
func getStatus(c echo.Context) error {
	code, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(code, "")
}
