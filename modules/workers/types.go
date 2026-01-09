package workers

// WorkerMessage represents the message structure for the workers.
type WorkerMessage struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
	ID      string      `json:"id"`
}
