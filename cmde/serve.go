package cmde

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prdcthandler "ecommerce/rest/handlers/product"
	usrhandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middlewares.NewMiddlewares(cnf)

	productHandler := prdcthandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrhandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
