package drivens

import (
    "context"
    "encoding/json"
    "errors"
    "io/ioutil"
    "math/rand"
    "net/http"
    "os"
    "time"
    "youtube-api/models"

    "google.golang.org/api/googleapi/transport"
    "google.golang.org/api/youtube/v3"
)

type YouTubeAdapter struct {
    apiKey       string
    channelIDs   []string
    maxResults   int64
    service      *youtube.Service
}

func NewYouTubeAdapter() (*YouTubeAdapter, error) {
    apiKey := os.Getenv("YOUTUBE_API_KEY")
    channelIDs := os.Getenv("YOUTUBE_CHANNEL_IDS")
    maxResults := os.Getenv("YOUTUBE_MAX_RESULTS")

    if apiKey == "" || channelIDs == "" || maxResults == "" {
        return nil, errors.New("missing required environment variables")
    }

    maxResultsInt, err := strconv.ParseInt(maxResults, 10, 64)
    if err != nil {
        return nil, err
    }

    channels := strings.Split(channelIDs, ",")

    client := &http.Client{
        Transport: &transport.APIKey{Key: apiKey},
    }

    service, err := youtube.New(client)
    if err != nil {
        return nil, err
    }

    return &YouTubeAdapter{
        apiKey:     apiKey,
        channelIDs: channels,
        maxResults: maxResultsInt,
        service:    service,
    }, nil
}

func (yt *YouTubeAdapter) GetSearching(ctx context.Context, searching string) ([]models.DataYT, error) {
    call := yt.service.Search.List([]string{"id", "snippet"}).
        Q(searching).
        MaxResults(yt.maxResults)

    response, err := call.Do()
    if err != nil {
        return nil, err
    }

    var results []models.DataYT
    for _, item := range response.Items {
        results = append(results, models.DataYT{
            ID:          item.Id.VideoId,
            Title:       item.Snippet.Title,
            Description: item.Snippet.Description,
            URL:         "https://www.youtube.com/watch?v=" + item.Id.VideoId,
        })
    }
    return results, nil
}

func (yt *YouTubeAdapter) GetRandomVideo(ctx context.Context) (models.DataYT, error) {
    call := yt.service.Search.List([]string{"id", "snippet"}).
        MaxResults(yt.maxResults)

    response, err := call.Do()
    if err != nil {
        return models.DataYT{}, err
    }

    rand.Seed(time.Now().Unix())
    randomIndex := rand.Intn(len(response.Items))

    item := response.Items[randomIndex]
    return models.DataYT{
        ID:          item.Id.VideoId,
        Title:       item.Snippet.Title,
        Description: item.Snippet.Description,
        URL:         "https://www.youtube.com/watch?v=" + item.Id.VideoId,
    }, nil
}

func (yt *YouTubeAdapter) GetSearchPlaylists(ctx context.Context) ([]models.DataYT, error) {
    call := yt.service.Playlists.List([]string{"id", "snippet"}).
        MaxResults(yt.maxResults).
        ChannelId(yt.channelIDs[0]) // assuming first channel for example

    response, err := call.Do()
    if err != nil {
        return nil, err
    }

    var results []models.DataYT
    for _, item := range response.Items {
        results = append(results, models.DataYT{
            ID:          item.Id,
            Title:       item.Snippet.Title,
            Description: item.Snippet.Description,
            URL:         "https://www.youtube.com/playlist?list=" + item.Id,
        })
    }
    return results, nil
}

func (yt *YouTubeAdapter) GetSearchPlaylistsItems(ctx context.Context, playlistID string) ([]models.DataYT, error) {
    call := yt.service.PlaylistItems.List([]string{"id", "snippet"}).
        MaxResults(yt.maxResults).
        PlaylistId(playlistID)

    response, err := call.Do()
    if err != nil {
        return nil, err
    }

    var results []models.DataYT
    for _, item := range response.Items {
        results = append(results, models.DataYT{
            ID:          item.Snippet.ResourceId.VideoId,
            Title:       item.Snippet.Title,
            Description: item.Snippet.Description,
            URL:         "https://www.youtube.com/watch?v=" + item.Snippet.ResourceId.VideoId,
        })
    }
    return results, nil
}

func (yt *YouTubeAdapter) GetCheckLive(ctx context.Context) (models.Live, error) {
    call := yt.service.LiveBroadcasts.List([]string{"id", "snippet", "contentDetails", "status"}).
        BroadcastStatus("active").
        BroadcastType("all").
        MaxResults(yt.maxResults)

    response, err := call.Do()
    if err != nil {
        return models.Live{}, err
    }

    if len(response.Items) > 0 {
        item := response.Items[0]
        return models.Live{
            ID:    item.Id,
            Title: item.Snippet.Title,
            URL:   "https://www.youtube.com/watch?v=" + item.Id,
            Live:  true,
        }, nil
    }

    return models.Live{
        Live: false,
    }, nil
}
