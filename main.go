package main

import (
	"github.com/Hulhay/jk-pengker/config"
	"github.com/Hulhay/jk-pengker/controller"
	"github.com/Hulhay/jk-pengker/middleware"
	"github.com/Hulhay/jk-pengker/repository"
	"github.com/Hulhay/jk-pengker/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepo  repository.UserRepository  = repository.NewUserRepository(db)
	storeRepo repository.StoreRepository = repository.NewStoreRepository(db)

	tokenUC usecase.Token = usecase.NewTokenUc()
	authUC  usecase.Auth  = usecase.NewAuthUC(userRepo)
	storeUC usecase.Store = usecase.NewStoreUC(storeRepo)

	ac controller.AuthController  = controller.NewAuthController(authUC, tokenUC)
	sc controller.StoreController = controller.NewStoreController(storeUC, tokenUC)
)

func main() {

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
	}

	storeRoutes := r.Group("api/stores", middleware.AuthorizeJWT(tokenUC))
	{
		storeRoutes.GET("", sc.GetStoreList)
	}

	r.Run(":8081")
}
