package collabyt

type errorRepository interface {
}

type Error struct {
	repo errorRepository
}

func NewError(r errorRepository) *Error {
	return &Error{r}
}
