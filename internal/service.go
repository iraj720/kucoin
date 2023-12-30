package internal

func SymbolsTopics(syms *Symbols) []string {
	topics := make([]string, 0)
	for _, v := range syms.Data {
		topics = append(topics, v.Symbol)
	}
	return topics
}
