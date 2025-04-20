package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/omikuu/sr/domain/video_info"
	"github.com/omikuu/sr/infrastructure/reoisitory_imple"
	video_usecase "github.com/omikuu/sr/usecase/video"
)

func main() {
	apiKey, title, limit, afterDays := loadEnvVariables()
	yt, err := reoisitory_imple.NewYouTubeClient(apiKey)
	if err != nil {
		log.Fatalf("YouTube client error: %v", err)
	}

	usecase := video_usecase.FetchVideosUseCase{Repo: yt}

	videos, err := usecase.Execute(title, limit, afterDays)
	if err != nil {
		log.Fatalf("Failed to fetch videos: %v", err)
	}

	// 再生回数順にソート
	videos = sortByViewCount(videos)

	printVideos(videos)
}

func loadEnvVariables() (string, string, int, int) {
	_ = godotenv.Load()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY not found")
	}

	title := os.Getenv("YOUTUBE_SEARCH_TITLE")
	if title == "" {
		log.Fatal("YOUTUBE_SERCH_TITLE not found")
	}

	limitStr := os.Getenv("YOUTUBE_SEARCH_LIMIT")
	if limitStr == "" {
		log.Fatal("YOUTUBE_SEARCH_LIMIT not found")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatalf("Invalid YOUTUBE_SEARCH_LIMIT: %v", err)
	}

	afterDaysStr := os.Getenv("YOUTUBE_SEARCH_AFTER_DAYS")
	if afterDaysStr == "" {
		log.Fatal("YOUTUBE_SEARCH_AFTER_DAYS not found")
	}
	afterDays, err := strconv.Atoi(afterDaysStr)
	if err != nil {
		log.Fatalf("Invalid YOUTUBE_SEARCH_AFTER_DAYS: %v", err)
	}
	return apiKey, title, limit, afterDays
}

func sortByViewCount(videos []video_info.VideoInfo) []video_info.VideoInfo {
	sort.Slice(videos, func(i, j int) bool {
		return videos[i].ViewCount > videos[j].ViewCount
	})
	return videos
}

func printVideos(videos []video_info.VideoInfo) {
	fmt.Println("🎬 最新動画リスト:")
	for _, v := range videos {
		fmt.Println(v.URL)
	}
}
