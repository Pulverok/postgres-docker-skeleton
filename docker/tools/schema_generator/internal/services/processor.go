package services

//go:generate mockgen -source=processor.go -destination=processor_mock.go -package=services

// Processor is an interface for processing data.
type Processor interface {
	Process() error
}
