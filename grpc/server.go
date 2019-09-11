package grpc

import (
	"context"

	"google.golang.org/grpc/codes"

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
	println("in two step")
	return nil
}

func (s *DropServer) PutLink(c context.Context, req *proto.PutLinkRequest) (*proto.PutLinkResponse, error) {
	println("in put link")
	code := uint32(codes.OK)
	if err := s.storage.PutObject("public", req.LinkText, []byte(req.LinkAddress)); err != nil {
		code = uint32(codes.Internal)
	}

	resp := &proto.PutLinkResponse{
		Base: &proto.BaseResponse{Code: code},
	}
	return resp, nil
}

func (s *DropServer) RemoveLink(c context.Context, req *proto.RemoveLinkRequest) (*proto.RemoveLinkResponse, error) {
	code := uint32(codes.OK)
	if err := s.storage.DeleteObject("public", req.LinkText); err != nil {
		code = uint32(codes.Internal)
	}

	resp := &proto.RemoveLinkResponse{
		Base: &proto.BaseResponse{Code: code},
	}
	return resp, nil
}

func (s *DropServer) GetLinks(req *proto.GetLinksRequest, stream proto.DropApi_GetLinksServer) error {
	println("in get links")
	links, err := s.storage.GetObjectList("public")
	if err != nil {
		println("err:", err.Error())
		stream.Send(
			&proto.GetLinksResponse{
				Base: &proto.BaseResponse{Code: uint32(codes.Internal)},
			},
		)
		return nil
	}

	println("links: ", len(links))

	for _, linkText := range links {
		println("sending...")
		linkAddress, err := s.storage.GetObjectValue("public", linkText)
		if err != nil {
			println(err.Error())
			continue
		}
		err = stream.Send(
			&proto.GetLinksResponse{
				Base:        &proto.BaseResponse{Code: uint32(codes.OK)},
				LinkText:    linkText,
				LinkAddress: string(linkAddress),
			},
		)
		if err != nil {
			println(err.Error())
		}
		println("sent ", linkText, " -> ", string(linkAddress))
	}

	return nil
}
