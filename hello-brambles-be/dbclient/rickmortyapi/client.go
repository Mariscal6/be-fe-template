package rickmortyapi

import rnm "github.com/pitakill/rickandmortyapigowrapper"

// The API will automatically paginate the responses.
// You will receive up to 20 documents per page.
func _() {
	_, _ = rnm.GetCharacters(map[string]interface{}{"page": 1})
}
