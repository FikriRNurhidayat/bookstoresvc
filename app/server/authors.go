package server

import (
	"context"
	proto "github.com/fikrirnurhidayat/bookstoresvc/proto"
)

func (b *Backend) GetAuthors(ctx context.Context, req *proto.PageRequest) (*proto.AuthorsResponse, error) {
  data, err := b.GetAuthorsUseCase.GetAuthors()
  if err != nil {
    return nil, err
  }

  rsp := &proto.AuthorsResponse{
		Meta: &proto.Meta{
      Pagination: &proto.Pagination{
        Page:     req.Page,
        PageSize: req.PageSize,
      },
		},
	}


	for _, author := range data {
		pack := &proto.Author{
      Id: author.Id,
      Name: author.Name,
		}
		rsp.Data = append(rsp.Data, pack)
	}
	return rsp, nil
}

func (b *Backend) GetAuthor(ctx context.Context, req *proto.GetAuthorRequest) (*proto.AuthorResponse, error) {
  data, err := b.GetAuthorUseCase.GetAuthor(req.GetId())
  if err != nil {
    return nil, err
  }

  rsp := &proto.AuthorResponse{
    Meta: &proto.Meta{},
    Data: &proto.Author{
      Id: data.Id,
      Name: data.Name,
    },
	}

	return rsp, nil
}
