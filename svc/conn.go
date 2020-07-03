package svc

// TODO: Channel, RPC and HTTP implementations.
// Conn represents a bidirectional means
// of communication between two services.
type Conn interface {
	Send()
	Receive()
	Open()
	Close()
}
