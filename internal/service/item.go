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
	Insert(ctx context.Context, b dto.ItemCreateDTO) (*model.Item, error)
	Update(ctx context.Context, b dto.ItemUpdateDTO) (*model.Item, error)
	Delete(ctx context.Context, b model.Item) error
	All(ctx context.Context) ([]*model.Item, error)
	FindByID(ctx context.Context, itemID uint64) (*model.Item, error)
	IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) (bool, error)
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

func (service *itemService) Insert(ctx context.Context, i dto.ItemCreateDTO) (*model.Item, error) {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	itemM, errI := service.itemRepository.InsertItem(ctx, item)
	if errI != nil {
		log.Errorf("item insert error: %v", errI)
		return nil, err
	}
	return itemM, nil
}

func (service *itemService) Update(ctx context.Context, i dto.ItemUpdateDTO) (*model.Item, error) {
	item := model.Item{}
	err := smapping.FillStruct(&item, smapping.MapFields(&i))
	if err != nil {
		log.Errorf("Failed map %v: ", err)
	}
	itemM, errI := service.itemRepository.UpdateItem(ctx, item)
	if errI != nil {
		log.Errorf("item update error: %v", errI)
		return nil, errI
	}
	return itemM, nil
}

func (service *itemService) Delete(ctx context.Context, i model.Item) error {
	err := service.itemRepository.DeleteItem(ctx, i)
	if err != nil {
		log.Errorf("item delete error: %v", err)
		return err
	}
	return nil
}

func (service *itemService) All(ctx context.Context) ([]*model.Item, error) {
	items, err := service.itemRepository.AllItem(ctx)
	if err != nil {
		log.Errorf("get all items error: %v", err)
		return nil, err
	}
	return items, nil
}

func (service *itemService) FindByID(ctx context.Context, itemID uint64) (*model.Item, error) {
	item, err := service.itemRepository.FindItemByID(ctx, itemID)
	if err != nil {
		log.Errorf("find item by id error: %v", err)
		return nil, err
	}
	return item, nil
}

func (service *itemService) IsAllowedToEdit(ctx context.Context, userID string, itemID uint64) (bool, error) {
	item, err := service.itemRepository.FindItemByID(ctx, itemID)
	if err != nil {
		log.Errorf("is allowed to edit item error: %v", err)
		return false, err
	}
	id := fmt.Sprintf("%v", item.UserID)
	return userID == id, nil
}
