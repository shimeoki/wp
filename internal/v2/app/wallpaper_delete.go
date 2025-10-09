package app

import "errors"

type DeleteWallpaperCommand struct {
	Hash
}

type DeleteWallpaperResult struct{}

func (s *WallpaperService) Delete(
	ctx Ctx,
	cmd *DeleteWallpaperCommand,
) (*DeleteWallpaperResult, error) {
	w, _ := s.wallpapers.ByHash(ctx, string(cmd.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	if err := s.wallpapers.Delete(ctx, w.ID); err != nil {
		return nil, err
	}

	if err := s.store.Remove(cmd.Hash); err != nil {
		return nil, err
	}

	return &DeleteWallpaperResult{}, nil
}
