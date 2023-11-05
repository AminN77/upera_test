package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		input       *Product
		srv         Service
		expected    *Product
		expectedErr error
	}{
		{
			name:        "success",
			srv:         NewService(&mockRepository{}, &mockPublisher{}),
			input:       &Product{},
			expected:    &Product{},
			expectedErr: nil,
		},
		{
			name:        "repo err",
			srv:         NewService(&errMockRepository{}, &mockPublisher{}),
			input:       &Product{},
			expected:    nil,
			expectedErr: ErrRepo,
		},
		{
			name:        "publisher err",
			srv:         NewService(&mockRepository{}, &errMockPublisher{}),
			input:       &Product{},
			expected:    nil,
			expectedErr: ErrPublisher,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.srv.Add(tc.input)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name        string
		input       *Product
		id          int
		srv         Service
		expected    *Product
		expectedErr error
	}{
		{
			name:        "success",
			srv:         NewService(&mockRepository{}, &mockPublisher{}),
			id:          1,
			input:       &Product{},
			expected:    &Product{},
			expectedErr: nil,
		},
		{
			name:        "repo err",
			srv:         NewService(&errMockRepository{}, &mockPublisher{}),
			id:          1,
			input:       &Product{},
			expected:    nil,
			expectedErr: ErrRepo,
		},
		{
			name:        "publisher err",
			srv:         NewService(&mockRepository{}, &errMockPublisher{}),
			id:          1,
			input:       &Product{},
			expected:    nil,
			expectedErr: ErrPublisher,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.srv.Update(tc.input, tc.id)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestFetch(t *testing.T) {
	testCases := []struct {
		name        string
		id          int
		srv         Service
		expected    *Product
		expectedErr error
	}{
		{
			name:        "success",
			srv:         NewService(&mockRepository{}, &mockPublisher{}),
			id:          1,
			expected:    &Product{},
			expectedErr: nil,
		},
		{
			name:        "repo err",
			srv:         NewService(&errMockRepository{}, &mockPublisher{}),
			id:          1,
			expected:    nil,
			expectedErr: ErrRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.srv.Fetch(tc.id)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
