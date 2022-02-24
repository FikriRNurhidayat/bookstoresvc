package server

import (
  "net/http"
	"context"

	proto "github.com/fikrirnurhidayat/bookstoresvc/proto"
  "google.golang.org/protobuf/types/known/timestamppb"
  "google.golang.org/grpc/status"
)

func (b *Backend) GetBooks(ctx context.Context, req *proto.PageRequest) (*proto.BooksResponse, error) {
  data, err := b.GetBooksUseCase.GetBooks()
  if err != nil {
    return nil, err
  }

  rsp := &proto.BooksResponse{
		Meta: &proto.Meta{
      Pagination: &proto.Pagination{
        Page:     req.Page,
        PageSize: req.PageSize,
      },
		},
	}


	for _, book := range data {
		pack := &proto.Book{
      Id: book.Id,
      Title: book.Title,
      Synopsis: book.Synopsis,
      Isbn: book.Isbn,
      PageCount: book.PageCount,
      CoverType: proto.CoverType(proto.CoverType_value[book.CoverType]),
      CoverImageUrl: book.CoverImageUrl,
      AuthorId: book.AuthorId,
      AuthorName: book.AuthorName,
      PublisherId: book.PublisherId,
      PublisherName: book.PublisherName,
      PublishedAt: timestamppb.New(book.PublishedAt),
		}
		rsp.Data = append(rsp.Data, pack)
	}
	return rsp, nil
}

func (b *Backend) GetBook(ctx context.Context, req *proto.GetBookRequest) (*proto.BookResponse, error) {
  data, err := b.GetBookUseCase.GetBook(req.GetId())
  if err != nil {
    return nil, status.Error(http.StatusNotFound, "Book not found")
  }

  resp := &proto.BookResponse{
    Status: proto.Status_OK,
    Data: &proto.Book{
      Id: data.Id,
      Title: data.Title,
      Synopsis: data.Synopsis,
      Isbn: data.Isbn,
      PageCount: data.PageCount,
      CoverType: proto.CoverType(proto.CoverType_value[data.CoverType]),
      CoverImageUrl: data.CoverImageUrl,
      AuthorId: data.AuthorId,
      AuthorName: data.AuthorName,
      PublisherId: data.PublisherId,
      PublisherName: data.PublisherName,
      PublishedAt: timestamppb.New(data.PublishedAt),
    },
  }

  return resp, nil
}
