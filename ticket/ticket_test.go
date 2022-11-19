package ticket

import "testing"

func TestTicketPrice(t *testing.T) {

	tests := []struct {
		name string
		age  int
		want float64
	}{
		{"should return 0 when age is 0", 0, 0.0},
		{"free Ticket when age under 3", 3, 0.0},
		{"Ticket $15 when age at 4 year old", 4, 15.0},
		{"Ticket $15 when age at 15 year old", 15, 15.0},
		{"Ticket $30 when age over 16", 16, 30.0},
		{"Ticket $5 when age over 50", 51, 5.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Price(uint(tc.age))

			if got != tc.want {
				t.Errorf("Price(%d) = %f; want %f", tc.age, got, tc.want)
			}
		})
	}
	// t.Run("should return 0 when age is 0", func(t *testing.T) {
	// 	want := 0.0
	// 	var age uint = 0

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(0) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("free Ticket when age under 3", func(t *testing.T) {
	// 	want := 0.0
	// 	var age uint = 3

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(3) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("Ticket $15 when age at 4 year old", func(t *testing.T) {
	// 	want := 15.0
	// 	var age uint = 4

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(4) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("Ticket $15 when age at 15 year old", func(t *testing.T) {
	// 	want := 15.0
	// 	var age uint = 15

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(15) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("Ticket $30 when age over 16", func(t *testing.T) {
	// 	want := 30.0
	// 	var age uint = 16

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(16) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("Ticket $30 when age over 50", func(t *testing.T) {
	// 	want := 30.0
	// 	var age uint = 50

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(50) = %f; want %f", got, want)
	// 	}
	// })

	// t.Run("Ticket $5 when age over 50", func(t *testing.T) {
	// 	want := 5.0
	// 	var age uint = 51

	// 	got := Price(age)

	// 	if got != want {
	// 		t.Errorf("Price(51) = %f; want %f", got, want)
	// 	}
	// })
}
