package global_variable_jit

// Global singleton of connections to our data store
var storage UserStorage

type Saver struct {
	storage UserStorage
}

func (s *Saver) Do(in *User) error {
	err := s.validate(in)
	if err != nil {
		return err
	}

	return s.getStorage().Save(in)
}

// Just-in-time DI
func (s *Saver) getStorage() UserStorage {
	if s.storage == nil {
		s.storage = storage
	}

	return s.storage
}

func (s *Saver) validate(in *User) error {
	// validate user and return error when there is a problem
	return nil
}

type UserStorage interface {
	Save(in *User) error
}

type User struct {
	Name     string
	Password string
}
