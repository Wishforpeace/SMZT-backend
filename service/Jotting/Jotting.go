package service

import (
	"SMZT/model/jottings"
	"SMZT/pkg/errno"
)

func GetJotting(student_id string) ([]jottings.Jotting, error) {
	jottings, err := jottings.GetJottings(student_id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return jottings, nil
}

func PostJotting(student_id string, title string, content string) error {
	if err := jottings.CreateJotting(student_id, title, content); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
