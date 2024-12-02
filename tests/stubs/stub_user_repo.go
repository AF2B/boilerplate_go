package stubs

type StubUserRepository struct {
	Users map[string]string
}

func (s *StubUserRepository) GetUserByID(id string) string {
	return s.Users[id]
}
