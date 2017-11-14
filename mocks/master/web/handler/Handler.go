// Code generated by mockery v1.0.0
package serializer

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Debug provides a mock function with given fields: w, r
func (_m *Handler) Debug(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// WriteExperiment provides a mock function with given fields: w, r
func (_m *Handler) WriteExperiment(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
