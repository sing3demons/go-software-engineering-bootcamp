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
}
