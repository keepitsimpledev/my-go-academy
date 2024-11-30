package maps

import "testing"

const testWord = "test"
const testDefinition = "this is just a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{testWord: testDefinition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(testWord)
		want := testDefinition

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := testDefinition

		err := dictionary.Add(testWord, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, testWord, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		definition := testDefinition
		dictionary := Dictionary{testWord: definition}
		err := dictionary.Add(testWord, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, testWord, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{testWord: testDefinition}
		newDefinition := "new definition"

		err := dictionary.Update(testWord, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, testWord, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(testWord, testDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{testWord: testDefinition}

	dictionary.Delete(testWord)

	_, err := dictionary.Search(testWord)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", testWord)
	}
}

// helpers

func assertDefinition(tb testing.TB, dictionary Dictionary, word, definition string) {
	tb.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		tb.Fatal("should find added word:", err)
	}

	assertStrings(tb, got, definition)
}

func assertStrings(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}
