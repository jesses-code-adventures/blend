package ingest

type Ingestor interface {
	// Set the locations and contents for the struct.
	Ingest()
	// A map of fully qualified location strings for the files.
	Locations() map[string]struct{}
	// A map where the keys are identical to Ingestor.Locations and the values are the contents of the files.
	Contents() map[string]string
}
