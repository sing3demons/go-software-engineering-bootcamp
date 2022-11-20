package double

import "testing"

type SpySearch struct {
	phone             string
	searchWatchCalled bool
}

func (ss *SpySearch) Search(people []*Person, firstName string, lastName string) *Person {
	ss.searchWatchCalled = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+66 965 436 633"
	phoneBook := &PhoneBook{}

	spy := &SpySearch{phone: fakePhone}

	phone, _ := phoneBook.Find(spy, "Jane", "Doe")

	if !spy.searchWatchCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
