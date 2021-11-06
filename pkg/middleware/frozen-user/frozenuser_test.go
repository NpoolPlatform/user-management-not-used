package frozenuser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/user-management/message/npool"
	userinfo "github.com/NpoolPlatform/user-management/pkg/crud/user-info"
	testinit "github.com/NpoolPlatform/user-management/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestFrozenUserMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	resp, err := userinfo.Create(context.Background(), &npool.AddUserRequest{
		UserInfo: &npool.UserBasicInfo{
			Username:    "test-frozen" + uuid.New().String(),
			Password:    "123456789",
			PhoneNumber: uuid.New().String(),
		},
	})
	assert.Nil(t, err)

	frozenUserInfo := npool.FrozenUser{
		UserId:      resp.UserInfo.UserId,
		FrozenBy:    uuid.New().String(),
		FrozenCause: "user has some illegal operations",
	}

	resp1, err := FrozenUser(context.Background(), &npool.FrozenUserRequest{
		UserId:      frozenUserInfo.UserId,
		FrozenBy:    frozenUserInfo.FrozenBy,
		FrozenCause: frozenUserInfo.FrozenCause,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.FrozenUserInfo.Id, uuid.UUID{})
		assert.Equal(t, resp1.FrozenUserInfo.UserId, frozenUserInfo.UserId)
		assert.Equal(t, resp1.FrozenUserInfo.FrozenBy, frozenUserInfo.FrozenBy)
		assert.Equal(t, resp1.FrozenUserInfo.FrozenCause, frozenUserInfo.FrozenCause)
		frozenUserInfo.Id = resp1.FrozenUserInfo.Id
	}

	frozenUserInfo.UnfrozenBy = uuid.New().String()
	resp2, err := UnfrozenUser(context.Background(), &npool.UnfrozenUserRequest{
		Id:         frozenUserInfo.Id,
		UserId:     frozenUserInfo.UserId,
		UnfrozenBy: frozenUserInfo.UnfrozenBy,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.UnFrozenUserInfo.Id, frozenUserInfo.Id)
		assert.Equal(t, resp2.UnFrozenUserInfo.UserId, frozenUserInfo.UserId)
		assert.Equal(t, resp2.UnFrozenUserInfo.FrozenBy, frozenUserInfo.FrozenBy)
		assert.Equal(t, resp2.UnFrozenUserInfo.FrozenCause, frozenUserInfo.FrozenCause)
		assert.Equal(t, resp2.UnFrozenUserInfo.UnfrozenBy, frozenUserInfo.UnfrozenBy)
	}
}
