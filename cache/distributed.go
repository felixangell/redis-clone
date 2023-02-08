package cache

type Page struct {
	Data []byte
}

type Segment struct {
	Pages []Page
}

type Partition struct {
	// Segments are 'open' for either 1 week of time
	// or until they reach 1GB of keys. After this a
	// new segment is created
	Segments []Segment
}
