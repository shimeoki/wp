package app

type AddWallpaperTagCommand struct {
	WallpaperHash Hash
	TagName       string
}

type AddWallpaperTagResult struct{}

func (s *WallpaperService) AddTag(
	Ctx,
	*AddWallpaperTagCommand,
) (*AddWallpaperTagResult, error) {
	return nil, nil // TODO: implement
}
