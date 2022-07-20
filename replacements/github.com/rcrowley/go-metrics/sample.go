package metrics

type Sample interface {
	Clear()
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Size() int
	Snapshot() Sample
	StdDev() float64
	Sum() int64
	Update(int64)
	Values() []int64
	Variance() float64
}

func NewExpDecaySample(int, float64) Sample {
	return NilSample{}
}

// NilSample is a no-op Sample.
type NilSample struct{}

// Clear is a no-op.
func (NilSample) Clear() {}

// Count is a no-op.
func (NilSample) Count() int64 { return 0 }

// Max is a no-op.
func (NilSample) Max() int64 { return 0 }

// Mean is a no-op.
func (NilSample) Mean() float64 { return 0.0 }

// Min is a no-op.
func (NilSample) Min() int64 { return 0 }

// Percentile is a no-op.
func (NilSample) Percentile(p float64) float64 { return 0.0 }

// Percentiles is a no-op.
func (NilSample) Percentiles(ps []float64) []float64 {
	return make([]float64, len(ps))
}

// Size is a no-op.
func (NilSample) Size() int { return 0 }

// Sample is a no-op.
func (NilSample) Snapshot() Sample { return NilSample{} }

// StdDev is a no-op.
func (NilSample) StdDev() float64 { return 0.0 }

// Sum is a no-op.
func (NilSample) Sum() int64 { return 0 }

// Update is a no-op.
func (NilSample) Update(v int64) {}

// Values is a no-op.
func (NilSample) Values() []int64 { return []int64{} }

// Variance is a no-op.
func (NilSample) Variance() float64 { return 0.0 }

func NewUniformSample(reservoirSize int) Sample {
	return NilSample{}
}
