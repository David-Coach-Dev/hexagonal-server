package drivers

import (
    "context"
    "youtube-api/models"
    "youtube-api/ports/drivens"
)

type YouTubeDriverAdapter struct {
    drivenAdapter drivens.YouTubeDrivenPort
}

func NewYouTubeDriverAdapter(drivenAdapter drivens.YouTubeDrivenPort) *YouTubeDriverAdapter {
    return &YouTubeDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

func (yda *YouTubeDriverAdapter) GetSearching(ctx context.Context, searching string) ([]models.DataYT, error) {
    return yda.drivenAdapter.GetSearching(ctx, searching)
}

func (yda *YouTubeDriverAdapter) GetRandomVideo(ctx context.Context) (models.DataYT, error) {
    return yda.drivenAdapter.GetRandomVideo(ctx)
}

func (yda *YouTubeDriverAdapter) GetSearchPlaylists(ctx context.Context) ([]models.DataYT, error) {
    return yda.drivenAdapter.GetSearchPlaylists(ctx)
}

func (yda *YouTubeDriverAdapter) GetSearchPlaylistsItems(ctx context.Context, playlistID string) ([]models.DataYT, error) {
    return yda.drivenAdapter.GetSearchPlaylistsItems(ctx, playlistID)
}

func (yda *YouTubeDriverAdapter) GetCheckLive(ctx context.Context) (models.Live, error) {
    return yda.drivenAdapter.GetCheckLive(ctx)
}
