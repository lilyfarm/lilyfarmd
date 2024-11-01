package geography

// Point is any point in a metric space
type Point interface {
	Dimensions() int
	Value(dimension int) (float64, error)
}

// MetricSpace defines a metric space in the Real Analysis sense [1].
// We have a set from which our co-ordinates are derived
// as well as a distance function. To simplify things, let
// us assume we are working in \mathbb{R}^2 and the "metric"
// returns values in \mathbb{R}.
//
// [1]: "Introduction to Analysis" by Maxwell Rosenlicht
// [2]: https://en.wikipedia.org/wiki/Metric_space
type MetricSpace interface {
	Distance(a Point, b Point) (float64, error)
}
