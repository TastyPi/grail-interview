package server

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mennanov/fmutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/TastyPi/grail-interview/internal/participant/storage"
)

type server struct {
	storage.ParticipantStorage
}

func Create(s storage.ParticipantStorage) pb.ParticipantServiceServer {
	return &server{s}
}

func (s *server) CreateParticipant(
	ctx context.Context, request *pb.CreateParticipantRequest) (*pb.Participant, error) {
	participant := request.Participant
	if participant == nil {
		participant = &pb.Participant{}
	}
	if participant.Name != "" {
		return nil, status.Error(codes.InvalidArgument, "Name should not be set")
	}
	return s.Insert(*participant)
}

func (s *server) GetParticipant(
	ctx context.Context, request *pb.GetParticipantRequest) (*pb.Participant, error) {
	p, err := s.Read(request.Name)
	if err != nil {
		return nil, convertStorageError(err)
	}
	return p, nil
}

func (s *server) UpdateParticipant(
	ctx context.Context, request *pb.UpdateParticipantRequest) (*pb.Participant, error) {
	if !request.UpdateMask.IsValid(request.Participant) {
		return nil, status.Error(codes.InvalidArgument, "update_mask does not match the participant")
	}
	name := request.Participant.Name
	fmutils.Filter(request.Participant, request.UpdateMask.Paths)
	p, err := s.Update(name, func(old pb.Participant) pb.Participant {
		proto.Merge(&old, request.Participant)
		return old
	})
	if err != nil {
		return nil, convertStorageError(err)
	}
	return p, nil
}

func (s *server) DeleteParticipant(
	ctx context.Context, request *pb.DeleteParticipantRequest) (*empty.Empty, error) {
	err := s.Delete(request.Name)
	if err != nil {
		return nil, convertStorageError(err)
	}
	return &empty.Empty{}, nil
}

func convertStorageError(err error) error {
	var notFound *storage.ParticipantNotFoundError
	if errors.As(err, &notFound) {
		return status.Errorf(codes.NotFound, "%s not found", notFound.Name)
	}
	return err
}
