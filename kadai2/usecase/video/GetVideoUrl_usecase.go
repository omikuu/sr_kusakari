package video_usecase

import (
	"github.com/omikuu/sr/domain/video"
	"github.com/omikuu/sr/domain/video_info"
)

type FetchVideosUseCase struct {
	Repo video.Repository
}

func (u *FetchVideosUseCase) Execute(title string, limit int, afterDays int) ([]video_info.VideoInfo, error) {
	return u.Repo.GetVideoUrl(title, limit, afterDays)
}
