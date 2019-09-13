package grpc

import (
	"github.com/farnasirim/drop"
	"github.com/farnasirim/drop/proto"
)

func protoFromRecord(r drop.Record) *proto.Record {
	return &proto.Record{
		LinkText:    r.Text(),
		LinkAddress: r.Address(),
		Id:          r.ID(),
	}
}

func recordFromProto(r *proto.Record) *Record {
	return &Record{
		TextField:    r.LinkText,
		AddressField: r.LinkAddress,
		IDField:      r.Id,
	}
}
