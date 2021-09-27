package member

import (
	"context"
	"database/sql"
	"time"

	"github.com/fauzanmh/olp-user/constant"
	"github.com/fauzanmh/olp-user/entity"
	"github.com/fauzanmh/olp-user/entity/microservice"
	appInit "github.com/fauzanmh/olp-user/init"
	authAdapter "github.com/fauzanmh/olp-user/repository/adapter/auth"
	mysqlRepo "github.com/fauzanmh/olp-user/repository/mysql"
	"github.com/fauzanmh/olp-user/schema/member"
)

type usecase struct {
	config      *appInit.Config
	mysqlRepo   mysqlRepo.Repository
	authAdapter authAdapter.AuthAdapter
}

func NewMemberUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository, authAdapter authAdapter.AuthAdapter) Usecase {
	return &usecase{
		config:      config,
		mysqlRepo:   mysqlRepo,
		authAdapter: authAdapter,
	}
}

// --- register --- ///
func (u *usecase) Register(ctx context.Context, req *member.RegisterRequest) (err error) {
	// begin transaction
	tx, err := u.mysqlRepo.BeginTx(ctx)
	if err != nil {
		return
	}

	// check if email exists
	exist, err := u.mysqlRepo.WithTx(tx).CheckEmail(ctx, req.Email)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}
	if exist {
		u.mysqlRepo.RollbackTx(tx)
		err = constant.ErrorMessageUniqueEmail
		return
	}

	// store register to database
	registerParams := &entity.RegisterParams{
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}
	id, err := u.mysqlRepo.WithTx(tx).Register(ctx, registerParams)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}

	// store register to ms auth
	createUserRequest := &microservice.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		MemberID: id,
	}
	err = u.authAdapter.CreateUser(ctx, createUserRequest)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}

	// commit transaction
	u.mysqlRepo.CommitTx(tx)

	return
}

// --- delete member --- //
func (u *usecase) DeleteMember(ctx context.Context, req *member.DeleteMemberRequest) (err error) {
	// begin transaction
	tx, err := u.mysqlRepo.BeginTx(ctx)
	if err != nil {
		return
	}

	// check if member exists
	member, err := u.mysqlRepo.WithTx(tx).CheckMember(ctx, req.ID)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}
	if member.DeletedAt.Valid {
		u.mysqlRepo.RollbackTx(tx)
		err = constant.ErrorMessageUserHasBeenDeleted
		return
	}

	// delete user from database
	deleteMemberParams := &entity.DeleteMemberParams{
		ID:        req.ID,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}
	err = u.mysqlRepo.WithTx(tx).DeleteMember(ctx, deleteMemberParams)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}

	err = u.authAdapter.DeleteUser(ctx, req.ID)
	if err != nil {
		u.mysqlRepo.RollbackTx(tx)
		return
	}

	// commit transaction
	u.mysqlRepo.CommitTx(tx)

	return
}
