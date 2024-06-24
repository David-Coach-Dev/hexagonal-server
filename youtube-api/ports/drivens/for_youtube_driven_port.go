package drivens

import (
    "context"
    "youtube-api/models"
)

type YouTubeDrivenPort interface {
    GetSearching(ctx context.Context, searching string) ([]models.DataYT, error)
    GetRandomVideo(ctx context.Context) (models.DataYT, error)
    GetSearchPlaylists(ctx context.Context) ([]models.DataYT, error)
    GetSearchPlaylistsItems(ctx context.Context, playlistID string) ([]models.DataYT, error)
    GetCheckLive(ctx context.Context) (models.Live, error)
}
