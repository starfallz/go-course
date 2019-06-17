package main

import (
	"errors"
	"fmt"
)

func main() {
	s := &MapStore{}

	puppy1 := Puppy{111, "Pug", "Black", 3000.00};
	puppy2 := Puppy{120, "Poodle", "White", 1000.00};
	updatedPuppy1 := Puppy{111, "Pug", "Brown", 3333.33};

	fmt.Println(s.CreatePuppy(&puppy1))
	fmt.Println(s.ReadPuppy(111))
	fmt.Println(s.UpdatePuppy(111, &puppy2))
	fmt.Println(s.UpdatePuppy(111, &updatedPuppy1))
	fmt.Println(s.DeletePuppy(111))
}

type MapStore map[uint32]*Puppy

type Puppy struct {
	ID      uint32                 	`json:"id"`
	Breed   string                	`json:"breed"`
	Colour 	string                 	`json:"colour"`
	Value   float64                 `json:"value"`
}

type Storer interface {
	CreatePuppy(puppy *Puppy) error
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, puppy *Puppy) error
	DeletePuppy(ID uint32) (bool, error)
}

//type SyncStore struct {}

func (s *MapStore) CreatePuppy(puppy *Puppy) error {
	if (*s)[puppy.ID] == nil {
		(*s)[puppy.ID] = puppy
		return nil
	} else {
		return errors.New(fmt.Sprintf("ERROR: Create failed, puppy with ID %d already exist", puppy.ID))
	}
}

func (s *MapStore) ReadPuppy(ID uint32) (*Puppy, error) {
	if (*s)[ID] != nil {
		return (*s)[ID], nil
	} else {
		return nil, errors.New(fmt.Sprintf("ERROR: Read failed, puppy with ID %d not found", ID))
	}
}

func (s *MapStore) UpdatePuppy(ID uint32, puppy *Puppy) error {
	if (*s)[ID] == nil {
		return errors.New(fmt.Sprintf("ERROR: Update failed, puppy with ID %d does not exist", ID))
	} else if (*s)[ID].ID != puppy.ID {
		return errors.New(fmt.Sprintf("ERROR: Update failed, puppy with ID %d, %d does not match", ID, puppy.ID))
	} else {
		(*s)[ID] = puppy
		return nil
	}
}

func (s *MapStore) DeletePuppy(ID uint32) (bool, error) {
	if (*s)[ID] != nil {
		(*s)[ID] = nil
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("ERROR: Delete failed, puppy with ID %d does not exist", ID))
	}
}
