package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type Response struct {
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Lifespan    string `json:"lifespan"`
	Temperament string `json:"temperament"`
}

func (app *Application) GetRandomDog(w http.ResponseWriter, r *http.Request) {
	dogInfo, err := app.Client.GetRandomDogInfo()
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.ErrorLog.Output(2, trace)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dogImage := dogInfo[0].URL
	dog := dogInfo[0].Breeds[0]
	resp := &Response{
		ImageURL:    dogImage,
		Name:        dog.Name,
		Height:      dog.Height.Metric,
		Weight:      dog.Weight.Metric,
		Lifespan:    dog.LifeSpan,
		Temperament: dog.Temperament,
	}

	json.NewEncoder(w).Encode(resp)
}
