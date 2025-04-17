package handler

import (
	"context"
	"goproject/global"
	"goproject/model"
	"strconv"
	"time"
	"webthree/goproject/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"gorm.io/gorm"
)

type UserServiceServer struct {
}

func ModelToResponse(user *model.User) *proto.UserInfoResponse {
	userInfoRsp := &proto.UserInfoResponse{
		Id:       int32(user.ID),
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
		Birthday: user.Birthday.Format("2006-01-02"),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func (s *UserServiceServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//here is the implementation of GetUserList function
	var users []model.User
	res := global.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(res.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Page), int(req.PageSize))).Find(&users)
	for _, user := range users {
		userinfoRsp := ModelToResponse(&user)
		rsp.Data = append(rsp.Data, userinfoRsp)
	}
	return rsp, nil
}
func (s *UserServiceServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return ModelToResponse(&user), nil
}
func (s *UserServiceServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	//here is the implementation of GetUserById function
	var user model.User
	result := global.DB.Where(&model.User{Id: req.Id}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return ModelToResponse(&user), nil
}
func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//here is the implementation of CreateUser function
	var user model.User
	// 检查输入参数是否有效
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "手机号已存在")
	}
	// 创建用户模型
	user := model.User{
		Mobile:   req.Mobile,
		NickName: req.NickName,
	}
	//密码加密
	options := &password.Options{16, 100, 50, sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	NewPasswoed := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	user.Password = NewPasswoed
	// 插入用户到数据库
	result := global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	// 返回用户信息
	return ModelToResponse(&user), nil
}
func (s *UserServiceServer) UpdateUser(ctx context.Context,req *protoUpdateUserInfo) (*proto.emptypb.Empty, error) {
	//here is the implementation of UpdateUser function
	// 更新用户模型
	user := model.User{
		ID:       req.Id,
		Mobile:   req.Mobile,
		NickName: req.NickName,
		Gender:   req.Gender,
		Role:     req.Role,
		Birthday: time.Unix(int64(req.Birthday), 0),
	}
	// 更新用户到数据库
	result := global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	// 返回空响应
	return &emptypb.Empty{}, nil
}
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *proto.IdRequest) (*proto.emptypb.Empty, error) {
	//here is the implementation of DeleteUser function
	// 检查输入参数是否有效
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "缺少必要的用户ID")
	}
	// 删除用户
	result := global.DB.Delete(&model.User{}, "id = ?", req.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	// 返回空响应
	return &emptypb.Empty{}, nil
}

func (s *UserServiceServer) CheckPassword(ctx context.Context,req *proto.EncyptedPassword) (*proto.CheckPasswordResponse, error){
	//here is the implementation of CheckPassword function
	// 检查输入参数是否有效
	if req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "缺少必要的密码")
    }
	// 检查密码是否正确
	options := &password.Options{16, 100, 50, sha512.New}
	passwordInfo := strings.Split(req.EncyptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
    return &proto.CheckPasswordResponse{success: check}, nil
}