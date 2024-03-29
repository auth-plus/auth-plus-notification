// Package mock contains all mocks and factory/generator
package mock

// MockedData is used as factory contract for faker
type MockedData struct {
	Email    string `faker:"email"`
	Title    string `faker:"sentence"`
	Subject  string `faker:"sentence"`
	Content  string `faker:"paragraph"`
	DeviceID string `faker:"uuid_digit"`
	Phone    string `faker:"phone_number"`
	ChatID   int64  `faker:"unix_time"`
	Error    string `faker:"sentence"`
}
