package server

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response contains the http response body data
type Response struct {
	Image       Image  `json:"image"`
	Name        string `json:"name"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Lifespan    string `json:"lifespan"`
	Temperament string `json:"temperament"`
	BreedGroup  string `json:"breed_group"`
}

type 

// Image holds Image information for a Dog
type Image struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

// GetHealth returns OK status
func (s *AnimalAPIServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := struct{ Status string }{http.StatusText(http.StatusOK)}

	render.JSON(w, r, resp)
}

// GetRandomDog returns random dog data from the DogAPI
func (s *AnimalAPIServer) GetRandomDog(w http.ResponseWriter, r *http.Request) {
	dogInfo, err := s.DogAPIClient.GetRandomDogInfo()
	if err != nil {
		s.Logger.Errorf("%s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dogImage := Image{
		URL:    dogInfo[0].URL,
		Width:  dogInfo[0].Width,
		Height: dogInfo[0].Height,
	}

	dog := dogInfo[0].Breeds[0]

	resp := &Response{
		Image:       dogImage,
		Name:        dog.Name,
		Height:      dog.Height.Metric + " cm",
		Weight:      dog.Weight.Metric + " kgs",
		Lifespan:    dog.LifeSpan,
		Temperament: dog.Temperament,
		BreedGroup:  dog.BreedGroup,
	}

	render.JSON(w, r, resp)
}

// GetRandomCat returns random cat data from the CatAPI
func (s *AnimalAPIServer) GetRandomCat(w http.ResponseWriter, r *http.Request) {
}
