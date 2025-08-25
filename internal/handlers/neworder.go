package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (h *Handler) NewOrder(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Type"), ContentTypeTextPlain) {

		err := fmt.Errorf("Content-Type != %s", ContentTypeTextPlain)
		errorGettingData(w, err)
		return
	}

	defer r.Body.Close()
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		errorGettingData(w, err)
		return
	}

	context := r.Context()
	orderID := string(responseData)
	err = h.Service.NewOrder(context, orderID)
	if err != nil {
		internalError(w, err)
		return
	}

}
