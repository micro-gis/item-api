package services

import (
	geoalert "github.com/micro-gis/geo-alert-api/domain/geoalert"
	"github.com/micro-gis/geo-alert-api/domain/queries"
	"github.com/micro-gis/utils/rest_errors"
)

var (
	GeoAlertService geoalertServiceInterface = &geoalertService{}
)

type geoalertServiceInterface interface {
	Create(geoalert.GeoAlert) (*geoalert.GeoAlert, rest_errors.RestErr)
	Get(string) (*geoalert.GeoAlert, rest_errors.RestErr)
	Search(queries.EsQuery, bool) ([]geoalert.GeoAlert, rest_errors.RestErr)
	GetUserGeoAlerts(int64, bool) ([]geoalert.GeoAlert, rest_errors.RestErr)
	Delete(id string) (*geoalert.GeoAlert, rest_errors.RestErr)
	Upsert(geoalert.GeoAlert, string) (*geoalert.GeoAlert, rest_errors.RestErr)
}

type geoalertService struct{}

func (is *geoalertService) Create(geo geoalert.GeoAlert) (*geoalert.GeoAlert, rest_errors.RestErr) {
	if err := geo.Save(); err != nil {
		return nil, err
	}
	return &geo, nil
}

func (is *geoalertService) Get(id string) (*geoalert.GeoAlert, rest_errors.RestErr) {
	geo := geoalert.GeoAlert{Id: id}
	if err := geo.Get(); err != nil {
		return nil, err
	}
	return &geo, nil
}

func (is *geoalertService) Search(query queries.EsQuery, isPublic bool) ([]geoalert.GeoAlert, rest_errors.RestErr) {
	dao := geoalert.GeoAlert{}
	if isPublic {
		queryPublic := queries.FieldValue{
			Field: "scope",
			Value: "Public",
		}
		query.Equals = append(query.Equals, queryPublic)
	}
	return dao.Search(query)
}

func (is *geoalertService) GetUserGeoAlerts(id int64, isPublic bool) ([]geoalert.GeoAlert, rest_errors.RestErr) {
	var query = queries.EsQuery{Equals: []queries.FieldValue{{
		Field: "user_id",
		Value: id,
	}}}
	if isPublic {
		queryPublic := queries.FieldValue{
			Field: "scope",
			Value: "Public",
		}
		query.Equals = append(query.Equals, queryPublic)
	}
	dao := geoalert.GeoAlert{}
	return dao.Search(query)
}
func (s *geoalertService) Delete(id string) (*geoalert.GeoAlert, rest_errors.RestErr) {
	geoa := geoalert.GeoAlert{Id: id}
	if err := geoa.Delete(); err != nil {
		return nil, err
	}
	return &geoa, nil
}

func (s *geoalertService) Upsert(geoa geoalert.GeoAlert, geoaId string) (*geoalert.GeoAlert, rest_errors.RestErr) {
	if err := geoa.Upsert(geoaId); err != nil {
		return nil, err
	}
	return &geoa, nil
}
