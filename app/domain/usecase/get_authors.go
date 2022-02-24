package usecase

import (
  "google.golang.org/grpc/grpclog"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/repository"
)

type GetAuthorsUseCase interface {
  GetAuthors() ([]entity.Author, error)
}

type GetAuthorsUseCaseImpl struct {
  logger grpclog.LoggerV2;
  authorRepository repository.IAuthorRepository;
}

func NewGetAuthorsUseCase(logger grpclog.LoggerV2, authorRepository repository.IAuthorRepository) *GetAuthorsUseCaseImpl {
  return &GetAuthorsUseCaseImpl{
    logger: logger,
    authorRepository: authorRepository,
  }
}

func (s GetAuthorsUseCaseImpl) GetAuthors() ([]entity.Author, error) {
  return s.authorRepository.GetAuthors()
}
