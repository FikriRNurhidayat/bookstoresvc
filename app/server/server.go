package server

import (
  "os"
  "io/ioutil"

  "github.com/jmoiron/sqlx"
	"google.golang.org/grpc/grpclog"
	"github.com/fikrirnurhidayat/bookstoresvc/app/config"
	"github.com/fikrirnurhidayat/bookstoresvc/app/driver/database"
	"github.com/fikrirnurhidayat/bookstoresvc/app/domain/usecase"
	"github.com/fikrirnurhidayat/bookstoresvc/app/domain/repository"

	proto "github.com/fikrirnurhidayat/bookstoresvc/proto"
)

type Backend struct {
  proto.BookstoreServiceServer
  Logger grpclog.LoggerV2
  db     *sqlx.DB;

  // Repository Object
  BookRepository repository.IBookRepository
  AuthorRepository repository.IAuthorRepository
  
  // Use Case Object
  GetBookUseCase usecase.GetBookUseCase;
  GetBooksUseCase usecase.GetBooksUseCase;
  GetAuthorsUseCase usecase.GetAuthorsUseCase;
  GetAuthorUseCase usecase.GetAuthorUseCase;
}

func NewServer() *Backend {
  // Setup Logger
  log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

  // Connect to database
  dataSourceName := config.GetDatabaseConnectionString()
  db, err := database.ConnectDB(dataSourceName) 
  if err != nil {
    log.Fatalf("Cannot connect to the database: %v", err)
  }

  log.Infof("Connected to database: %s", dataSourceName)

  // Create server
	server := &Backend{
		Logger: log,
    db: db,
	}

  // Create repository
  server.BookRepository = repository.NewBookRepository(db)
  server.AuthorRepository = repository.NewAuthorRepository(db)

  // Create use case
  server.GetBookUseCase = usecase.NewGetBookUseCase(server.Logger, server.BookRepository)
  server.GetBooksUseCase = usecase.NewGetBooksUseCase(server.Logger, server.BookRepository)
  server.GetAuthorsUseCase = usecase.NewGetAuthorsUseCase(server.Logger, server.AuthorRepository)
  server.GetAuthorUseCase = usecase.NewGetAuthorUseCase(server.Logger, server.AuthorRepository)

  return server;
} 
