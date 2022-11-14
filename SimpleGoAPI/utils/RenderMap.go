package utils

import (
	"TarlaIO/users"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"io"
	"math/rand"
	"os"
	"strconv"
)

type GeoChartCreator struct{}

//geoBase The function that will be designed as we want. Color,data format etc.
func geoBase(geoData []opts.GeoData) *charts.Geo {
	//Map Title, Map Type and Map Colour
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Location of Users"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "world",
			ItemStyle: &opts.ItemStyle{Color: "#faedcd"},
		}),
	)
	//Features of pointer
	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     8,
			BrushType: "stroke",
		}),
	)
	//Show name of users while location point strokes
	geo.SetSeriesOptions(charts.WithLabelOpts(opts.Label{
		Show:      true,
		Formatter: "{b}",
		Color: "#000000",
		Position: "bottom",
	}))

	return geo
}

// Creates map to show selected user's locations on the map
func (GeoChartCreator) CreateMap(users []users.User, values []int) {

	//Creates struct of map data and array of it
	var locationOfUser opts.GeoData
	var geoData []opts.GeoData

	//Search for all user to find equal parameters with users on id
	for _, user := range users {
		//When client select ex. 3.,6.,9. user on row, it will be get as 'values' in GO
		for _, value := range values {
			if user.Id == value {
				locationOfUser.Name = user.Name
				if lat, err := strconv.ParseFloat(user.Address.Geo.Lat, 64); err == nil {
					if lng, err := strconv.ParseFloat(user.Address.Geo.Lng, 64); err == nil {
						locationOfUser.Value = []float64{lat, lng, float64(rand.Intn(100))}
					}
				}
				geoData = append(geoData, locationOfUser)
			}
		}
	}

	//Create map component
	page := components.NewPage()
	page.AddCharts(
		geoBase(geoData),
	)
	//Create new html file
	f, err := os.Create("view/map.html")
	if err != nil {
		panic(err)
	}
	//Write source code to created file
	page.Render(io.MultiWriter(f))

	//Call this function to add go template keywords to created HTML
	HtmlDesigner("view/map.html", "map")
}
