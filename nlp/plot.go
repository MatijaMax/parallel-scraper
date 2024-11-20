package nlp

import (
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts/v2"
)

func SaveChart(buf []byte) error {
	file := filepath.Join("data", "pie-chart.png")
	err := os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func CreatePieChart(pos float64, neg float64) error {
	values := []float64{
		pos,
		neg,
	}

	p, err := charts.PieRender(
		values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text: "KOMENTARI: Pozitivni i Negativni",
			Left: charts.PositionCenter,
		}),
		charts.PaddingOptionFunc(charts.Box{
			Top:    20,
			Right:  20,
			Bottom: 20,
			Left:   20,
		}),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data: []string{
				"Pozitivni",
				"Negativni",
			},
			Left: charts.PositionLeft,
		}),
		charts.PieSeriesShowLabel(),
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = SaveChart(buf)
	if err != nil {
		panic(err)
	}
	return err
}
