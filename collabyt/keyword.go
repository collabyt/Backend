package collabyt

type keywordRepository interface {
}

type Keyword struct {
	repo keywordRepository
}

func NewKeyword(r keywordRepository) *Keyword {
	return &Keyword{r}
}
