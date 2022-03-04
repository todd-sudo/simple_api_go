package service

import (
	"context"
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	log "github.com/todd-sudo/todo/pkg/logger"
)

type ItemService interface {
	Insert(ctx context.Context, b dto.ItemCreateDTO) model.Item
	Update(ctx context.Context, b dto.ItemUpdateDTO) model.Item
	Delete(ctx context.Context, b model.Item)
	All(ctx context.Context) []model.Item
	FindByID(ctx context.Context, itemID uint64) model.Item
	IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) bool
}

type itemService struct {
	ctx            context.Context
	itemRepository repository.ItemRepository
}

func NewItemService(ctx context.Context, itemRepo repository.ItemRepository) ItemService {
	return &itemService{
		ctx:            ctx,
		itemRepository: itemRepo,
	}
}

func (service *itemService) Insert(ctx context.Context, i dto.ItemCreateDTO) model.Item {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	res := service.itemRepository.InsertItem(ctx, item)
	return res
}

func (service *itemService) Update(ctx context.Context, i dto.ItemUpdateDTO) model.Item {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	res := service.itemRepository.UpdateItem(ctx, item)
	return res
}

func (service *itemService) Delete(ctx context.Context, i model.Item) {
	service.itemRepository.DeleteItem(ctx, i)
}

func (service *itemService) All(ctx context.Context) []model.Item {
	return service.itemRepository.AllItem(ctx)
}

func (service *itemService) FindByID(ctx context.Context, itemID uint64) model.Item {
	return service.itemRepository.FindItemByID(ctx, itemID)
}

func (service *itemService) IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) bool {
	i := service.itemRepository.FindItemByID(ctx, itemID)
	id := fmt.Sprintf("%v", i.UserID)
	return userID == id
}
