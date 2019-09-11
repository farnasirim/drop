package main

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"google.golang.org/grpc"

	drop_grpc "github.com/farnasirim/drop/grpc"
	"github.com/farnasirim/drop/proto"
	"github.com/farnasirim/drop/storage/redis"
)

var (
	serveAddrKey = "serve_addr"
	redisAddrKey = "reddis_addr"
)

func serveCmdFunc(cmd *cobra.Command, args []string) {
	addr := viper.GetString(serveAddrKey)
	lis, err := net.Listen("tcp", addr)

	redisAddr := viper.GetString(redisAddrKey)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	storageServer := redis.NewStorageService(redisAddr)

	grpcServer := grpc.NewServer()
	proto.RegisterDropApiServer(grpcServer, drop_grpc.NewDropServer(storageServer))

	log.Printf("About to listen on %s\n", addr)
	grpcServer.Serve(lis)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve drop server with the grpc api",
	Long:  ``,
	Run:   serveCmdFunc,
}

func serveInit() {
	serveCmd.Flags().String("addr", "", "[host]:port to listen on")
	viper.BindPFlag(serveAddrKey, serveCmd.Flags().Lookup("addr"))
	viper.BindEnv(serveAddrKey)
	viper.SetDefault(serveAddrKey, ":20080")

	serveCmd.Flags().String("redis-addr", "", "[host]:port to access redis on")
	viper.BindPFlag(redisAddrKey, serveCmd.Flags().Lookup("redis-addr"))
	viper.BindEnv(redisAddrKey)
	viper.SetDefault(redisAddrKey, "redis:6379")
}
