package app

type FindWallpaperQuery struct {
	Hash
}

type FindWallpaperResult struct{}

func (s *WallpaperService) Find(
	Ctx,
	*FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	return nil, nil // TODO: implement
}
