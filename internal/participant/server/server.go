package server

import (
	"context"
	"errors"
	"regexp"

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
	if err := validateParticipant(participant); err != nil {
		return nil, err
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
	if err := validateParticipant(request.Participant); err != nil {
		return nil, err
	}

	name := request.Participant.Name
	// Filter out fields that are not part of the update_mask
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

var phoneNumberRegexp = regexp.MustCompile(`^\+?[0-9\s]+$`)

func validateParticipant(p *pb.Participant) error {
	if p.PhoneNumber != "" && !phoneNumberRegexp.MatchString(p.PhoneNumber) {
		return status.Errorf(codes.InvalidArgument, "phone_number %s is invalid", p.PhoneNumber)
	}
	return nil
}

func convertStorageError(err error) error {
	var notFound *storage.ParticipantNotFoundError
	if errors.As(err, &notFound) {
		return status.Errorf(codes.NotFound, "%s not found", notFound.Name)
	}
	return err
}
