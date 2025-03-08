package service

import (
	"errors"

	"github.com/lemonoa/LemonOA-Go/model"

	"gorm.io/gorm"
)

// BasicContractService 基础数据-合同模块服务
type BasicContractService struct {
	db *gorm.DB
}

func NewBasicContractService(db *gorm.DB) *BasicContractService {
	return &BasicContractService{db: db}
}

// GetContractCategoryList 获取合同分类列表
func (s *BasicContractService) GetContractCategoryList() ([]model.ContractCategory, error) {
	var categories []model.ContractCategory
	err := s.db.Order("sort asc").Find(&categories).Error
	return categories, err
}

// GetContractCategoryByID 根据ID获取合同分类
func (s *BasicContractService) GetContractCategoryByID(id uint) (*model.ContractCategory, error) {
	var category model.ContractCategory
	err := s.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateContractCategory 创建合同分类
func (s *BasicContractService) CreateContractCategory(category *model.ContractCategory) error {
	return s.db.Create(category).Error
}

// UpdateContractCategory 更新合同分类
func (s *BasicContractService) UpdateContractCategory(category *model.ContractCategory) error {
	if category.ID == 0 {
		return errors.New("contract category id is required")
	}
	return s.db.Model(category).Updates(category).Error
}

// DeleteContractCategory 删除合同分类
func (s *BasicContractService) DeleteContractCategory(id uint) error {
	return s.db.Delete(&model.ContractCategory{}, id).Error
}

// GetProductCategoryList 获取产品分类列表
func (s *BasicContractService) GetProductCategoryList(parentID *uint) ([]model.ProductCategory, error) {
	var categories []model.ProductCategory
	query := s.db.Model(&model.ProductCategory{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&categories).Error
	return categories, err
}

// GetProductCategoryByID 根据ID获取产品分类
func (s *BasicContractService) GetProductCategoryByID(id uint) (*model.ProductCategory, error) {
	var category model.ProductCategory
	err := s.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateProductCategory 创建产品分类
func (s *BasicContractService) CreateProductCategory(category *model.ProductCategory) error {
	return s.db.Create(category).Error
}

// UpdateProductCategory 更新产品分类
func (s *BasicContractService) UpdateProductCategory(category *model.ProductCategory) error {
	if category.ID == 0 {
		return errors.New("product category id is required")
	}
	return s.db.Model(category).Updates(category).Error
}

// DeleteProductCategory 删除产品分类
func (s *BasicContractService) DeleteProductCategory(id uint) error {
	// 检查是否有子分类
	var count int64
	if err := s.db.Model(&model.ProductCategory{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with sub-categories")
	}

	// 检查是否有关联的产品
	if err := s.db.Model(&model.Product{}).Where("category_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with associated products")
	}

	return s.db.Delete(&model.ProductCategory{}, id).Error
}

// GetProductList 获取产品列表
func (s *BasicContractService) GetProductList(categoryID uint) ([]model.Product, error) {
	var products []model.Product
	query := s.db.Model(&model.Product{})
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	err := query.Find(&products).Error
	return products, err
}

// GetProductByID 根据ID获取产品
func (s *BasicContractService) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := s.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// CreateProduct 创建产品
func (s *BasicContractService) CreateProduct(product *model.Product) error {
	return s.db.Create(product).Error
}

// UpdateProduct 更新产品
func (s *BasicContractService) UpdateProduct(product *model.Product) error {
	if product.ID == 0 {
		return errors.New("product id is required")
	}
	return s.db.Model(product).Updates(product).Error
}

// DeleteProduct 删除产品
func (s *BasicContractService) DeleteProduct(id uint) error {
	return s.db.Delete(&model.Product{}, id).Error
}

// GetServiceContentList 获取服务内容列表
func (s *BasicContractService) GetServiceContentList() ([]model.ServiceContent, error) {
	var contents []model.ServiceContent
	err := s.db.Order("sort asc").Find(&contents).Error
	return contents, err
}

// GetServiceContentByID 根据ID获取服务内容
func (s *BasicContractService) GetServiceContentByID(id uint) (*model.ServiceContent, error) {
	var content model.ServiceContent
	err := s.db.First(&content, id).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// CreateServiceContent 创建服务内容
func (s *BasicContractService) CreateServiceContent(content *model.ServiceContent) error {
	return s.db.Create(content).Error
}

// UpdateServiceContent 更新服务内容
func (s *BasicContractService) UpdateServiceContent(content *model.ServiceContent) error {
	if content.ID == 0 {
		return errors.New("service content id is required")
	}
	return s.db.Model(content).Updates(content).Error
}

// DeleteServiceContent 删除服务内容
func (s *BasicContractService) DeleteServiceContent(id uint) error {
	return s.db.Delete(&model.ServiceContent{}, id).Error
}

// GetSupplierList 获取供应商列表
func (s *BasicContractService) GetSupplierList() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := s.db.Find(&suppliers).Error
	return suppliers, err
}

// GetSupplierByID 根据ID获取供应商
func (s *BasicContractService) GetSupplierByID(id uint) (*model.Supplier, error) {
	var supplier model.Supplier
	err := s.db.First(&supplier, id).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

// CreateSupplier 创建供应商
func (s *BasicContractService) CreateSupplier(supplier *model.Supplier) error {
	return s.db.Create(supplier).Error
}

// UpdateSupplier 更新供应商
func (s *BasicContractService) UpdateSupplier(supplier *model.Supplier) error {
	if supplier.ID == 0 {
		return errors.New("supplier id is required")
	}
	return s.db.Model(supplier).Updates(supplier).Error
}

// DeleteSupplier 删除供应商
func (s *BasicContractService) DeleteSupplier(id uint) error {
	return s.db.Delete(&model.Supplier{}, id).Error
}

// GetPurchaseCategoryList 获取采购品分类列表
func (s *BasicContractService) GetPurchaseCategoryList(parentID *uint) ([]model.PurchaseCategory, error) {
	var categories []model.PurchaseCategory
	query := s.db.Model(&model.PurchaseCategory{})
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}
	err := query.Order("sort asc").Find(&categories).Error
	return categories, err
}

// GetPurchaseCategoryByID 根据ID获取采购品分类
func (s *BasicContractService) GetPurchaseCategoryByID(id uint) (*model.PurchaseCategory, error) {
	var category model.PurchaseCategory
	err := s.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreatePurchaseCategory 创建采购品分类
func (s *BasicContractService) CreatePurchaseCategory(category *model.PurchaseCategory) error {
	return s.db.Create(category).Error
}

// UpdatePurchaseCategory 更新采购品分类
func (s *BasicContractService) UpdatePurchaseCategory(category *model.PurchaseCategory) error {
	if category.ID == 0 {
		return errors.New("purchase category id is required")
	}
	return s.db.Model(category).Updates(category).Error
}

// DeletePurchaseCategory 删除采购品分类
func (s *BasicContractService) DeletePurchaseCategory(id uint) error {
	// 检查是否有子分类
	var count int64
	if err := s.db.Model(&model.PurchaseCategory{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with sub-categories")
	}

	// 检查是否有关联的采购品
	if err := s.db.Model(&model.PurchaseItem{}).Where("category_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with associated purchase items")
	}

	return s.db.Delete(&model.PurchaseCategory{}, id).Error
}

// GetPurchaseItemList 获取采购品列表
func (s *BasicContractService) GetPurchaseItemList(categoryID uint) ([]model.PurchaseItem, error) {
	var items []model.PurchaseItem
	query := s.db.Model(&model.PurchaseItem{})
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	err := query.Find(&items).Error
	return items, err
}

// GetPurchaseItemByID 根据ID获取采购品
func (s *BasicContractService) GetPurchaseItemByID(id uint) (*model.PurchaseItem, error) {
	var item model.PurchaseItem
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// CreatePurchaseItem 创建采购品
func (s *BasicContractService) CreatePurchaseItem(item *model.PurchaseItem) error {
	return s.db.Create(item).Error
}

// UpdatePurchaseItem 更新采购品
func (s *BasicContractService) UpdatePurchaseItem(item *model.PurchaseItem) error {
	if item.ID == 0 {
		return errors.New("purchase item id is required")
	}
	return s.db.Model(item).Updates(item).Error
}

// DeletePurchaseItem 删除采购品
func (s *BasicContractService) DeletePurchaseItem(id uint) error {
	return s.db.Delete(&model.PurchaseItem{}, id).Error
}
