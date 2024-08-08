package grid_cli

type Message struct {
	Parms   []interface{}          `json:"parms"`   // Parameters: first element is the promise
	Payload map[string]interface{} `json:"payload"` // Metadata about the promise
}
