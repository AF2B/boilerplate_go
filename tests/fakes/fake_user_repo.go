package fakes

type FakeUserRepository struct {
	Users map[string]string
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{Users: make(map[string]string)}
}

func (f *FakeUserRepository) SaveUser(name string) error {
	f.Users[name] = "saved"
	return nil
}
