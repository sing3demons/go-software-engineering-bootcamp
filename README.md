# software-engineering-bootcamp

```all test case
go test -v ./...
```

```check test cover ticket module
go test -cover ./ticket
go test -cover -coverprofile=c.out ./ticket
go tool cover -html=c.out -o coverage.html
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

## test library -> testify

```
go get -u github.com/stretchr/testify
```

## test library -> is
```
go get -u github.com/matryer/is
```