package models

import (
	"time"

	"github.com/pkg/errors"
)

type Bucket struct {
	Id    int64
	Title string

	Files []*File `xorm:"-"`

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func BucketCreate(title string) (*Bucket, error) {
	buck := new(Bucket)
	buck.Title = title

	e := Save(buck)

	return buck, e
}

func BucketGet(id int64) (*Bucket, error) {
	var (
		e     error
		buck  = new(Bucket)
		files []*File
	)
	_, e = x.Id(id).Get(buck)
	if e != nil {
		return nil, errors.Wrap(e, "get bucket by id")
	}
	e = x.Where("bucket_id = ?", id).Find(&files)
	if e != nil {
		return nil, errors.Wrap(e, "err get bucket files")
	}
	buck.Files = files
	return buck, e
}

func BucketList(limit int, offsets ...int) (res []*Bucket, err error) {
	err = x.Limit(limit, offsets...).OrderBy("id desc").Find(&res)
	return res, err
}
