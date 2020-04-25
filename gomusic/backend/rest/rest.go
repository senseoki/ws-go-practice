package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return runAPIWithHandler(address, h)
}

func runAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.GET("/products", h.GetProducts)

	// 프로모션 목록
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)

}
