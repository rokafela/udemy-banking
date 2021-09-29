package service

import (
	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	customersDto := make([]dto.CustomerResponse, 0)
	for _, v := range c {
		customersDto = append(customersDto, v.ToDto())
	}
	return customersDto, nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
