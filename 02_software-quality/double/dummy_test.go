package double

import "testing"


// Dummy ส่งของไปให้พารามิตเตอร์ผ่าน
type DummySearch struct{}

func (DummySearch) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestFindIsShouldReturnsErrorWhenFirstNameOrLastName(t *testing.T) {
	phoneBook := &PhoneBook{}
	want := ErrMissingArgs
	_, got := phoneBook.Find(DummySearch{}, "", "")

	if got != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}
