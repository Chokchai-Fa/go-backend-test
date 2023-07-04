package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Backend Team",
		})
	})

	router.POST("/create", createUser)

	router.GET("/get/:id", getUserById)
	router.GET("/getAll", getAllUser)

	router.PUT("/update/:id", updateAllDetail)

	router.PATCH("/updateSome/:id", updateSomeDetail)

	router.DELETE("/delete/:id", deleteUserbyId)

	server.router = router
}

// Start run HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
