package server

import (
	"encoding/json"
	"net/http"
)

// Response contains the http response body data
type Response struct {
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Lifespan    string `json:"lifespan"`
	Temperament string `json:"temperament"`
	BreedGroup  string `json:"breed_group"`
}

// GetRandomDog returns random dog data from the DogAPI
func (s *AnimalAPIServer) GetRandomDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dogInfo, err := s.DogAPIClient.GetRandomDogInfo()
	if err != nil {
		s.ErrorLog.Printf("%s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		s.ErrorLog.Printf("%s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
