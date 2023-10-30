package event

type Publisher interface {
	PublishCreatedEvent(e *ProductCreatedEvent) error
	PublishUpdatedEvent(e *ProductUpdatedEvent) error
}
