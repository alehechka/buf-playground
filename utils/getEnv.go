package utils

import "os"

var (
	// gRPC server endpoints
	GRPCInventoryServerEndpoint = GetEnv("GRCP_INVENTORY_SERVER_ENDPOINT", "localhost:3001")
	GRPCSessionServerEndpoint   = GetEnv("GRCP_SESSION_SERVER_ENDPOINT", "localhost:3002")
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
