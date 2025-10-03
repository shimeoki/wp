package app

type RemoveWallpaperTagCommand struct {
	WallpaperHash Hash
	TagName       string
}

type RemoveWallpaperTagResult struct{}

func (s *WallpaperService) RemoveTag(
	Ctx,
	*RemoveWallpaperTagCommand,
) (*RemoveWallpaperTagResult, error) {
	return nil, nil // TODO: implement
}
