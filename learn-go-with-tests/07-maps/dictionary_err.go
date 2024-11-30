package maps

const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}
