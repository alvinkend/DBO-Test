package main

import (
	"pustaka-api/handler"
	"pustaka-api/repository"
	"pustaka-api/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:makanenak@tcp(127.0.0.1:3306)/DBO?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	router := gin.Default()
	v1 := router.Group("/v1")

	authRepo := repository.NewCustomerRepository(db)
	authSrv := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authSrv)	
	v1.POST("/login", authHandler.LoginHandler)

	customerRepo := repository.NewCustomerRepository(db)
	customerSrv := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerSrv)
	
	v1.GET("/customers", customerHandler.GetcustomersHandler)
	v1.GET("/customer/:id", customerHandler.GetcustomerHandler)
	v1.POST("/register-customer", customerHandler.PostcustomerHandler)
	v1.PUT("/customer/:id", customerHandler.UpdatecustomerHandler)
	v1.DELETE("/customer/:id", customerHandler.DeletecustomerHandler)

	customerDetailRepo := repository.NewCustomerDetailRepository(db)
	customerDetailSrv := usecase.NewCustomerDetailUsecase(customerDetailRepo)
	customerDetailHandler := handler.NewCustomerDetailHandler(customerDetailSrv)

	v1.GET("/customers-details", customerDetailHandler.GetCustomerDetailsHandler)
	v1.GET("/customer-detail/:id", customerDetailHandler.GetCustomerDetailHandler)
	v1.POST("/customer-detail", customerDetailHandler.PostCustomerDetailHandler)
	v1.PUT("/customer-detail/:id", customerDetailHandler.UpdateCustomerDetailHandler)
	v1.DELETE("/customer-detail/:id", customerDetailHandler.DeleteCustomerDetailHandler)

	orderRepo := repository.NewOrderRepository(db)
	orderSrv := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSrv)

	v1.GET("/orders", orderHandler.GetOrdersHandler)
	v1.GET("/order/:id", orderHandler.GetOrderHandler)
	v1.POST("/order", orderHandler.PostOrderHandler)
	v1.PUT("/order/:id", orderHandler.UpdateOrderHandler)
	v1.DELETE("/order/:id", orderHandler.DeleteOrderHandler)

	orderDetailRepo := repository.NewOrderDetailRepository(db)
	orderDetailSrv := usecase.NewOrderDetailUsecase(orderDetailRepo)
	orderDetailHandler := handler.NewOrderDetailHandler(orderDetailSrv)

	v1.GET("/order-details", orderDetailHandler.GetOrderDetailsHandler)
	v1.GET("/order-details/:id", orderDetailHandler.GetOrderDetailHandler)
	v1.POST("/order-detail", orderDetailHandler.PostOrderDetailHandler)
	v1.PUT("/order-detail/:id", orderDetailHandler.UpdateOrderDetailHandler)
	v1.DELETE("/order-detail/:id", orderDetailHandler.DeleteOrderDetailHandler)

	router.Run()
}
