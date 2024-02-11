package cmd

import (
	"context"
	"github.com/tejiriaustin/ToW/database"
	"github.com/tejiriaustin/ToW/payment"

	"github.com/spf13/cobra"

	"github.com/tejiriaustin/ToW/env"
	"github.com/tejiriaustin/ToW/repository"
	"github.com/tejiriaustin/ToW/server"
	"github.com/tejiriaustin/ToW/services"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts Tree of Wally API",
	Long:  ``,
	Run:   startApi,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startApi(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	config := setApiEnvironment()

	dbConn, err := database.NewMongoDbClient().Connect(config.GetAsString(env.MongoDsn), config.GetAsString(env.MongoDbName))
	if err != nil {
		panic("Couldn't connect to mongo dsn: " + err.Error())
	}
	defer func() {
		_ = dbConn.Disconnect(context.TODO())
	}()

	rc := repository.New(dbConn)

	sc := services.New(
		services.WithPaymentProcessor(payment.NewPaymentProcessor()),
	)

	server.Start(ctx, sc, rc, &config)
}

func setApiEnvironment() env.Config {
	staticEnvironment := env.NewEnvironment()

	staticEnvironment.
		SetEnv(env.EnvPort, env.GetEnv(env.EnvPort, "8080")).
		SetEnv(env.JwtSecret, env.MustGetEnv(env.JwtSecret)).
		SetEnv(env.MongoDsn, env.MustGetEnv(env.MongoDsn)).
		SetEnv(env.MongoDbName, env.MustGetEnv(env.MongoDbName)).
		SetEnv(env.MinimumFollowSpend, 1)

	return staticEnvironment
}
