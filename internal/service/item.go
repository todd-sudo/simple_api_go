package service

import (
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	log "github.com/todd-sudo/todo/pkg/logger"
)

type ItemService interface {
	Insert(b dto.ItemCreateDTO) model.Item
	Update(b dto.ItemUpdateDTO) model.Item
	Delete(b model.Item)
	All() []model.Item
	FindByID(itemID uint64) model.Item
	IsAllowedToEdit(userID string, itemID uint64) bool
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) ItemService {
	return &itemService{
		itemRepository: itemRepo,
	}
}

func (service *itemService) Insert(i dto.ItemCreateDTO) model.Item {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	res := service.itemRepository.InsertItem(item)
	return res
}

func (service *itemService) Update(i dto.ItemUpdateDTO) model.Item {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	res := service.itemRepository.UpdateItem(item)
	return res
}

func (service *itemService) Delete(i model.Item) {
	service.itemRepository.DeleteItem(i)
}

func (service *itemService) All() []model.Item {
	return service.itemRepository.AllItem()
}

func (service *itemService) FindByID(itemID uint64) model.Item {
	return service.itemRepository.FindItemByID(itemID)
}

func (service *itemService) IsAllowedToEdit(userID string, itemID uint64) bool {
	i := service.itemRepository.FindItemByID(itemID)
	id := fmt.Sprintf("%v", i.UserID)
	return userID == id
}
