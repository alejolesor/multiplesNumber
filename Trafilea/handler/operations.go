package handler

import (
	"TRAFILEA/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

type Operations struct {
	servicesMultipliers usecases.ServicesMultipliers
}

func NewOperations(servicesMultipliers usecases.ServicesMultipliers) *Operations {
	return &Operations{
		servicesMultipliers: servicesMultipliers,
	}
}

func (b *Operations) GetMultipliers(w http.ResponseWriter, r *http.Request) {

	num, _ := strconv.Atoi(chi.URLParam(r, "num"))
	result := b.servicesMultipliers.GetMultipliers(num)

	marshal, _ := json.Marshal(result)

	_, err := w.Write(marshal)
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (b *Operations) SaveNumberCollection(w http.ResponseWriter, r *http.Request) {

	num, _ := strconv.Atoi(chi.URLParam(r, "num"))
	b.servicesMultipliers.SaveMultiplier(num)

	_, err := w.Write([]byte(fmt.Sprintf("num created : %d", num)))
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *Operations) GetValueByNumber(w http.ResponseWriter, r *http.Request) {

	num, _ := strconv.Atoi(chi.URLParam(r, "num"))
	result, err := b.servicesMultipliers.GetValueByNumber(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	marshal, _ := json.Marshal(result)

	_, err = w.Write(marshal)
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (b *Operations) GetCollection(w http.ResponseWriter, r *http.Request) {

	result := b.servicesMultipliers.GetMultiplierCollection()
	marshal, _ := json.Marshal(result)
	_, err := w.Write(marshal)
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
