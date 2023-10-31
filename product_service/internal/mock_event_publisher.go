package internal

import "fmt"

type mockPublisher struct {
}

func NewMockEventPublisher() EventPublisher {
	return &mockPublisher{}
}

func (*mockPublisher) PublishCreatedEvent(e *ProductCreatedEvent) error {
	fmt.Printf("Publishing Created Event, e: %+v\n", e)
	return nil
}

func (*mockPublisher) PublishUpdatedEvent(e *ProductUpdatedEvent) error {
	fmt.Printf("Publishing Updated Event, e: %+v\n", e)
	return nil
}
