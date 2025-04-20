package reoisitory_imple

import (
	"context"
	"fmt"
	"time"

	"github.com/omikuu/sr/domain/video_info"
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

func (c *YouTubeClient) GetVideoUrl(title string, limit int, afterDays int) ([]video_info.VideoInfo, error) {
	// ISO形式
	publishedAfter := time.Now().AddDate(0, 0, afterDays).Format(time.RFC3339)

	call := c.service.Search.List([]string{"id", "snippet"}).
		Q(title).
		Type("video").
		RegionCode("JP").
		Order("date").
		PublishedAfter(publishedAfter).
		MaxResults(int64(limit))

	res, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videoIDs []string
	for _, item := range res.Items {
		videoIDs = append(videoIDs, item.Id.VideoId)
	}

	// 動画の詳細取得
	videoCall := c.service.Videos.List([]string{"snippet", "statistics"}).Id(videoIDs...)
	videoRes, err := videoCall.Do()
	if err != nil {
		return nil, err
	}

	var videos []video_info.VideoInfo
	for _, item := range videoRes.Items {
		viewCount := item.Statistics.ViewCount
		title := item.Snippet.Title
		videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id)

		videos = append(videos, video_info.VideoInfo{
			Title:     title,
			URL:       videoURL,
			ViewCount: int64(viewCount),
		})
	}

	return videos, nil
}
