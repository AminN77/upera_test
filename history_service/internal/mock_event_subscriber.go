package internal

type mockEventSubscriber struct {
}

func NewMockEventSubscriber() EventSubscriber {
	return &mockEventSubscriber{}
}

func (*mockEventSubscriber) Subscribe() error {
	return nil
}

func (*mockEventSubscriber) UnSubscribe() error {
	return nil
}
