package mlscratchlib

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

// PlotVectors does shit
func PlotVectors(x []float64, y []float64, title string) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i := range x {
		err = plotutil.AddLinePoints(p, makePoint(x[i], y[i]))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "points.png"); err != nil {
		log.Fatal(err)
	}
}

func makePoint(x, y float64) plotter.XYs {
	point := make(plotter.XYs, 1)
	point[0].X = x
	point[0].Y = y
	return point
}
