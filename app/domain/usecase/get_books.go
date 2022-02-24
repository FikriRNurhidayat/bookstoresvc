package usecase

import (
  "google.golang.org/grpc/grpclog"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/repository"
)

type GetBooksUseCase interface {
  GetBooks() ([]entity.Book, error)
}

type GetBooksUseCaseImpl struct {
  logger grpclog.LoggerV2;
  bookRepository repository.IBookRepository;
}

func NewGetBooksUseCase(logger grpclog.LoggerV2, bookRepository repository.IBookRepository) *GetBooksUseCaseImpl {
  return &GetBooksUseCaseImpl{
    logger: logger,
    bookRepository: bookRepository,
  }
}

func (s GetBooksUseCaseImpl) GetBooks() ([]entity.Book, error) {
  return s.bookRepository.GetBooks()
}
