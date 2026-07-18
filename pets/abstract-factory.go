package pets

import (
	"breeders/configuration"
	"breeders/models"
	"fmt"
	"log"
)

type AnimalInterface interface {
	Show() string
}
type DogFromFactory struct {
	Pet *models.Dog
}

func (f *DogFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", f.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (f *CatFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", f.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactory struct {
}

func (f *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{Pet: &models.Dog{}}
}

func (f *DogAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := configuration.GeInstance()
	breed, _ := app.Models.DogBreed.GetBreedByName(b)

	return &DogFromFactory{Pet: &models.Dog{Breed: *breed}}
}

type CatAbstractFactory struct {
}

func (f *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{Pet: &models.Cat{}}
}

func (f *CatAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := configuration.GeInstance()
	breed, err := app.CatService.Remote.GetCatBreedByName(b)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &CatFromFactory{Pet: &models.Cat{
		Breed: *breed,
	}}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, fmt.Errorf("unknown species %s", species)
	}
}

func NewPetWithBreedFromAbstractFactory(species string, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)

		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPetWithBreed(breed)

		return cat, nil
	default:
		return nil, fmt.Errorf("unknown species %s", species)
	}
}
