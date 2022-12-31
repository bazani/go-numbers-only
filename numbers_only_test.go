package numbers_only

import "testing"

func TestNumbersOnly(t *testing.T) {
	t.Run("removing other chars from input", func(t *testing.T) {
		got := onlyNumbers("123.456 789 10")
		want := "12345678910"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but wanted %q", got, want)
	}
}
