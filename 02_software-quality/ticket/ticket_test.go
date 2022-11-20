package ticket

import "testing"

type testCase struct {
	name string
	age  int
	want float64
}

func TestTicketPrice(t *testing.T) {

	tests := []testCase{
		{name: "should return 0 when age is 0", age: 0, want: 0.0},
		{name: "free Ticket when age under 3", age: 3, want: 0.0},
		{name: "Ticket $15 when age at 4 year old", age: 4, want: 15.0},
		{name: "Ticket $15 when age at 15 year old", age: 15, want: 15.0},
		{name: "Ticket $30 when age over 15", age: 16, want: 30.0},
		{name: "Ticket $30 when age is 50", age: 50, want: 30.0},
		{name: "Ticket $5 when age over 50", age: 51, want: 5.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Price(uint(tc.age))

			if got != tc.want {
				t.Errorf("Price(%d) = %f; want %f", tc.age, got, tc.want)
			}
		})
	}
}
