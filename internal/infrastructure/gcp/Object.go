package gcp

import "strings"

type Object struct {
	Bucket string
	Key    string
}

func (o *Object) GetPath() string {
	return "gs://" + o.Bucket + "/" + o.Key
}

func ParsePath(path []string) (Object, error) {
	return Object{
		Bucket: path[0],
		Key:    strings.Join(path[1:], "/"),
	}, nil
}
