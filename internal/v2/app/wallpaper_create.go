package app

type CreateWallpaperCommand struct {
	Image
}

type CreateWallpaperResult struct {
	Hash
}

func (s *WallpaperService) Create(
	Ctx,
	*CreateWallpaperCommand,
) (*CreateWallpaperResult, error) {
	return nil, nil // TODO: implement
}
