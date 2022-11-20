package double

import "testing"

// Stub return ของ เฉยๆ
type StubSearch struct{ phone string }

func (ss StubSearch) Search(people *[]Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+66 965 436 633"
	phoneBook := &PhoneBook{}

	phone, _ := phoneBook.Find(StubSearch{fakePhone}, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("")
	}

}
