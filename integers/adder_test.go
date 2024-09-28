package integers

import (
	"fmt"
	"testing"
)

func Test_adder(t *testing.T) {
	t.Run("adding two integer values", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
