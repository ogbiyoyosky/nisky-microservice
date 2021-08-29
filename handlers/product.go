package handlers

import (
	"di/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	productSlice, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(productSlice)
}
