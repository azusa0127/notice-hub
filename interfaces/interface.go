package interfaces

import "context"

// SendClient is the general interface for all notification sending clents
type SendClient interface {
	// Send the message out with using request key from the context passed in
	Send(ctx context.Context) error
}
