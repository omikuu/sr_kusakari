package video_usecase

import "github.com/omikuu/sr/domain/video"

type FetchVideosUseCase struct {
	Repo video.Repository
}

func (u *FetchVideosUseCase) Execute(title string, limit int) ([]video.Video, error) {
	return u.Repo.GetVideoUrl(title, limit)
}
