package service

import (
	"context"
	"grpc-go-learning/internal/database"
	"grpc-go-learning/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {

	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {

	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesList []*pb.Category
	for _, c := range categories {
		categoriesList = append(categoriesList, &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categoriesList,
	}, nil
}

func (c *CategoryService) FindCategory(ctx context.Context, in *pb.FindCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindCategory(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
