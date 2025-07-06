package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"task-processing-service/internal/web"

	"task-processing-service/cmd"
)

func main() {
	configs := getConfigs()

	compositionRoot := cmd.NewCompositionRoot(configs)
	defer compositionRoot.CloseAll()

	router := web.NewRouter(compositionRoot)

	log.Printf("Task server running on :%s", configs.HttpPort)
	err := http.ListenAndServe(":"+configs.HttpPort, router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getConfigs() cmd.Config {
	return cmd.Config{
		HttpPort:            getEnv("HTTP_PORT"),
		TaskRunnerWorkers:   mustParseInt(getEnv("TASK_RUNNER_WORKERS")),
		TaskRunnerQueueSize: mustParseInt(getEnv("TASK_RUNNER_QUEUE_SIZE")),
	}
}

func getEnv(key string) string {
	_ = godotenv.Load(".env")
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing env var: %s", key)
	}
	return val
}

func mustParseInt(value string) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Failed to parse int: %v", err)
	}
	return n
}
