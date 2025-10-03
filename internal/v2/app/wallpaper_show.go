package app

type ShowWallpaperQuery struct {
	Hash
}

type ShowWallpaperResult struct {
	Image
}

func (s *WallpaperService) Show(
	Ctx,
	*ShowWallpaperQuery,
) (*ShowWallpaperResult, error) {
	return nil, nil // TODO: implement
}
