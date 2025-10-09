package app

import "github.com/shimeoki/wp/internal/v2/domain"

type BackupService struct {
	store domain.Store
}

func NewBackupService(s domain.Store) *BackupService {
	return &BackupService{
		store: s,
	}
}

type BackupResult struct {
	Version
	// wallpapers
}
