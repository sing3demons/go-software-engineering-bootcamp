package double

import "testing"

type FakeSearch struct{}

func (fs FakeSearch) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}
	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringNoPerson(t *testing.T) {
	phoneBook := &PhoneBook{}
	fake := &FakeSearch{}

	phone, _ := phoneBook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}

}
