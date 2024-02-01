package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GeoServicer interface {
	GeoSearch(input string) ([]*Address, error)
	GeoCode(lat, lng string) ([]*Address, error)
}

type GeoService struct {
}

func NewGeoService() GeoService {
	return GeoService{}
}

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResponse struct {
	Suggestions []*Suggest `json:"suggestions"`
}

type Suggest struct {
	Value             string   `json:"value"`
	UnrestrictedValue string   `json:"unrestricted_value"`
	Data              *Address `json:"data"`
}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type GeocodeResponse struct {
	Addresses []*Address `json:"addresses"`
}

func main() {

	rejs := SearchRequest{Query: "москва"}
	jsd, _ := json.Marshal(rejs)
	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address", bytes.NewBuffer(jsd))
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 24133541982a4f8baa0497ac37c7661de6598b13")
	req.Header.Set("X-Secret", "bbff5cda452ec7ecbf4eea2f3c4e97538e599b46")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err read body:", err)
	}

	fmt.Println(string(body))
	var sugg SearchResponse
	err = json.Unmarshal(body, &sugg)
	if err != nil {
		log.Fatal("err unmarsh:", err)
	}

	var address []*Address

	for _, s := range sugg.Suggestions {
		address = append(address, s.Data)
	}
	fmt.Println("===================================================")
	for _, m := range address {
		fmt.Println(m.GeoLon)
	}

}

type Address struct {
	PostalCode           interface{} `json:"postal_code"`
	Country              string      `json:"country"`
	CountryISOCode       string      `json:"country_iso_code"`
	FederalDistrict      interface{} `json:"federal_district"`
	RegionFIASID         string      `json:"region_fias_id"`
	RegionKLADRID        string      `json:"region_kladr_id"`
	RegionISOCode        string      `json:"region_iso_code"`
	RegionWithType       string      `json:"region_with_type"`
	RegionType           string      `json:"region_type"`
	RegionTypeFull       string      `json:"region_type_full"`
	Region               string      `json:"region"`
	AreaFIASID           interface{} `json:"area_fias_id"`
	AreaKLADRID          interface{} `json:"area_kladr_id"`
	AreaWithType         interface{} `json:"area_with_type"`
	AreaType             interface{} `json:"area_type"`
	AreaTypeFull         interface{} `json:"area_type_full"`
	Area                 interface{} `json:"area"`
	CityFIASID           string      `json:"city_fias_id"`
	CityKLADRID          string      `json:"city_kladr_id"`
	CityWithType         string      `json:"city_with_type"`
	CityType             string      `json:"city_type"`
	CityTypeFull         string      `json:"city_type_full"`
	City                 string      `json:"city"`
	CityArea             interface{} `json:"city_area"`
	CityDistrictFIASID   interface{} `json:"city_district_fias_id"`
	CityDistrictKLADRID  interface{} `json:"city_district_kladr_id"`
	CityDistrictWithType interface{} `json:"city_district_with_type"`
	CityDistrictType     interface{} `json:"city_district_type"`
	CityDistrictTypeFull interface{} `json:"city_district_type_full"`
	CityDistrict         interface{} `json:"city_district"`
	StreetFIASID         string      `json:"street_fias_id"`
	StreetKLADRID        string      `json:"street_kladr_id"`
	StreetWithType       string      `json:"street_with_type"`
	StreetType           string      `json:"street_type"`
	StreetTypeFull       string      `json:"street_type_full"`
	Street               string      `json:"street"`
	SteadFIASID          interface{} `json:"stead_fias_id"`
	SteadCadnum          interface{} `json:"stead_cadnum"`
	SteadType            interface{} `json:"stead_type"`
	SteadTypeFull        interface{} `json:"stead_type_full"`
	Stead                interface{} `json:"stead"`
	HouseFIASID          interface{} `json:"house_fias_id"`
	HouseKLADRID         interface{} `json:"house_kladr_id"`
	HouseCadnum          interface{} `json:"house_cadnum"`
	HouseType            interface{} `json:"house_type"`
	HouseTypeFull        interface{} `json:"house_type_full"`
	House                interface{} `json:"house"`
	BlockType            interface{} `json:"block_type"`
	BlockTypeFull        interface{} `json:"block_type_full"`
	Block                interface{} `json:"block"`
	Entrance             interface{} `json:"entrance"`
	Floor                interface{} `json:"floor"`
	FlatFIASID           interface{} `json:"flat_fias_id"`
	FlatCadnum           interface{} `json:"flat_cadnum"`
	FlatType             interface{} `json:"flat_type"`
	FlatTypeFull         interface{} `json:"flat_type_full"`
	Flat                 interface{} `json:"flat"`
	FlatArea             interface{} `json:"flat_area"`
	SquareMeterPrice     interface{} `json:"square_meter_price"`
	FlatPrice            interface{} `json:"flat_price"`
	PostalBox            interface{} `json:"postal_box"`
	FIASID               string      `json:"fias_id"`
	FIASCadastreNumber   string      `json:"fias_cadastre_number"`
	FIASLevel            string      `json:"fias_level"`
	FIASActualityState   string      `json:"fias_actuality_state"`
	KLADRID              string      `json:"kladr_id"`
	GeonameID            string      `json:"geoname_id"`
	CapitalMarker        string      `json:"capital_marker"`
	OKATO                string      `json:"okato"`
	OKTMO                string      `json:"oktmo"`
	TaxOffice            string      `json:"tax_office"`
	TaxOfficeLegal       string      `json:"tax_office_legal"`
	Timezone             interface{} `json:"timezone"`
	GeoLat               string      `json:"geo_lat"`
	GeoLon               string      `json:"geo_lon"`
	BeltwayHit           interface{} `json:"beltway_hit"`
	BeltwayDistance      interface{} `json:"beltway_distance"`
	Metro                interface{} `json:"metro"`
	Divisions            interface{} `json:"divisions"`
	QCGeo                string      `json:"qc_geo"`
	QCComplete           interface{} `json:"qc_complete"`
	QCHouse              interface{} `json:"qc_house"`
	HistoryValues        []string    `json:"history_values"`
	UnparsedParts        interface{} `json:"unparsed_parts"`
	Source               interface{} `json:"source"`
	QC                   interface{} `json:"qc"`
}
