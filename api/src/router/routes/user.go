package routes

import (
	controller "api/src/Controller"
	"net/http"
)

var routeUser = []Rout{
	{
		Uri:                   "/user",
		Method:                http.MethodPost,
		Function:              controller.CreateUser, // Remova os parênteses
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user",
		Method:                http.MethodGet,
		Function:              controller.GetUsers, // Remova os parênteses
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodGet,
		Function:              controller.GetUser, // Remova os parênteses
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodPut,
		Function:              controller.UpdateUser, // Remova os parênteses
		RequireAuthentication: false,
	},
	{
		Uri:                   "/user/{id}",
		Method:                http.MethodDelete,
		Function:              controller.DeleteUser, // Remova os parênteses
		RequireAuthentication: false,
	},
}
