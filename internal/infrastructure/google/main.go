package google

type ObjectHandle interface {
	GetObject(bucket string, key string)
}
