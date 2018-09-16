package global_variable

// Global singleton of connections to our data store
var storage UserStorage

type Saver struct {
}

func (s *Saver) Do(in *User) error {
	err := s.validate(in)
	if err != nil {
		return err
	}

	return storage.Save(in)
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
