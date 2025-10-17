package cli

import (
	"github.com/shimeoki/wp/internal/infra/store"
)

func openStore() *store.LocalStore {
	hasher := &store.SHA256Hasher{}

	s, err := store.NewLocalStore(cfg.Store.Path, hasher)
	if err != nil {
		fatal(err)
	}

	return s
}
