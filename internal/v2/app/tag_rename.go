package app

type RenameTagCommand struct {
	Before string
	After  string
}

type RenameTagResult struct{}

func (s *TagService) Rename(Ctx, *RenameTagCommand) (*RenameTagResult, error) {
	return nil, nil // TODO: implement
}
