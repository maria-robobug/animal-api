package model

type DogInfo struct {
	Breeds []Breed `json:"breeds"`
	URL    string  `json:"url"`
}

type Breed struct {
	Name        string  `json:"name"`
	Height      Measure `json:"height"`
	Weight      Measure `json:"weight"`
	BreedGroup  string  `json:"breed_group"`
	LifeSpan    string  `json:"life_span"`
	Temperament string  `json:"temperament"`
}

type Measure struct {
	Metric string `json:"metric"`
}
