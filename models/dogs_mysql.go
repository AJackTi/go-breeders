package models

import (
	"context"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
				id,
				breed,
				weight_low_lbs,
				weight_high_lbs,
				CAST(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) AS average_weight,
				lifespan,
				COALESCE(details, ''),
				COALESCE(alternate_names, ''),
				COALESCE(geographic_origin, '')
				FROM dog_breeds ORDER BY breed;`

	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.AverageWeight,
			&b.Lifespan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}

	return breeds, nil
}
