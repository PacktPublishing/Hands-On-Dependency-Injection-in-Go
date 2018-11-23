package humans

// PetFetcher searches the data store for pets whos name matches the search string.
// Limit is optional (default is 100).  Offset is optional (default 0).
// sortBy is optional (default name).  sortAscending is optional
func PetFetcher(search string, limit int, offset int, sortBy string, sortAscending bool) []Pet {
	return []Pet{}
}

func PetFetcherTypicalUsage() {
	_ = PetFetcher("Fido", 0, 0, "", true)
}
