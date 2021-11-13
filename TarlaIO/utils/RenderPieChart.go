package utils

import (
	"TarlaIO/users"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
)

type PieChartGenerator struct{}

//pieBase The function that will be designed as we want. Color,data format etc.
func pieBase(pieDatas []opts.PieData) *charts.Pie {

	//Setting title of Pie Chart
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Post Count of Users"}),
	)
	//Show each user's ,that is selected, name and value of posts
	pie.AddSeries("pie", pieDatas).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)
	return pie
}

// Creates map to show selected user's locations on the map
func (PieChartGenerator) CreatePieChart(users []users.User, values []int) {

	//Creates struct of pie data and array of it
	var pieDatas []opts.PieData
	var pieData opts.PieData

	//Search for all user to find equal parameters with users on id
	for _, i := range users {
		//When client select ex. 3.,6.,9. user on row, it will be get as 'values' in GO
		for _, j := range values {
			if i.Id == j {
				pieData.Name = i.Name
				pieData.Value = i.PostNumber
				pieDatas = append(pieDatas, pieData)
			}
		}
	}

	// If there is no selected parameters show default valued chart
	if len(pieDatas) == 0 {
		pieDatas = append(pieDatas, opts.PieData{
			Name:  "Please select an user. Default Value is ",
			Value: 0,
		})
	}

	//Create pie chart components
	page := components.NewPage()
	page.AddCharts(
		pieBase(pieDatas),
	)
	//Create new html file
	f, err := os.Create("view/pie.html")
	if err != nil {
		panic(err)
	}
	//Write source code to created file
	page.Render(io.MultiWriter(f))
	//Call this function to add go template keywords to created HTML
	HtmlDesigner("view/pie.html", "pie")

}
