package video

type Repository interface {
	GetVideoUrl(title string, limit int) ([]Video, error)
}
