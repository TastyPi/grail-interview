package storage_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/TastyPi/grail-interview/internal/participant/storage"
)

func TestInsert_ParticipantHasName(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()

	p, err := s.Insert(pb.Participant{Name: "participant/AAA-001"})

	assert.Nil(t, p)
	assert.NotNil(t, err)
}

func TestInsert_SetsName(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	original := pb.Participant{
		GivenName:   "Foo",
		FamilyName:  "Bar",
		BirthDate:   &date.Date{Year: 1970, Month: 1, Day: 1},
		PhoneNumber: "01189998819991197253",
		Address:     "221B Baker Street",
	}

	p, err := s.Insert(original)

	assert.Nil(t, err)
	assert.NotEmpty(t, p.Name)
	original.Name = p.Name
	assert.Empty(t, cmp.Diff(original, p, protocmp.Transform()))
}

func TestRead_ParticipantNotFoundError(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	name := "participant/does-not-exist"
	// Add a participant to ensure it is not returned.
	s.Insert(pb.Participant{})

	p, err := s.Read(name)

	assert.Nil(t, p)
	assert.Equal(t, &storage.ParticipantNotFoundError{name}, err)
}

func TestUpdate_ParticipantNotFoundError(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	name := "participant/does-not-exist"
	// Add a participant to ensure it is not used in the update.
	s.Insert(pb.Participant{})

	p, err := s.Update(name, func(p pb.Participant) pb.Participant {
		t.Error("This function should not be called")
		return p
	})

	assert.Nil(t, p)
	assert.Equal(t, &storage.ParticipantNotFoundError{name}, err)
}

func TestUpdate_ChangingNameReturnsError(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	inserted, _ := s.Insert(pb.Participant{})

	p, err := s.Update(inserted.Name, func(p pb.Participant) pb.Participant {
		p.Name = "participant/another-name"
		return p
	})

	assert.Nil(t, p)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	original := &pb.Participant{
		GivenName:   "Foo",
		FamilyName:  "Bar",
		BirthDate:   &date.Date{Year: 1970, Month: 1, Day: 1},
		PhoneNumber: "01189998819991197253",
		Address:     "221B Baker Street",
	}
	updatedFields := &pb.Participant{
		GivenName:   "Baz",
		FamilyName:  "Qux",
		BirthDate:   &date.Date{Year: 2000, Month: 12, Day: 31},
		PhoneNumber: "0987654321",
		Address:     "10 Downing Street",
	}
	original, _ = s.Insert(*original)

	p, err := s.Update(original.Name, func(p pb.Participant) pb.Participant {
		proto.Merge(&p, updatedFields)
		return p
	})

	assert.Nil(t, err)

	// Sanity check that original was not changed
	assert.Equal(t, "Foo", original.GivenName)

	expected := &(*original)
	proto.Merge(expected, updatedFields)
	assert.Empty(t, cmp.Diff(expected, p, protocmp.Transform()))
}

func TestUpdateThenRead(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	original := &pb.Participant{
		GivenName:   "Foo",
		FamilyName:  "Bar",
		BirthDate:   &date.Date{Year: 1970, Month: 1, Day: 1},
		PhoneNumber: "01189998819991197253",
		Address:     "221B Baker Street",
	}
	updatedFields := &pb.Participant{
		GivenName:   "Baz",
		FamilyName:  "Qux",
		BirthDate:   &date.Date{Year: 2000, Month: 12, Day: 31},
		PhoneNumber: "0987654321",
		Address:     "10 Downing Street",
	}
	original, _ = s.Insert(*original)

	expected, _ := s.Update(original.Name, func(p pb.Participant) pb.Participant {
		proto.Merge(&p, updatedFields)
		return p
	})

	actual, _ := s.Read(original.Name)

	assert.Empty(t, cmp.Diff(expected, actual, protocmp.Transform()))
}

func TestDelete_ParticipantNotFoundError(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	name := "participant/does-not-exist"
	// Add a participant to ensure it is not deleted instead.
	other, _ := s.Insert(pb.Participant{})

	err := s.Delete(name)

	assert.Equal(t, &storage.ParticipantNotFoundError{name}, err)
	// Check the other participant still exists
	p, _ := s.Read(other.Name)
	assert.NotNil(t, p)
}

func TestDelete(t *testing.T) {
	s := storage.CreateInMemoryParticipantStorage()
	p, _ := s.Insert(pb.Participant{})

	err := s.Delete(p.Name)

	assert.Nil(t, err)
	// Check that the participant can no longer be read.
	deleted, _ := s.Read(p.Name)
	assert.Nil(t, deleted)
}
