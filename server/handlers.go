package server

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Lifespan    string `json:"lifespan"`
	Temperament string `json:"temperament"`
	BreedGroup  string `json:"breed_group"`
}

func (s *Server) GetRandomDog(w http.ResponseWriter, r *http.Request) {
	dogInfo, err := s.Client.GetRandomDogInfo()
	if err != nil {
		s.ErrorLog.Printf("%s", err.Error())

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
		Height:      dog.Height.Metric + " cm",
		Weight:      dog.Weight.Metric + " kgs",
		Lifespan:    dog.LifeSpan,
		Temperament: dog.Temperament,
		BreedGroup:  dog.BreedGroup,
	}

	json.NewEncoder(w).Encode(resp)
}
