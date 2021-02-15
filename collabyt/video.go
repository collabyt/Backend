package collabyt

type videoRepository interface {
}

type Video struct {
	repo videoRepository
}

func NewVideo(r videoRepository) *Video {
	return &Video{r}
}
