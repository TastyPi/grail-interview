// In-memory implementation of ParticipantStorage
// In a real app this would be implemented using a database, but the assignment
// said persistance was not required so we'll just use a map for simplicity.
package storage

import (
	"errors"
	"fmt"

	pb "github.com/TastyPi/grail-interview/api/participant"
)

type storage map[string]pb.Participant

var lastReference = 0

func createName() string {
	// In real code this should do something more sensible than ordered IDs, but the assignment says
	// to assume some mechanism is provided, so I'm pretending this is that mechanism.
	lastReference += 1
	return fmt.Sprintf("participants/AAA-%03d", lastReference)
}

func CreateInMemoryParticipantStorage() ParticipantStorage {
	return make(storage)
}

func (s storage) Insert(participant pb.Participant) (*pb.Participant, error) {
	if participant.Name != "" {
		return nil, errors.New("Participant must not have a Name when inserting")
	}
	name := createName()
	participant.Name = name
	s[name] = participant
	return &participant, nil
}

func (s storage) Read(name string) (*pb.Participant, error) {
	p, found := s[name]
	if !found {
		return nil, &ParticipantNotFoundError{name}
	}
	return &p, nil
}

func (s storage) Update(
	name string, update func(pb.Participant) pb.Participant) (*pb.Participant, error) {
	old, found := s[name]
	if !found {
		return nil, &ParticipantNotFoundError{name}
	}
	new := update(old)
	if new.Name != name {
		return nil, errors.New("Participant's Name must not be changed")
	}
	s[name] = new
	return &new, nil
}

func (s storage) Delete(name string) error {
	if _, found := s[name]; !found {
		return &ParticipantNotFoundError{name}
	}
	delete(s, name)
	return nil
}
