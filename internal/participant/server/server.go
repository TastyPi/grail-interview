package server

import (
	"context"
	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/golang/protobuf/ptypes/empty"
)

type server struct{}

func Create() pb.ParticipantServiceServer {
	return &server{}
}

func (s *server) CreateParticipant(
	ctx context.Context, request *pb.CreateParticipantRequest) (*pb.Participant, error) {
	return &pb.Participant{}, nil
}

func (s *server) GetParticipant(
	ctx context.Context, request *pb.GetParticipantRequest) (*pb.Participant, error) {
	return &pb.Participant{}, nil
}

func (s *server) UpdateParticipant(
	ctx context.Context, request *pb.UpdateParticipantRequest) (*pb.Participant, error) {
	return &pb.Participant{}, nil
}

func (s *server) DeleteParticipant(
	ctx context.Context, request *pb.DeleteParticipantRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
