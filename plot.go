package main

import (
	"image/color"
	"log"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type offsetUnitGrid struct {
	minVec *Vector
	maxVec *Vector
	N      int
	Data   mat.Matrix
}

func (g offsetUnitGrid) Dims() (c, r int)   { r, c = g.Data.Dims(); return c, r }
func (g offsetUnitGrid) Z(c, r int) float64 { return g.Data.At(r, c) }
func (g offsetUnitGrid) X(c int) float64 {
	_, n := g.Data.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}

	hx := (g.maxVec.x - g.minVec.x) / float64(g.N)
	return g.minVec.x + float64(c)*hx //float64(c) + g.XOffset
}
func (g offsetUnitGrid) Y(r int) float64 {
	m, _ := g.Data.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	hy := (g.maxVec.y - g.minVec.y) / float64(g.N)
	return g.minVec.y + float64(r)*hy //float64(c) + g.XOffset
}

func addPointCharges(p *plot.Plot, Qs []*PointCharge) {
	Qpts := make(plotter.XYs, len(Qs))

	for i := range Qpts {
		Qpts[i].X = Qs[i].rVec.x
		Qpts[i].Y = Qs[i].rVec.y
	}

	qScatter, err := plotter.NewScatter(Qpts)
	if err != nil {
		panic(err)
	}
	qScatter.Radius = 4
	qScatter.Shape = draw.CrossGlyph{}
	qScatter.GlyphStyle.Color = color.RGBA{R: 128, B: 255, A: 255}

	p.Add(qScatter)
}

func calcEletricFieldValues(EFn func(rVec *Vector) *Vector, minVec *Vector, maxVec *Vector, N int) []float64 {
	values := make([]float64, N*N)

	hx := (maxVec.x - minVec.x) / float64(N)
	hy := (maxVec.y - minVec.y) / float64(N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := minVec.x + float64(i)*hx
			y := minVec.y + float64(j)*hy
			point := &Vector{float64(x), float64(y)}
			values[i+N*j] = Magnitude(EFn(point))
		}
	}

	return values
}

func addContourPlot(p *plot.Plot, EFn func(rVec *Vector) *Vector, minVec *Vector, maxVec *Vector, N int) {
	m := offsetUnitGrid{
		N:      N,
		minVec: minVec,
		maxVec: maxVec,
		Data:   mat.NewDense(N, N, calcEletricFieldValues(EFn, minVec, maxVec, N))}
	pal := palette.Heat(1, 1)
	h := plotter.NewContour(m, nil, pal)

	p.Add(h)
}

func convertFn(fn func(vec *Vector) *Vector) func(float64, float64) float64 {
	return func(x float64, y float64) float64 {
		return Magnitude(fn(&Vector{x, y}))
	}
}

func Draw(Qs []*PointCharge, EFn func(rVEc *Vector) *Vector, minVec *Vector, maxVec *Vector, N int) {
	p, err := plot.New()
	if err != nil {
		log.Fatalln("Failed to initialize the plo library.")
	}

	p.Title.Text = "Electric Field E(x, y)"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	addPointCharges(p, Qs)
	addContourPlot(p, EFn, minVec, maxVec, N)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "result.png"); err != nil {
		log.Fatalln("Failed to draw the contour plot.")
	}
}
