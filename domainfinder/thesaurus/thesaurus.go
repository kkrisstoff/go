package thesaurus

// Thesaurus describes a method that takes a term string
// and returns either a slice of strings containing the synonyms or an error
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
