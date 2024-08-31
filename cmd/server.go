package cmd

import (
	"context"
	"fmt"
	"log"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dhamanutd/golang-clean-architecture/api/rest"
	orders_controller "github.com/dhamanutd/golang-clean-architecture/api/rest/controller/v1/orders"
	"github.com/dhamanutd/golang-clean-architecture/configs"
	"github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql"
	postgresql_orders "github.com/dhamanutd/golang-clean-architecture/internal/repositories/postgresql/orders"
	orders_usecase "github.com/dhamanutd/golang-clean-architecture/internal/usecases/orders"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Start HTTP server",
	Long:  `Start Boilerplate HTTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init config
		config, secret, err := configs.LoadConfig(cfgFile, scrtFile)
		if err != nil {
			log.Fatalf("Unable to load configuration and secret: %v", err)
		}

		// PostgreSQL Database
		postgresConfig := postgresql.Config{
			Host:     config.PostgreSQLConfig.Host,
			Port:     config.PostgreSQLConfig.Port,
			Username: secret.PostgreSQLSecret.Username,
			Password: secret.PostgreSQLSecret.Password,
			DBName:   secret.PostgreSQLSecret.Database,
		}

		dbClient, err := postgresql.New(postgresConfig)
		if err != nil {
			fmt.Printf("Unable to init mysql gateway, %v", err)
			panic(err)
		}

		// Repositories
		orderRepo := postgresql_orders.New(dbClient)

		// Usecases
		orderUsecase := orders_usecase.New(orderRepo)

		// Controllers
		orderController := orders_controller.New(orderUsecase)

		routerModule := rest.RouterModule{
			V1OrderController: orderController,
		}

		if err := routerModule.Validate(); err != nil {
			log.Fatalf("error new http router: %v", err)
		}

		r := rest.NewRouter(routerModule)

		server := &netHttp.Server{Addr: ":3000", Handler: r}
		// Start the server in a goroutine
		go func() {
			if err := server.ListenAndServe(); err != nil && err != netHttp.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Set up signal capturing
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal.
		<-quit

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline or until all connections have returned.
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

		log.Println("Server gracefully stopped")
	},
}
