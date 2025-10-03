package app

type DeleteWallpaperCommand struct {
	Hash
}

type DeleteWallpaperResult struct{}

func (s *WallpaperService) Delete(
	Ctx,
	*DeleteWallpaperCommand,
) (*DeleteWallpaperResult, error) {
	return nil, nil // TODO: implement
}
