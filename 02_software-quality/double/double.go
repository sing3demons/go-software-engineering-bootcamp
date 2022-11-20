package double

import "errors"

var (
	ErrMissingArgs   = errors.New("FirstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type PhoneBook struct{ People []*Person }

func (p *PhoneBook) Find(query Queryer, firstName string, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := query.Search(p.People, firstName, lastName)
	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
