package controllers

import "event_backend/api/middlewares"

func (s *Server) intializeRoutes() {

	v1 := s.Router.Group("api/v1")
	{
		// Register Route
		v1.POST("/register", s.CreateUser)

		// Login
		v1.POST("/login", s.Login)

		// Users
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUserByID)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		// User Profile
		v1.PUT("/profile", s.initProfile)

		// Events
		v1.POST("/create-event", s.CreateEvent)

	}

}
