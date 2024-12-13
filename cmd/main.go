package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/MCPutro/E-commerce/internal/delivery"
	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/internal/middleware"
	"github.com/MCPutro/E-commerce/internal/repository/user"
	"github.com/google/uuid"

	"github.com/MCPutro/E-commerce/internal/db"
	"github.com/MCPutro/E-commerce/internal/repository"
	"github.com/MCPutro/E-commerce/internal/usecase"
	"github.com/MCPutro/E-commerce/pkg/constant"
	"github.com/MCPutro/E-commerce/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main2() {

	logger.NewLogger(logrus.DebugLevel)

	newBb, _ := db.GetMysqlConnection()
	defer newBb.Close()

	fmt.Println("start")

	productRepository := repository.NewProductRepository()
	cartRepository := repository.NewCartRepository()
	orderRepository := repository.NewOrderRepository()

	orderUseCase := usecase.NewOrderUseCase(productRepository, cartRepository, orderRepository, newBb)

	productUseCase := usecase.NewProductUseCase(productRepository, newBb)

	cartUseCase := usecase.NewCartUseCase(cartRepository, productRepository, newBb)

	mux := http.NewServeMux()

	handler := delivery.NewHandler(productUseCase, cartUseCase, orderUseCase)

	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":9999", middleware.NewMiddleware(mux)))

	////routing
	//handler := delivery.NewHandler(productUseCase, cartUseCase, orderUseCase)
	//
	//handler.RegisterRoutes(mux)

	//for i := 1; i <= 100; i++ {
	//	// add to cart
	//	cartService.AddToCart(context.Background(), uint(i), &entity.CartItem{
	//		ProductID: 1,
	//		Quantity:  1,
	//	})
	//}
	//
	//cart, err := cartService.GetCart(context.Background(), 101)
	//if err != nil {
	//	if errors.Is(err, newError.ErrCartNotFound) {
	//		fmt.Println("cart not found ------------")
	//		fmt.Printf("newError.GetErrorCode(err): %v\n", newError.GetErrorCode(err))
	//	} else {
	//		fmt.Println("other error =============")
	//	}
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("cart : ", cart)
	//}

	// group := sync.WaitGroup{}

	// for i := 1; i <= 1000; i++ {
	// 	group.Add(1)
	// 	go func(x uint) {
	// 		defer group.Done()
	// 		checkout, err := orderService.Checkout(context.Background(), x)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("checkout : ", checkout)
	// 		}

	// 	}(uint(i))
	// }

	// group.Wait()
	// fmt.Println("selesai")

	//app := fiber.New()
	//
	//app.Get("/", func(c *fiber.Ctx) error {
	//	query, _ := db.Query("select version();")
	//	defer query.Close()
	//	var data string
	//	query.Scan(&data)
	//	return c.SendString("Hello, World! -> " + data)
	//})
	//
	//fmt.Println("running in port 3000")
	//
	//app.Listen(":3000")
}

func main() {
	newBb, _ := db.GetMysqlConnection()
	defer newBb.Close()

	tx, err := newBb.Begin()
	if err != nil {
		logrus.Infoln("error :", err)
		return
	}
	defer tx.Rollback()

	userRepo := user.NewUserRepository()

	newUser := domain.User{
		Id:       uuid.NewString(),
		Name:     "name",
		Email:    "email111@gmail.com",
		Password: "password",
		Role:     constant.Staff,
	}

	userRepo.Create(context.Background(), tx, &newUser)

	newUser2 := domain.User{
		Id:       uuid.NewString(),
		Name:     "name",
		Email:    "email222@gmail.com",
		Password: "password",
		Role:     constant.Staff,
		Address: []domain.UserAddress{
			{Address: "Jakarta", City: "jkt", PostalCode: "123"},
			{Address: "Bandung", City: "bdg", PostalCode: "456"},
		},
	}
	userRepo.Create(context.Background(), tx, &newUser2)

	all, err := userRepo.FindAll(context.Background(), tx)
	fmt.Println(err)
	fmt.Println(all)

	tx.Commit()

}
