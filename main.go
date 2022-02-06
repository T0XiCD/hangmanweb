package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Hangman struct {
	word         string
	guessesLeft  int
	wrongGuesses []string
	state        bool
	guess        string
	draw         []string
	show         []string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		data := Hangman{
			word:        r.FormValue("w"),
			guessesLeft: 10,
			state:       true,
			guess:       "bonsoir",
			draw: []string{"  +-----+\n  |     |\n	|\n	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n  |	|\n	|\n	|\n========="},
		}
		if len(data.word) == 0 {
			fmt.Fprintln(w, "Il n'ya pas d'argument envoyer")
			return
		}
		if len(data.word) == 1 {
			for i := 0; i < len(data.guess); i++ {
				if data.word[0] == data.guess[i] {
					fmt.Fprintln(w, "Tu as trouver une lettre")
					data.show = append(data.show, data.word)
				} else {
					data.guessesLeft -= 1
					data.show = append(data.show, data.word)
					data.wrongGuesses = append(data.wrongGuesses, data.word)
					fmt.Fprintln(w, "Dommage mauvaise lettre")
					fmt.Fprintln(w, "tu as perdue 1pv il te reste", data.guessesLeft, "pv")
				}
				break
			}
		}
		if len(data.word) >= 2 {
			data.guessesLeft -= 2
			data.wrongGuesses = append(data.wrongGuesses, data.word)
			data.show = append(data.show, data.word)
			fmt.Fprintln(w, "sale merde")
		}
		if data.guessesLeft == 10 {
			fmt.Fprintln(w, data.draw[0])
			fmt.Fprintln(w, "Mauvais mot déja utiliser : ", data.wrongGuesses)
			fmt.Fprintln(w, "Tout les mots déja utiliser : ", data.show)
		}
		if data.guessesLeft == 9 {
			fmt.Fprintln(w, data.draw[1])
			fmt.Fprintln(w, "Mauvais mot déja utiliser : ", data.wrongGuesses)
			fmt.Fprintln(w, "Tout les mots déja utiliser : ", data.show)
		}
		if data.guessesLeft == 8 {
			fmt.Fprintln(w, data.draw[2])
			fmt.Fprintf(w, "ou la la mince alors")
			fmt.Fprintln(w, "Mauvais mot déja utiliser : ", data.wrongGuesses)
			fmt.Fprintln(w, "Tout les mots déja utiliser : ", data.show)
		}
		if data.guessesLeft == 7 {
			fmt.Fprintf(w, "tu es vraiment dans la mouisse")
		}
		tmpl.Execute(w, data)
	})
	// premier paramètre string url, après le port
	http.ListenAndServe(":8000", nil)

}
