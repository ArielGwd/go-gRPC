package controllers

import (
	"context"
	"database/sql"
	"proyek/models"
	"proyek/pb/cities"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// City struct
type City struct {
	DB *sql.DB

	cities.UnimplementedCitiesServiceServer
}

// fungsi GetCity
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}

// List func
func (s *City) GetCities(in *cities.EmptyMessage, stream cities.CitiesService_GetCitiesServer) error {
	ctx := stream.Context()

	for i := 1; i < 50; i++ {
		err := contextError(ctx)
		if err != nil {
			return err
		}

		res := &cities.CitiesStream{
			City: &cities.City{Id: int32(i), Name: "Jakarta"},
		}

		err = stream.Send(res)
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
		}
	}

	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}

func (s *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, s.DB, in)
	return &cityModel.Pb, err
}

func (s *City) Delete(ctx context.Context, in *cities.Id) (*cities.MyBoolean, error) {
	var cityModel models.City
	err := cityModel.Delete(ctx, s.DB, in)
	if err != nil {
		return &cities.MyBoolean{Boolean: false}, err
	}
	return &cities.MyBoolean{Boolean: true}, nil
}