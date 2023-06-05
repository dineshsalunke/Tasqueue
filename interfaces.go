package tasqueue

import "context"

type Results interface {
	Get(ctx context.Context, uuid string) ([]byte, error)
	Set(ctx context.Context, uuid string, b []byte) error
	// DeleteJob removes the job's saved metadata from the store
	DeleteJob(ctx context.Context, uuid string) error
	GetFailed(ctx context.Context) ([]string, error)
	GetSuccess(ctx context.Context) ([]string, error)
	SetFailed(ctx context.Context, uuid string) error
	SetSuccess(ctx context.Context, uuid string) error
}

type Broker interface {
	// Enqueue places a task in the queue
	Enqueue(ctx context.Context, msg []byte, queue string) error

	// Consume listens for tasks on the queue and calls processor
	Consume(ctx context.Context, work chan []byte, queue string)

	// GetPending returns a list of stored job messages on the particular queue
	GetPending(ctx context.Context, queue string) ([]string, error)
}
