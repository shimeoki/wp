package app

type ImportBackupCommand struct {
	Backup *BackupResult
}

type ImportBackupResult struct{}

func (s *BackupService) Import(
	Ctx,
	*ImportBackupCommand,
) (*ImportBackupResult, error) {
	return nil, nil // TODO: implement
}
