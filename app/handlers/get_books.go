package handlers

import (
	"context"
	proto "github.com/fikrirnurhidayat/bookstoresvc/gen"
)

func (b *Backend) GetBooks(ctx context.Context, req *proto.PageRequest) (*proto.BooksResponse, error) {
	return &proto.BooksResponse{
    Meta: &proto.MetaResponse{
      PageSize: 1,
      PageCount: 1,
      Total: 1,
    },
	}, nil
}
