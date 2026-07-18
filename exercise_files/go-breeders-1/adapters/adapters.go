package adapters

import (
	"encoding/json"
	"encoding/xml"
	"go-breeders/models"
	"io"
	"net/http"
)

// CatBreedsInterface is simply our target interface, which defines all the methods that
// any type which implements this interface must have.
type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
	GetCatBreedByName(b string) (*models.CatBreed, error)
}

// RemoteService is the Adaptor type. It embeds a DataInterface interface
// (which is critical to the pattern).
type RemoteService struct {
	Remote CatBreedsInterface
}

// GetAllBreeds is the function on RemoteService which lets us
// call any adaptor which implements the DataInterface type.
func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

// JSONBackend is the JSON adaptee, which needs to satisfy the CatBreedsInterface by
// have all required method.
type JSONBackend struct{}

// GetAllCatBreeds is necessary so that JSONBackend satisfies the CatBreedsInterface requirements.
func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

// GetCatBreedByName gets a cat breed by name from a remote service which
// returns JSON.
func (jb *JSONBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breed models.CatBreed

	err = json.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

// XMLBackend is a type which will satisfy our adapter requirements (CatBreedsInterface)
// by having all required methods.
type XMLBackend struct{}

// GetAllCatBreeds gets a slice of all cat breeds from a remote service which
// returns XML.
func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{} `xml:"cat-breeds"`
		Breeds []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds

	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil
}

// GetCatBreedByName gets a cat breed by name from a remote service which
// returns XML.
func (xb *XMLBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breed models.CatBreed

	err = xml.Unmarshal(body, &breed)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}


type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{ID: 1, Breed: "Tomcat", Details: "Some details"},
	}

	return breeds, nil
}

func (tb *TestBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	return nil, nil
}