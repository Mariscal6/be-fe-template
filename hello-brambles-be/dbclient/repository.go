package dbclient

import (
	"context"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

type DbClient interface {
	GetCharacters(ctx context.Context, page int) ([]rnm.Character, error)
}
