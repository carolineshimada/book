package Routes

import (
	controller "api/src/Controller"
	"net/http"
)

var routeUser = []Rout{
	{
		Uri:                   "/user",
		Method:                http.MethodPost,
		Function:              controller.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user",
		Method:                http.MethodGet,
		Function:              controller.GetUsers,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodGet,
		Function:              controller.GetUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodPut,
		Function:              controller.UpdateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodDelete,
		Function:              controller.DeleteUser,
		RequireAuthentication: false,
	},
}
