package mock

type MockedData struct {
	Email    string `faker:"email"`
	Title    string `faker:"sentence"`
	Content  string `faker:"paragraph"`
	DeviceID string `faker:"uuid_digit"`
	Phone    string `faker:"phone_number"`
}
