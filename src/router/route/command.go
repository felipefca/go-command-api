package route

import (
	"api/src/controllers"
	"net/http"
)

var commandRoute = []Route{
	{
		Uri:      "/stock",
		Method:   http.MethodPost,
		Function: controllers.SaveStock,
		Auth:     false,
	},
	{
		Uri:      "/fiat",
		Method:   http.MethodPost,
		Function: controllers.SaveFiat,
		Auth:     false,
	},
}
