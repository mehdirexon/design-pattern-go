package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id , breed , weight_low_lbs, weight_high_lbs, cast(((weight_low_lbs + weight_high_lbs) /2) as unsigned) as average_weight,
																		lifespan, coalesce(details, ''),
																		coalesce(alternate_names, ''), coalesce(geographic_origin, '')
																		from dog_breeds order by breed`

	var dogBreeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dogBreed DogBreed
		err := rows.Scan(
			&dogBreed.ID,
			&dogBreed.Breed,
			&dogBreed.WeightLowLbs,
			&dogBreed.WeightHighLbs,
			&dogBreed.AverageWeight,
			&dogBreed.Lifespan,
			&dogBreed.Details,
			&dogBreed.AlternateNames,
			&dogBreed.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		dogBreeds = append(dogBreeds, &dogBreed)
	}

	return dogBreeds, nil
}

func (m *mysqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id , breed , weight_low_lbs, weight_high_lbs, cast(((weight_low_lbs + weight_high_lbs) /2) as unsigned) as average_weight,
																		lifespan, coalesce(details, ''),
																		coalesce(alternate_names, ''), coalesce(geographic_origin, '')
																		from dog_breeds where breed = ?`

	row := m.DB.QueryRowContext(ctx, query, b)

	var dogBreed DogBreed

	err := row.Scan(
		&dogBreed.ID,
		&dogBreed.Breed,
		&dogBreed.WeightLowLbs,
		&dogBreed.WeightHighLbs,
		&dogBreed.AverageWeight,
		&dogBreed.Lifespan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)

	if err != nil {
		log.Println("error getting dogBreed:", err)
		return nil, err
	}
	return &dogBreed, nil
}
