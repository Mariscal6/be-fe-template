package rickmortyapi

import (
	"context"
	"sync"

	rnm "github.com/pitakill/rickandmortyapigowrapper"
	"golang.org/x/sync/errgroup"
)

type RickMortyAPI struct{}

// The API will automatically paginate the responses.
// You will receive up to 20 documents per page.
func (api RickMortyAPI) GetCharacters(ctx context.Context, page int) ([]rnm.Character, error) {
	g := errgroup.Group{}
	m := sync.Mutex{}
	allCharacters := []rnm.Character{}
	apiPage := page * 2
	g.Go(func() error {
		characters, err := rnm.GetCharacters(map[string]interface{}{"page": apiPage})
		if err != nil {
			return err
		}
		m.Lock()
		allCharacters = append(allCharacters, characters.Results...)
		return nil
	})

	g.Go(func() error {
		characters, err := rnm.GetCharacters(map[string]interface{}{"page": apiPage + 1})
		if err != nil {
			return err
		}
		m.Lock()
		allCharacters = append(allCharacters, characters.Results...)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return allCharacters, nil
}
