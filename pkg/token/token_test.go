package token

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToken(t *testing.T) {
	var (
		token      string
		id         uint   = 2020
		student_id string = "2021213966"
	)

	Convey("Test token", t, func() {
		Convey("Test token generation", func() {
			var err error
			var tokenPayload = &TokenPayload{
				Id:        id,
				StudentID: student_id,
				Expired:   time.Hour * 2,
			}
			token, err = tokenPayload.GenerateToken()
			So(err, ShouldBeNil)
		})

		Convey("Test token resolution", func() {
			t, err := ResolveToken(token)
			So(err, ShouldBeNil)
			So(t.Id, ShouldEqual, id)
			So(t.StudentID, ShouldEqual, student_id)
		})
	})
}
