package weatherCaller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Khan/genqlient/graphql"
)

func WeatherCaller(city string) map[string]string {
	graphqlClient := graphql.NewClient("https://graphql-weather-api.herokuapp.com", nil)
	getCityResp, err := getCity(context.Background(), graphqlClient, city)
	if err != nil {
		fmt.Printf("An error occurred: %s", err)
	}

	PublicMap := make(map[string]string)
	PublicMap["Title"] = getCityResp.GetCityByName.Weather.Summary.Title
	PublicMap["Description"] = getCityResp.GetCityByName.Weather.Summary.Description
	PublicMap["Actual"] = strconv.FormatFloat(getCityResp.GetCityByName.Weather.Temperature.Actual, 'f', 6, 64)
	PublicMap["FeelsLike"] = strconv.FormatFloat(getCityResp.GetCityByName.Weather.Temperature.FeelsLike, 'f', 6, 64)
	PublicMap["Min"] = strconv.FormatFloat(getCityResp.GetCityByName.Weather.Temperature.Min, 'f', 6, 64)
	PublicMap["Max"] = strconv.FormatFloat(getCityResp.GetCityByName.Weather.Temperature.Max, 'f', 6, 64)

	return PublicMap

}
