package collabyt

type authRepository interface {
}

type Auth struct {
	repo authRepository
}

func NewAuth(r authRepository) *Auth {
	return &Auth{r}
}
