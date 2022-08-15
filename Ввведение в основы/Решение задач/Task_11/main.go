package main

import (
	"fmt"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	lang, _ := language.Parse("ru")
	_ = message.Set(language.Russian, "{count} korov",
		plural.Selectf(1, "",
			plural.One, "%d korova",
			plural.Few, "%d korovy",
			plural.Other, "%d korov",
		),
	)

	fmt.Print(message.NewPrinter(lang).Sprintf("{count} korov", n))
}
