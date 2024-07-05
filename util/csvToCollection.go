package util

func CsvToCollection(rows [][]string) map[string]string {
	table := make(map[string]string)

	for _, row := range rows {
		kanji := row[1]
		meaning := row[0]
		table[kanji] = meaning
	}

	return table
}
