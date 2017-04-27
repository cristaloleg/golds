package filter

// Filter interface that all filter-like collections implement
type Filter interface {
	Add(value interface{})
	Has(value interface{}) bool
	Count() int
}
