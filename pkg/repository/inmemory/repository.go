package inmemory

type Repository struct {
	briefToFull map[string]string
	fullToBrief map[string]string
}

func NewRepository() *Repository {
	briefToFull := make(map[string]string)
	fullToBrief := make(map[string]string)
	return &Repository{
		briefToFull: briefToFull,
		fullToBrief: fullToBrief,
	}
}
