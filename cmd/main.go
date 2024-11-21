package main

import (
	"fmt"
	"github.com/MCPutro/E-commerce/internal/delivery"
	"github.com/MCPutro/E-commerce/internal/middleware"
	"log"
	"net/http"

	"github.com/MCPutro/E-commerce/internal/db"
	"github.com/MCPutro/E-commerce/internal/repository"
	"github.com/MCPutro/E-commerce/internal/usecase"
	"github.com/MCPutro/E-commerce/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {

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
