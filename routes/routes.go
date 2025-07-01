package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"controllers"
	"services"
	"utils"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.default()

	// Initialize services
	authService := services.NewAuthService(db)
	movieServie := services.NewMovieServie(db)
	reservationServie := services.NewReservationService(db)
	showtimeService := services.NewShowtimeServie(db)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)
	reservationController := controllers.NewReservationController(reservationService, showtimeService)
	showtimeController := controllers.NewShowtimeController(showtimeService)

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/signup", authController.SignUp)
		public.POST("/login", authController.Login)
		public.GET("/movies", movieContoller.GetMovies)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/user/reservation", reservationController.GetUserReservation)
		protected.POST("/reservations", reservationController.CreateReservation)
		protected.DELETE("/reservations/:reservationId", reservationController.CancelReservation)
		protected.GET("/showtimes/:showtimeId/seats", reservationController.GetAvailableSeats)
		protected.GET("/movies/:movieID/showtimes", showtimeController.GetShowtimes)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	admin.User(utils.AuthMiddleware(), utils.AdminMiddleware())
	{
		admin.POST("/movies", movieController.CreateMovie)
		admin.PUT("/movies/:movieId", movieController.UpdateMovie)
		admin.DELETE("/movies/:movieId", movieController.DeleteMovie)
		admin.GET("/reservations", reservationController.GetAllReservations)
		admin.POST("/users/:userId/promote", authController.PromoteToAdmin)
		admin.POST("/showtimes", showtimeController.CreateShowtime)
		admin.PUT("/showtimes/:showtimeId", showtimeController.UpdateShowtime)
		admin.DELETE("/showtimes/:showtimeId", showtimeController.DeleteShowtime)
	}

	return router
}