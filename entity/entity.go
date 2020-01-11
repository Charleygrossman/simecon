package entity

// An Entity is any struct that implements Routes,
// which implies it has a Server struct field.
type Entity interface {
	Routes()
}
