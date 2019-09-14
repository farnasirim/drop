package grpc

import (
	"context"

	"github.com/farnasirim/drop"
	"github.com/farnasirim/drop/proto"
)

type DropServer struct {
	storage drop.StorageService
}

func NewDropServer(s drop.StorageService) *DropServer {
	return &DropServer{
		storage: s,
	}
}

func (s *DropServer) TwoStepLogin(*proto.LoginRequest, proto.DropApi_TwoStepLoginServer) error {
	return nil
}

func (s *DropServer) PutLink(c context.Context, req *proto.PutLinkRequest) (*proto.PutLinkResponse, error) {
	record := recordFromProto(req.GetLink())
	id, err := s.storage.PutRecord("public", record)
	if err != nil {
		id = -1
	}

	record.IDField = id

	resp := &proto.PutLinkResponse{
		Link: protoFromRecord(record),
	}
	return resp, nil
}

func (s *DropServer) RemoveLink(c context.Context, req *proto.RemoveLinkRequest) (*proto.RemoveLinkResponse, error) {
	rec := recordFromProto(req.GetLink())
	if err := s.storage.DeleteRecord("public", req.GetLink().GetId()); err != nil {
		rec.IDField = -1
	}

	resp := &proto.RemoveLinkResponse{
		Link: protoFromRecord(rec),
	}
	return resp, nil
}

func (s *DropServer) Subscribe(req *proto.SubscribeRequest, stream proto.DropApi_SubscribeServer) error {
	var lastRec int64 = 0

	if req.GetExcludePast() {
		lastRec = req.GetLink().GetId()
	}

	base, lastId := s.storage.AllRecordsAfter(stream.Context(), "public", lastRec)
	creates := s.storage.AllCreateEventsAfter(stream.Context(), "public", lastId)
	deletes := s.storage.AllDeleteEvents(stream.Context(), "public")

	done := false
	for !done {
		select {
		case <-stream.Context().Done():
			done = true

		case record := <-base:
			stream.Send(
				&proto.SubscribeResponse{
					Action: proto.SubscribeResponse_CREATE,
					Record: protoFromRecord(record),
				},
			)
		case record := <-creates:
			stream.Send(
				&proto.SubscribeResponse{
					Action: proto.SubscribeResponse_CREATE,
					Record: protoFromRecord(record),
				},
			)
		case record := <-deletes:
			stream.Send(
				&proto.SubscribeResponse{
					Action: proto.SubscribeResponse_REMOVE,
					Record: protoFromRecord(record),
				},
			)
		}
	}
	return nil
}
