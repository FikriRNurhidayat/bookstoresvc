package usecase

import (
  "google.golang.org/grpc/grpclog"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/repository"
)

type GetAuthorUseCase interface {
  GetAuthor(id uint32) (entity.Author, error)
}

type GetAuthorUseCaseImpl struct {
  logger grpclog.LoggerV2;
  authorRepository repository.IAuthorRepository;
}

func NewGetAuthorUseCase(logger grpclog.LoggerV2, authorRepository repository.IAuthorRepository) *GetAuthorUseCaseImpl {
  return &GetAuthorUseCaseImpl{
    logger: logger,
    authorRepository: authorRepository,
  }
}

func (s GetAuthorUseCaseImpl) GetAuthor(id uint32) (entity.Author, error) {
  return s.authorRepository.GetAuthor(id)
}
