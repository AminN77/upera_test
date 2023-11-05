package internal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchRevision(t *testing.T) {
	testCases := []struct {
		name           string
		revisionNumber string
		srv            Service
		expected       *Product
		expectedErr    error
	}{
		{
			name:           "success",
			srv:            NewService(&mockRepository{}),
			revisionNumber: "some revision number",
			expected:       &Product{},
			expectedErr:    nil,
		},
		{
			name:           "repo err",
			srv:            NewService(&errMockRepository{}),
			revisionNumber: "some revision number",
			expected:       nil,
			expectedErr:    ErrRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.srv.FetchRevision(tc.revisionNumber, context.Background())
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestFetchRevisionsOfOneProduct(t *testing.T) {
	testCases := []struct {
		name        string
		productId   int64
		srv         Service
		expectedErr error
	}{
		{
			name:        "success",
			srv:         NewService(&mockRepository{}),
			productId:   17,
			expectedErr: nil,
		},
		{
			name:        "repo err",
			srv:         NewService(&errMockRepository{}),
			productId:   17,
			expectedErr: ErrRepo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.srv.FetchRevisionsOfOneProduct(10, 10, tc.productId, context.Background())
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
