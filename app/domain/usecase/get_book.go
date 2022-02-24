package usecase

import (
  "google.golang.org/grpc/grpclog"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/entity"
  "github.com/fikrirnurhidayat/bookstoresvc/app/domain/repository"
)

type GetBookUseCase interface {
  GetBook(id uint32) (entity.Book, error)
}

type GetBookUseCaseImpl struct {
  logger grpclog.LoggerV2;
  bookRepository repository.IBookRepository;
}

func NewGetBookUseCase(logger grpclog.LoggerV2, bookRepository repository.IBookRepository) *GetBookUseCaseImpl {
  return &GetBookUseCaseImpl{
    logger: logger,
    bookRepository: bookRepository,
  }
}

func (s GetBookUseCaseImpl) GetBook(id uint32) (entity.Book, error) {
  return s.bookRepository.GetBook(id)
}
