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
		v1.GET("/me", middlewares.TokenAuthMiddleware(), s.GetMe)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdatePassword)
		v1.PUT("/users-email/:id", middlewares.TokenAuthMiddleware(), s.UpdateEmail)

		// User Profile
		v1.PUT("/profile", s.initProfile)

		// Events
		v1.POST("/create-event", middlewares.TokenAuthMiddleware(), s.CreateEvent)
		v1.GET("/user_events/:id", s.GetUserEvents)
		v1.GET("/events/:id", s.GetEvent)
		v1.GET("/events", s.GetEvents)

		// Events-Types Routes
		v1.GET("/music", s.GetMusicEvents)
		v1.GET("/sports", s.GetSportEvents)
		// v1.GET("/entertainment")
		// v1.GET("/cinema")
		// v1.GET("/arts")
		// v1.GET("/dance")
		// v1.GET("/kids")
		// v1.GET("/social_events")
	}

}
