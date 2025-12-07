package service

import (
	"errors"
	"testing"

	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/repository"
	"github.com/R4yL-dev/pkmc/internal/repository/mocks"
	"github.com/R4yL-dev/pkmc/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestItemService_CreateItem(t *testing.T) {
	tests := []struct {
		name          string
		extCode       string
		langCode      string
		typeName      string
		price         *float64
		setupMocks    func(*mocks.MockUnitOfWork, *mocks.MockItemRepository, *mocks.MockExtensionRepository, *mocks.MockLanguageRepository, *mocks.MockItemTypeRepository)
		expectedError string
		validateItem  func(*testing.T, *models.Item)
	}{
		{
			name:     "success - create item with price",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "Display",
			price:    testutil.FloatPtr(129.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				// Setup UoW to execute the function
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow) // Execute the function with the mock UoW
				}).Return(nil)

				// Setup repository mocks
				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)
				uow.On("Items").Return(items)

				exts.On("FindByCode", "DRI").Return(&models.Extension{
					Model: gorm.Model{ID: 1},
					Code:  "DRI",
					Name:  "Destinées Radieuses",
				}, nil)

				langs.On("FindByCode", "fr").Return(&models.Language{
					Model: gorm.Model{ID: 1},
					Code:  "fr",
					Name:  "Français",
				}, nil)

				types.On("FindByName", "Display").Return(&models.ItemType{
					Model: gorm.Model{ID: 1},
					Name:  "Display",
				}, nil)

				items.On("Create", mock.MatchedBy(func(item *models.Item) bool {
					return item.ExtensionID == 1 && item.TypeID == 1 && item.LanguageID == 1
				})).Run(func(args mock.Arguments) {
					item := args.Get(0).(*models.Item)
					item.ID = 10 // Simulate DB assigning ID
				}).Return(nil)

				items.On("FindByID", uint(10)).Return(&models.Item{
					Model:       gorm.Model{ID: 10},
					ExtensionID: 1,
					TypeID:      1,
					LanguageID:  1,
					Price:       testutil.FloatPtr(129.99),
					Extension: models.Extension{
						Model: gorm.Model{ID: 1},
						Code:  "DRI",
						Name:  "Destinées Radieuses",
					},
					Type: models.ItemType{
						Model: gorm.Model{ID: 1},
						Name:  "Display",
					},
					Language: models.Language{
						Model: gorm.Model{ID: 1},
						Code:  "fr",
						Name:  "Français",
					},
				}, nil)
			},
			validateItem: func(t *testing.T, item *models.Item) {
				assert.Equal(t, uint(10), item.ID)
				assert.Equal(t, uint(1), item.ExtensionID)
				assert.Equal(t, "DRI", item.Extension.Code)
				assert.Equal(t, "Destinées Radieuses", item.Extension.Name)
				assert.Equal(t, "Display", item.Type.Name)
				assert.Equal(t, "Français", item.Language.Name)
				assert.NotNil(t, item.Price)
				assert.InDelta(t, 129.99, *item.Price, 0.001)
			},
		},
		{
			name:     "success - create item without price",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "Display",
			price:    nil,
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(nil)

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)
				uow.On("Items").Return(items)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}, Code: "DRI"}, nil)
				langs.On("FindByCode", "fr").Return(&models.Language{Model: gorm.Model{ID: 1}, Code: "fr"}, nil)
				types.On("FindByName", "Display").Return(&models.ItemType{Model: gorm.Model{ID: 1}, Name: "Display"}, nil)

				items.On("Create", mock.Anything).Run(func(args mock.Arguments) {
					item := args.Get(0).(*models.Item)
					item.ID = 11
				}).Return(nil)

				items.On("FindByID", uint(11)).Return(&models.Item{
					Model:       gorm.Model{ID: 11},
					ExtensionID: 1,
					TypeID:      1,
					LanguageID:  1,
					Price:       nil,
					Extension:   models.Extension{Model: gorm.Model{ID: 1}},
					Type:        models.ItemType{Model: gorm.Model{ID: 1}},
					Language:    models.Language{Model: gorm.Model{ID: 1}},
				}, nil)
			},
			validateItem: func(t *testing.T, item *models.Item) {
				assert.Equal(t, uint(11), item.ID)
				assert.Nil(t, item.Price)
			},
		},
		{
			name:     "error - extension not found",
			extCode:  "INVALID",
			langCode: "fr",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("extension 'INVALID' not found: record not found"))

				uow.On("Extensions").Return(exts)

				exts.On("FindByCode", "INVALID").Return(nil, errors.New("record not found"))
			},
			expectedError: "extension 'INVALID' not found",
		},
		{
			name:     "error - language not found",
			extCode:  "DRI",
			langCode: "invalid",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("language 'invalid' not found: record not found"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}, Code: "DRI"}, nil)
				langs.On("FindByCode", "invalid").Return(nil, errors.New("record not found"))
			},
			expectedError: "language 'invalid' not found",
		},
		{
			name:     "error - item type not found",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "InvalidType",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("item type 'InvalidType' not found: record not found"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}}, nil)
				langs.On("FindByCode", "fr").Return(&models.Language{Model: gorm.Model{ID: 1}}, nil)
				types.On("FindByName", "InvalidType").Return(nil, errors.New("record not found"))
			},
			expectedError: "item type 'InvalidType' not found",
		},
		{
			name:     "error - failed to create item",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("failed to create item: database error"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)
				uow.On("Items").Return(items)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}}, nil)
				langs.On("FindByCode", "fr").Return(&models.Language{Model: gorm.Model{ID: 1}}, nil)
				types.On("FindByName", "Display").Return(&models.ItemType{Model: gorm.Model{ID: 1}}, nil)
				items.On("Create", mock.Anything).Return(errors.New("database error"))
			},
			expectedError: "failed to create item",
		},
		{
			name:     "error - failed to load created item",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("failed to load created item: database error"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)
				uow.On("Items").Return(items)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}}, nil)
				langs.On("FindByCode", "fr").Return(&models.Language{Model: gorm.Model{ID: 1}}, nil)
				types.On("FindByName", "Display").Return(&models.ItemType{Model: gorm.Model{ID: 1}}, nil)
				items.On("Create", mock.Anything).Run(func(args mock.Arguments) {
					item := args.Get(0).(*models.Item)
					item.ID = 12
				}).Return(nil)
				items.On("FindByID", uint(12)).Return(nil, errors.New("database error"))
			},
			expectedError: "failed to load created item",
		},
		{
			name:     "edge case - empty extension code",
			extCode:  "",
			langCode: "fr",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("extension '' not found: record not found"))

				uow.On("Extensions").Return(exts)

				exts.On("FindByCode", "").Return(nil, errors.New("record not found"))
			},
			expectedError: "extension '' not found",
		},
		{
			name:     "edge case - empty language code",
			extCode:  "DRI",
			langCode: "",
			typeName: "Display",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("language '' not found: record not found"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}}, nil)
				langs.On("FindByCode", "").Return(nil, errors.New("record not found"))
			},
			expectedError: "language '' not found",
		},
		{
			name:     "edge case - empty type name",
			extCode:  "DRI",
			langCode: "fr",
			typeName: "",
			price:    testutil.FloatPtr(99.99),
			setupMocks: func(uow *mocks.MockUnitOfWork, items *mocks.MockItemRepository, exts *mocks.MockExtensionRepository, langs *mocks.MockLanguageRepository, types *mocks.MockItemTypeRepository) {
				uow.On("Do", mock.AnythingOfType("func(repository.UnitOfWork) error")).Run(func(args mock.Arguments) {
					fn := args.Get(0).(func(repository.UnitOfWork) error)
					fn(uow)
				}).Return(errors.New("item type '' not found: record not found"))

				uow.On("Extensions").Return(exts)
				uow.On("Languages").Return(langs)
				uow.On("ItemTypes").Return(types)

				exts.On("FindByCode", "DRI").Return(&models.Extension{Model: gorm.Model{ID: 1}}, nil)
				langs.On("FindByCode", "fr").Return(&models.Language{Model: gorm.Model{ID: 1}}, nil)
				types.On("FindByName", "").Return(nil, errors.New("record not found"))
			},
			expectedError: "item type '' not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockUoW := mocks.NewMockUnitOfWork(t)
			mockItems := mocks.NewMockItemRepository(t)
			mockExts := mocks.NewMockExtensionRepository(t)
			mockLangs := mocks.NewMockLanguageRepository(t)
			mockTypes := mocks.NewMockItemTypeRepository(t)

			tt.setupMocks(mockUoW, mockItems, mockExts, mockLangs, mockTypes)

			// Create service
			service := NewItemService(mockUoW)

			// Execute
			item, err := service.CreateItem(tt.extCode, tt.langCode, tt.typeName, tt.price)

			// Assert
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, item)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, item)
				if tt.validateItem != nil {
					tt.validateItem(t, item)
				}
			}
		})
	}
}
