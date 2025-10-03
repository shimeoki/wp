package app

type BackupService struct {
	store Store
}

func NewBackupService(s Store) *BackupService {
	return &BackupService{
		store: s,
	}
}

type BackupResult struct {
	Version
	// wallpapers
}
