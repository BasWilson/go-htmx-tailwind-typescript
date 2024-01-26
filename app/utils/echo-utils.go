package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func Rt(e echo.Context, template templ.Component) error {
	return template.Render(e.Request().Context(), e.Response())
}

func WsHotReload(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
