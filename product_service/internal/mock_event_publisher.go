package internal

import "fmt"

type mockPublisher struct {
}

func (*mockPublisher) PublishCreatedEvent(e *ProductCreatedEvent) error {
	fmt.Printf("Publishing Created Event, e: %+v\n", e)
	return nil
}

func (*mockPublisher) PublishUpdatedEvent(e *ProductUpdatedEvent) error {
	fmt.Printf("Publishing Updated Event, e: %+v\n", e)
	return nil
}

type errMockPublisher struct {
}

func (*errMockPublisher) PublishCreatedEvent(e *ProductCreatedEvent) error {
	return ErrPublish
}

func (*errMockPublisher) PublishUpdatedEvent(e *ProductUpdatedEvent) error {
	return ErrPublish
}
