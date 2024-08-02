package repository

type Repositories struct {
	Book    IBookRepository
	User    IUserRepository
	History IHistoryRepository
}
