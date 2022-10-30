package service

import (
	"SMZT/model"
	"SMZT/model/user"
	"SMZT/pkg/errno"
	Token "SMZT/pkg/token"
	"SMZT/util"
	"fmt"
)

type Id struct {
	id string
}

//
// import (
// 	"SMZT/model/user"
// 	"SMZT/pkg/auth"
// 	"SMZT/pkg/constvar"
// 	"SMZT/pkg/token"
// 	// pb "SMZT/proto"
// 	"SMZT/util"
// )

// Login ... 登录
func Login(student_id string, pwd string) (string, error) {
	// 根据 studentId 在 DB 查询 user
	//user, err := userModel.GetUserByStudentId(studentId)
	var Id int
	var userInfo user.UserModel
	userInfo.StudentId = student_id
	userInfo.HashPassword = pwd
	//	if err := model.DB.Self.Where("student_id = ?", student_id).First(&userInfo); err.Error != nil {
	//		fmt.Println(err, err.Error)
	//		return "", errno.ErrUserNotExisted
	//	}

	//仿照team
	//首次登录，验证一站式
	//判断是否首次登录
	result := model.DB.Self.Where("student_id = ?", student_id).First(&Id)
	// result := model.DB.Self.Where("student_id = ?", student_id).First(&userInfo)
	if result.Error != nil {
		I, err := model.GetUserInfoFormOne(userInfo.StudentId, pwd)
		fmt.Println(I)
		if err != nil {
			return "", errno.ErrUserNotExisted
		}
		//对用户信息初始化
		//对密码进行base64加密
		//u.Password = base64.StdEncoding.EncodeToString([]byte(u.Password))
		userInfo.HashPassword = model.GeneratePasswordHash(userInfo.HashPassword)
		model.DB.Self.Create(&userInfo)
		//		model.DB.Self.Where("student_id = ?", userInfo.StudentId).Select("id").Find(&userInfo.ID)
	} else {
		//在数据库中解密比较
		//	password, _ := base64.StdEncoding.DecodeString(u.Password)
		fmt.Println(userInfo.HashPassword)
		fmt.Println(pwd)
		if model.CheckPassword(pwd, userInfo.HashPassword) == false {
			fmt.Printf("1")
			return "", errno.ErrPasswordIncorrect
		}
		model.DB.Self.Where("student_id = ?", userInfo.StudentId).Select("id").Find(&userInfo.ID)
		// if string(password) != pwd {
		// 	c.JSON(http.StatusUnauthorized, "password or account is wrong.")
		// 	return
		// }
	}
	fmt.Println(userInfo.ID)
	//team部分

	//	md5 := Md5.New()
	//	md5.Write([]byte(pwd))
	//	hashPwd := hex.EncodeToString(md5.Sum(nil))

	// 生成 auth token
	var payload = Token.TokenPayload{
		Id:        userInfo.ID,
		StudentID: userInfo.StudentId,
		Expired:   util.GetExpiredTime(),
	}
	token, err := payload.GenerateToken()
	if err != nil {
		fmt.Printf("3")
		return "", err
	}

	return token, nil

}
