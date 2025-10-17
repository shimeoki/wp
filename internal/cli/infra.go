package cli

import (
	"github.com/shimeoki/wp/internal/domain"
	"github.com/shimeoki/wp/internal/infra/store"
)

func openStore() domain.Store {
	hasher := &store.SHA256Hasher{}

	s, err := store.NewLocalStore(cfg.Store.Path, hasher)
	if err != nil {
		fatal(err)
	}

	return s
}
