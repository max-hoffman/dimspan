package main

import (
	"fmt"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func formatPlotData(data [][]float64) plotter.XYs {
	fmt.Printf("%v\n", len(data))
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = data[i][0]
		pts[i].Y = data[i][1]
		// pts[i].Z = data[i][2]
	}
	return pts
}

func createSVG(data plotter.XYs, title, file, axisOne, axisTwo string) error {
	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("Create new plot error: %v", err)
	}

	p.Title.Text = title
	p.X.Label.Text = axisOne
	p.Y.Label.Text = axisTwo

	err = plotutil.AddLinePoints(p, "Data", data)

	if err != nil {
		return fmt.Errorf("Failed to draw plot: %v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./plots/"+file); err != nil {
		return fmt.Errorf("Failed to save plot: %v", err)
	}

	return nil
}
