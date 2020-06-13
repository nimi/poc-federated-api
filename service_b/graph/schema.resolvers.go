package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/example/federation/reviews/graph/model"
)

// These two methods are required for gqlgen to resolve the internal id-only wrapper structs.
// This boilerplate might be removed in a future version of gqlgen that can no-op id only nodes.
func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	return &model.Product{
		Upc: upc,
	}, nil
}

func (r *entityResolver) FindByUserID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID: id,
	}, nil
}

// Here we implement the stitched part of this service, returning reviews for a product. Of course normally you would
// go back to the database, but we are just making some data up here.
func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	switch obj.Upc {
	case "top-1":
		return []*model.Review{{
			Body: "A highly effective form of birth control.",
		}}, nil

	case "top-2":
		return []*model.Review{{
			Body: "Fedoras are one of the most fashionable hats around and can look great with a variety of outfits.",
		}}, nil

	case "top-3":
		return []*model.Review{{
			Body: "This is the last straw. Hat you will wear. 11/10",
		}}, nil

	}
	return nil, nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	if obj.ID == "1234" {
		return []*model.Review{{
			Body: "Has an odd fascination with hats.",
		}}, nil
	}
	return nil, nil
}
