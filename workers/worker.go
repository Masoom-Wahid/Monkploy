package main

import (
	"fmt"
	"log"
	"platform/pkg/env"
	"platform/pkg/tasks"
	"strconv"

	"github.com/hibiken/asynq"
)

func main() {
	// webUI command
	// docker run --rm \
	// --name asynqmon --network drfamo_internal \
	// -p 2828:8080 \
	// hibiken/asynqmon --redis-addr=cache:6379 --redis-db=2
	env.SetupEnvFile()

	redisURL := env.GetEnv("REDIS_URL", "")
	var srv *asynq.Server

	if redisURL != "" {
		// Use Redis URL if provided (e.g., redis://:password@host:port/db)
		srv = asynq.NewServer(
			asynq.RedisClientOpt{
				Addr: redisURL,
			},
			asynq.Config{
				Concurrency: 10,
			},
		)
	} else {
		// Fallback to individual env vars
		host := env.GetEnv("REDIS_HOST", "cache")
		port := env.GetEnv("REDIS_PORT", "6379")
		dbNumber, _ := strconv.Atoi(env.GetEnv("REDIS_QUEUE_DB", "5"))
		pass := env.GetEnv("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")

		srv = asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     fmt.Sprintf("%s:%s", host, port),
				DB:       dbNumber,
				Password: pass,
			},
			asynq.Config{
				Concurrency: 10,
			},
		)
	}

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}

	fmt.Println("worker started")
}
