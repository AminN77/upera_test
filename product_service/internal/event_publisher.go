package internal

type EventPublisher interface {
	PublishCreatedEvent(e *ProductCreatedEvent) error
	PublishUpdatedEvent(e *ProductUpdatedEvent) error
}
