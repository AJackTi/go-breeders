package pets

import "errors"

type PetInterface interface {
	SetSpecies(string) *Pet
	SetBreed(string) *Pet
	SetMinWeight(int) *Pet
	SetMaxWeight(int) *Pet
	SetWeight(int) *Pet
	SetDescription(string) *Pet
	SetLifeSpan(int) *Pet
	SetGeographicOrigin(string) *Pet
	SetColor(string) *Pet
	SetAge(int) *Pet
	SetAgeEstimated(bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(b string) *Pet {
	p.Breed = b
	return p
}

func (p *Pet) SetMinWeight(m int) *Pet {
	p.MinWeight = m
	return p
}

func (p *Pet) SetMaxWeight(m int) *Pet {
	p.MaxWeight = m
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifeSpan(l int) *Pet {
	p.LifeSpan = l
	return p
}

func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(a bool) *Pet {
	p.AgeEstimated = a
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
