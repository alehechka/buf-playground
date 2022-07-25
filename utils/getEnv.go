package utils

import (
	"fmt"
	"os"
)

var (
	// gRPC server endpoints
	GRPCInventoryServerEndpoint = GetEnv("GRCP_INVENTORY_SERVER_ENDPOINT", "localhost:3001")
	GRPCSessionServerEndpoint   = GetEnv("GRCP_SESSION_SERVER_ENDPOINT", "localhost:3002")
	PORT                        = GetEnv("PORT", "80")
	ListenOn                    = fmt.Sprintf(":%s", PORT)
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
