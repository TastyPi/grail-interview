package storage

import (
	"fmt"

	pb "github.com/TastyPi/grail-interview/api/participant"
)

// Interface for the Participant storage.
//
// Does not perform any validation on the Participant, just storage operations.
type ParticipantStorage interface {

	// Insert a participant into the store.
	//
	// The provide Participant must not have Name set. The returned Participant will be the same
	// as the one provided but with Name set.
	//
	// Returns an error if the provided Participant has a Name.
	Insert(pb.Participant) (*pb.Participant, error)

	// Retrieve a participant from the store.
	//
	// Returns ParticipantNotFoundError if the participant could not be found.
	Read(name string) (*pb.Participant, error)

	// Updates a stored Participant.
	//
	// Returns the updated participant or ParticipantNotFoundError if the participant could not be
	// found.
	Update(name string, update func(pb.Participant) pb.Participant) (*pb.Participant, error)

	// Deletes a Participant from the storage.
	//
	// Returns ParticipantNotFoundError if the participant could not be found.
	Delete(name string) error
}

type ParticipantNotFoundError struct {
	Name string
}

func (e *ParticipantNotFoundError) Error() string {
	return fmt.Sprintf("Could not find %s", e.Name)
}
