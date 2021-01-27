package server_test

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"

	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/TastyPi/grail-interview/internal/participant/server"
	"github.com/TastyPi/grail-interview/internal/participant/storage"
)

func TestCreateParticipant_NameSet(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)

	response, err := server.CreateParticipant(context.Background(),
		&pb.CreateParticipantRequest{Participant: &pb.Participant{
			Name: "participant/AAA-001",
		}})

	assert.Nil(t, response)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestCreateParticipant(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	p := &pb.Participant{
		GivenName:   "Foo",
		FamilyName:  "Bar",
		BirthDate:   &date.Date{Year: 1970, Month: 1, Day: 1},
		PhoneNumber: "01189998819991197253",
		Address:     "221B Baker Street",
	}

	response, err := server.CreateParticipant(context.Background(),
		&pb.CreateParticipantRequest{Participant: p})

	assert.Nil(t, err)
	assert.Empty(t, cmp.Diff(response, p,
		protocmp.Transform(), protocmp.IgnoreFields(response, "name")))
}

func TestGetParticipant_NotFound(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	// Create a participant to ensure it is not returned.
	server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.GetParticipant(context.Background(), &pb.GetParticipantRequest{
		Name: "participant/does-not-exist",
	})

	assert.Nil(t, response)
	assert.Equal(t, codes.NotFound, status.Code(err))
}

func TestGetParticipant(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	// Create a participant to ensure it is not returned.
	p, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.GetParticipant(context.Background(), &pb.GetParticipantRequest{
		Name: p.Name,
	})

	assert.Nil(t, err)
	assert.Empty(t, cmp.Diff(p, response, protocmp.Transform()))
}

func TestUpdateParticipant_NotFound(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	// Create a participant to ensure it is not updated.
	other, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.UpdateParticipant(context.Background(), &pb.UpdateParticipantRequest{
		Participant: &pb.Participant{
			Name: "participant/does-not-exist",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{}},
	})

	assert.Nil(t, response)
	assert.Equal(t, codes.NotFound, status.Code(err), err)
	// Check that the other participant did not change in storage
	storedOther, _ := store.Read(other.Name)
	assert.Empty(t, cmp.Diff(other, storedOther, protocmp.Transform()))
}

func TestUpdateParticipant_UpdateMaskMissingField(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	p, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.UpdateParticipant(context.Background(), &pb.UpdateParticipantRequest{
		Participant: &pb.Participant{
			Name:       p.Name,
			GivenName:  "Foo",
			FamilyName: "Bar",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"given_name"}},
	})

	assert.Nil(t, err)
	// Only the given_name has been updated.
	p.GivenName = "Foo"
	assert.Empty(t, cmp.Diff(p, response, protocmp.Transform()))
}

func TestGetAfterUpdate(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	p, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	updateResp, _ := server.UpdateParticipant(context.Background(), &pb.UpdateParticipantRequest{
		Participant: &pb.Participant{
			Name:      p.Name,
			GivenName: "Foo",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"given_name"}},
	})
	getResp, _ := server.GetParticipant(context.Background(), &pb.GetParticipantRequest{Name: p.Name})

	assert.Empty(t, cmp.Diff(updateResp, getResp, protocmp.Transform()))
}

func TestDelete_NotFound(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	// Create a participant to ensure it is not deleted.
	other, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.DeleteParticipant(context.Background(), &pb.DeleteParticipantRequest{
		Name: "participant/does-not-exist",
	})

	assert.Nil(t, response)
	assert.Equal(t, codes.NotFound, status.Code(err), err)
	// Check that the other participant still exists
	storedOther, _ := store.Read(other.Name)
	assert.Empty(t, cmp.Diff(other, storedOther, protocmp.Transform()))

}

func TestDelete(t *testing.T) {
	store := storage.CreateInMemoryParticipantStorage()
	server := server.Create(store)
	p, _ := server.CreateParticipant(context.Background(), &pb.CreateParticipantRequest{})

	response, err := server.DeleteParticipant(context.Background(), &pb.DeleteParticipantRequest{
		Name: p.Name,
	})

	assert.Nil(t, err)
	assert.Empty(t, cmp.Diff(&empty.Empty{}, response, protocmp.Transform()))
	// Check that it is no longer in the store
	stored, _ := store.Read(p.Name)
	assert.Nil(t, stored)
}
