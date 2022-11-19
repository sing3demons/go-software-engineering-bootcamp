# software-engineering-bootcamp

```all test case
go test -v ./...
```

## Arrange, Act, Assert pattern

```example test pattern Go
func TestSum1(t *testing.T) {
	t.Run("should return 3 when 1 and 2", func(t *testing.T) {
		// Arrange
		want := 3

		// Act
		got := sum(1, 2)

		// Assert
		if got != want {
			t.Errorf("sum(1,2) = %d; want %d", got, want)
		}
	})
}
```
