package helpers

var preTagColors = []string{
	"#c0e4c5",
	"#b4d6db",
	"#a997ae",
	"#54c0f4",
	"#f1252f",
	"#9c90f4",
	"#67f5f5",
	"#3558b6",
	"#f0db4f",
	"#f00dca",
	"#ff8303",
	"#787cb5",
	"#e756c8",
	"#ffd43b",
	"#bc1142",
	"#0f0",
	"#49beb7",
	"#55ff55",
	"#5dacfd",
}

func RandomTagColors(idx int) string {
	if idx < len(preTagColors) {
		return preTagColors[idx]
	}
	for idx >= len(preTagColors) {
		idx = idx - len(preTagColors)
	}

	return preTagColors[idx]
}
