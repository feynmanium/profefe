package profile_test

import (
	"context"
	"testing"

	"github.com/profefe/profefe/pkg/logger"
	"github.com/profefe/profefe/pkg/profile"
	"github.com/profefe/profefe/pkg/storage/inmemory"
	"github.com/stretchr/testify/require"
)

func TestRepository_CreateProfile(t *testing.T) {
	log := logger.NewNop()
	st := inmemory.New()
	repo := profile.NewRepository(log, st)

	ctx := context.Background()

	createReq := &profile.CreateProfileRequest{
		ID:      "123abc",
		Service: "test",
		Labels:  profile.Labels{{"key", "value"}},
	}
	token, err := repo.CreateProfile(ctx, createReq)
	require.NoError(t, err)
	require.NotEmptyf(t, token, "CreateProfile: empty token")

	queryReq := &profile.QueryRequest{
		Service: "test",
	}
	profs, err := st.Query(ctx, queryReq)
	require.NoError(t, err)
	require.Len(t, profs, 1)
	require.Equal(t, token, profs[0].Service.Token)
}
