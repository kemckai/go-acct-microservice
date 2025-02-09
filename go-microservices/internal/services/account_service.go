package services

import (
    "errors"
    "sync"
    "go-microservices/internal/models"
)

type AccountService struct {
    mu       sync.Mutex
    accounts map[int]*models.Account
    nextID   int}

func NewAccountService() *AccountService {
    return &AccountService{
        accounts: make(map[int]*models.Account),
        nextID:   1,
    }
}

func (s *AccountService) CreateAccount(name string, initialBalance float64) *models.Account {
    s.mu.Lock()
    defer s.mu.Unlock()

    account := &models.Account{
        ID:      s.nextID,
        Name:    name,
        Balance: initialBalance,
    }
    s.accounts[s.nextID] = account
    s.nextID++
    return account
}

func (s *AccountService) GetAccount(id int) (*models.Account, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    account, exists := s.accounts[id]
    if (!exists) {
        return nil, errors.New("account not found")
    }
    return account, nil
}

func (s *AccountService) UpdateAccount(id int, name string, balance float64) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    account, exists := s.accounts[id]
    if (!exists) {
        return errors.New("account not found")
    }
    account.Name = name
    account.Balance = balance
    return nil
}

func (s *AccountService) DeleteAccount(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.accounts[id]; !exists {
        return errors.New("account not found")
    }
    delete(s.accounts, id)
    return nil
}