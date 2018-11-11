package list

import (
	"context"
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/data"
)

var (
	// error thrown when there are no people in the database
	errPeopleNotFound = errors.New("no people found")
)

// NewLister creates and initializes a Lister
func NewLister(cfg Config) *Lister {
	return &Lister{
		cfg: cfg,
	}
}

// Config is the config for Lister
type Config interface {
	Logger() logging.Logger
	DataDSN() string
}

// Lister will attempt to load all people in the database.
// It can return an error caused by the data layer
type Lister struct {
	cfg  Config
	data myLoader
}

// Exchange will load the people from the data layer
func (l *Lister) Do() ([]*Person, error) {
	// load all people
	people, err := l.load()
	if err != nil {
		return nil, err
	}

	if len(people) == 0 {
		// special processing for 0 people returned
		return nil, errPeopleNotFound
	}

	return l.convert(people), nil
}

// load all people
func (l *Lister) load() ([]*data.Person, error) {
	people, err := l.getLoader().LoadAll(context.TODO())
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPeopleNotFound
		}
		return nil, err
	}

	return people, nil
}

func (l *Lister) getLoader() myLoader {
	if l.data == nil {
		l.data = data.NewDAO(l.cfg)

		// temporarily add a log tracker
		l.data.(*data.DAO).Tracker = data.NewLogTracker(l.cfg.Logger())
	}

	return l.data
}

func (l *Lister) convert(in []*data.Person) []*Person {
	out := make([]*Person, len(in))

	for index, thisRecord := range in {
		out[index] = &Person{
			ID:       thisRecord.ID,
			FullName: thisRecord.FullName,
			Phone:    thisRecord.Phone,
		}
	}

	return out
}

//go:generate mockery -name=myLoader -case underscore -testonly -inpkg -note @generated
type myLoader interface {
	LoadAll(ctx context.Context) ([]*data.Person, error)
}

// Person is a copy/sub-set of data.Person so that the relationship does not leak.
// It also allows us to remove/hide and internal fields
type Person struct {
	ID       int
	FullName string
	Phone    string
}
