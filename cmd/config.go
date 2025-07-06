package cmd

type Config struct {
	HttpPort            string
	TaskRunnerWorkers   int
	TaskRunnerQueueSize int
}
