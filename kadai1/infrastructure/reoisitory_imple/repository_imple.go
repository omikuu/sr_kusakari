package reoisitory_imple

import (
	"context"
	"fmt"

	"github.com/omikuu/sr/domain/video"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeClient struct {
	service *youtube.Service
}

func NewYouTubeClient(apiKey string) (*YouTubeClient, error) {
	ctx := context.Background()
	srv, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &YouTubeClient{service: srv}, nil
}

func (c *YouTubeClient) GetVideoUrl(title string, limit int) ([]video.Video, error) {
	call := c.service.Search.List([]string{"id", "snippet"}).
		Q(title).
		Type("video").
		Order("date").
		MaxResults(int64(limit))

	res, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []video.Video
	for _, item := range res.Items {
		videos = append(videos, video.Video{
			URL: fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId),
		})
	}
	return videos, nil
}
