package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
)

var tpl *template.Template
var staticDir string

func init() {

	_, thisFile, _, _ := runtime.Caller(0)
	base := filepath.Dir(thisFile)

	htmlPath := filepath.Join(base, "visuel", "index.html")
	staticDir = filepath.Join(base, "style")

	tpl = template.Must(template.ParseFiles(htmlPath))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tpl.Execute(w, map[string]any{
		"Board":   board,
		"Current": current,
		"Winner":  winner,
		"Draw":    draw,
	})
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if winner != 0 || draw {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err == nil {
		if drop(col, current) {
			winner = checkWin()
			if winner == 0 && isFull() {
				draw = true
			}
			if winner == 0 && !draw {
				if current == 1 {
					current = 2
				} else {
					current = 1
				}
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	reset()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	reset()
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/reset", resetHandler)

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.ListenAndServe(":8080", nil)
}
