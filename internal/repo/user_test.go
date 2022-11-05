package repo_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/aria3ppp/watch-server/internal/models"
	"github.com/aria3ppp/watch-server/internal/repo"
	"github.com/aria3ppp/watch-server/internal/testutils"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestUserGet(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	user := &models.User{
		Email:          "username@example.com",
		HashedPassword: "jfdjsfks",
	}

	// no user

	fetchedUser, err := r.UserGet(ctx, user.ID)
	require.Equal(repo.ErrNoRecord, err)
	require.Nil(fetchedUser)

	// create user

	err = r.UserCreate(ctx, user)

	require.NoError(err)
	require.NotEqual(user.ID, 0)

	// fetch the user

	fetchedUser, err = r.UserGet(ctx, user.ID)
	require.NoError(err)
	require.Equal(user, fetchedUser)
}

func TestUserGetByEmail(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	user := &models.User{
		Email:          "username@example.com",
		HashedPassword: "jfdjsfks",
	}

	// no user

	fetchedUser, err := r.UserGetByEmail(ctx, user.Email)
	require.Equal(repo.ErrNoRecord, err)
	require.Nil(fetchedUser)

	// create user

	err = r.UserCreate(ctx, user)

	require.NoError(err)
	require.NotEqual(user.ID, 0)

	// fetch user

	fetchedUser, err = r.UserGetByEmail(ctx, user.Email)
	require.NoError(err)
	require.Equal(user, fetchedUser)
}

func TestUsersCount(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	// no user

	count, err := r.UsersCount(ctx)
	require.NoError(err)
	require.Equal(0, count)

	// insert users

	users := []*models.User{
		{
			Email:          "fdfjshfksjd@fdkjjd.com",
			HashedPassword: "jslffjlsfs",
		},
		{
			Email:          "afyuehf@dfhk.com",
			HashedPassword: "dsfjsljsd",
		},
		{
			Email:          "fhskgksj@kfdfj.com",
			HashedPassword: "fdsgfgafg",
		},
		{
			Email:          "igrhhgurhgur@wewef.com",
			HashedPassword: "fdsljfafjlkg",
		},
		{
			Email:          "qwdadasd@dkeffjejf.com",
			HashedPassword: "fjsldkgjksag",
		},
	}

	for _, user := range users {
		err := r.UserCreate(ctx, user)
		require.NoError(err)
		require.NotEqual(user.ID, 0)
	}

	// count

	count, err = r.UsersCount(ctx)
	require.NoError(err)
	require.Equal(len(users), count)
}

func TestUserCreate(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	user := &models.User{
		Email:          "username@example.com",
		HashedPassword: "jfdjsfks",
	}

	err = r.UserCreate(ctx, user)

	require.NoError(err)
	require.NotEqual(user.ID, 0)

	fetchedUser, err := r.UserGet(ctx, user.ID)
	require.NoError(err)
	require.Equal(user, fetchedUser)
}

func TestUserDelete(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	user1 := &models.User{
		Email:          "username1@example.com",
		HashedPassword: "jfdjsfks",
	}
	user2 := &models.User{
		Email:          "username2@example.com",
		HashedPassword: "lkflhlhkf",
	}

	// create user

	err = r.UserCreate(ctx, user1)

	require.NoError(err)
	require.NotEqual(user1.ID, 0)

	err = r.UserCreate(ctx, user2)

	require.NoError(err)
	require.NotEqual(user2.ID, 0)

	err = r.UserDelete(ctx, user1.ID)
	require.NoError(err)

	// no user

	fetchedUser1, err := r.UserGet(ctx, user1.ID)
	require.Equal(repo.ErrNoRecord, err)
	require.Nil(fetchedUser1)

	// check user

	fetchedUser2, err := r.UserGet(ctx, user2.ID)
	require.NoError(err)
	require.Equal(user2, fetchedUser2)
}

func TestUserUpdate(t *testing.T) {
	require := require.New(t)

	teardown, err := setup()
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	r := repo.NewRepository(db)
	ctx := context.Background()

	user1 := &models.User{
		Email:          "username1@example.com",
		HashedPassword: "jfdjsfks",
	}
	user2 := &models.User{
		Email:          "username2@example.com",
		HashedPassword: "lkflhlhkf",
	}

	// no user

	err = r.UserUpdate(
		ctx,
		1,
		models.M{models.UserColumns.ID: 99},
	)
	require.Equal(repo.ErrNoRecord, err)

	// create user

	err = r.UserCreate(ctx, user1)

	require.NoError(err)
	require.NotEqual(user1.ID, 0)

	err = r.UserCreate(ctx, user2)

	require.NoError(err)
	require.NotEqual(user2.ID, 0)

	// update user

	user1.Email = "changed_email"
	user1.HashedPassword = "changed_hashed_password"

	cols := models.M{
		models.UserColumns.Email:          user1.Email,
		models.UserColumns.HashedPassword: user1.HashedPassword,
	}

	err = r.UserUpdate(ctx, user1.ID, cols)
	require.NoError(err)

	// check user

	fetchedUser1, err := r.UserGet(ctx, user1.ID)
	require.NoError(err)
	require.Equal(user1, fetchedUser1)

	fetchedUser2, err := r.UserGet(ctx, user2.ID)
	require.NoError(err)
	require.Equal(user2, fetchedUser2)
}

func TestUserUpdate_UpdateFields(t *testing.T) {
	user := models.User{
		Email:          "email",
		HashedPassword: "hashed_password",
		FirstName:      null.StringFrom("first_name"),
		LastName:       null.StringFrom("last_name"),
		Bio:            null.StringFrom("bio"),
		Birthdate: null.TimeFrom(
			testutils.Date(1990, 1, 1),
		),
		Joindate: testutils.Date(2020, 8, 16),
	}

	updatedUser := models.User{
		ID:             1001,
		Email:          "updated_email",
		HashedPassword: "updated_hashed_password",
		FirstName:      null.StringFrom("updated_first_name"),
		LastName:       null.StringFrom("updated_last_name"),
		Bio:            null.StringFrom("updated_bio"),
		Birthdate: null.TimeFrom(
			testutils.Date(1997, 1, 1),
		),
		Joindate: testutils.Date(2022, 8, 16),
	}

	testCases := []struct {
		name           string
		user           models.User
		updatedColumns models.User
	}{
		{
			name:           "tc0",
			user:           user,
			updatedColumns: models.User{},
		},
		{
			name: "tc1",
			user: user,
			updatedColumns: models.User{
				ID: updatedUser.ID,
			},
		},
		{
			name: "tc2",
			user: user,
			updatedColumns: models.User{
				ID:    updatedUser.ID,
				Email: updatedUser.Email,
			},
		},
		{
			name: "tc3",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
			},
		},
		{
			name: "tc4",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
			},
		},
		{
			name: "tc5",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
			},
		},
		{
			name: "tc6",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
				Bio:            updatedUser.Bio,
			},
		},
		{
			name: "tc7",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
				Bio:            updatedUser.Bio,
				Birthdate:      updatedUser.Birthdate,
			},
		},
		{
			name: "tc8",
			user: user,
			updatedColumns: models.User{
				ID:             updatedUser.ID,
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
				Bio:            updatedUser.Bio,
				Birthdate:      updatedUser.Birthdate,
				Joindate:       updatedUser.Joindate,
			},
		},
		{
			name: "tc9",
			user: user,
			updatedColumns: models.User{
				Email:          updatedUser.Email,
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
				Bio:            updatedUser.Bio,
				Birthdate:      updatedUser.Birthdate,
				Joindate:       updatedUser.Joindate,
			},
		},
		{
			name: "tc10",
			user: user,
			updatedColumns: models.User{
				HashedPassword: updatedUser.HashedPassword,
				FirstName:      updatedUser.FirstName,
				LastName:       updatedUser.LastName,
				Bio:            updatedUser.Bio,
				Birthdate:      updatedUser.Birthdate,
				Joindate:       updatedUser.Joindate,
			},
		},
		{
			name: "tc11",
			user: user,
			updatedColumns: models.User{
				FirstName: updatedUser.FirstName,
				LastName:  updatedUser.LastName,
				Bio:       updatedUser.Bio,
				Birthdate: updatedUser.Birthdate,
				Joindate:  updatedUser.Joindate,
			},
		},
		{
			name: "tc12",
			user: user,
			updatedColumns: models.User{
				LastName:  updatedUser.LastName,
				Bio:       updatedUser.Bio,
				Birthdate: updatedUser.Birthdate,
				Joindate:  updatedUser.Joindate,
			},
		},
		{
			name: "tc13",
			user: user,
			updatedColumns: models.User{
				Bio:       updatedUser.Bio,
				Birthdate: updatedUser.Birthdate,
				Joindate:  updatedUser.Joindate,
			},
		},
		{
			name: "tc14",
			user: user,
			updatedColumns: models.User{
				Birthdate: updatedUser.Birthdate,
				Joindate:  updatedUser.Joindate,
			},
		},
		{
			name: "tc15",
			user: user,
			updatedColumns: models.User{
				Joindate: updatedUser.Joindate,
			},
		},
	}

	for i, tc := range testCases {
		user := tc.user
		updatedColumns := tc.updatedColumns

		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)

			teardown, err := setup()
			require.NoError(err)
			t.Cleanup(func() { teardown() })

			r := repo.NewRepository(db)
			ctx := context.Background()

			// make unique emails
			user.Email += strconv.Itoa(i)

			err = r.UserCreate(ctx, &user)
			require.NoError(err)
			require.NotEqual(0, user.ID)

			cols := make(map[string]any)
			if updatedColumns.ID != 0 {
				cols[models.UserColumns.ID] = updatedColumns.ID + i
			}
			if updatedColumns.Email != "" {
				// make unique emails
				updatedColumns.Email += strconv.Itoa(i)
				cols[models.UserColumns.Email] = updatedColumns.Email
			}
			if updatedColumns.HashedPassword != "" {
				cols[models.UserColumns.HashedPassword] = updatedColumns.HashedPassword
			}
			if updatedColumns.FirstName.Valid {
				cols[models.UserColumns.FirstName] = updatedColumns.FirstName.String
			}
			if updatedColumns.LastName.Valid {
				cols[models.UserColumns.LastName] = updatedColumns.LastName.String
			}
			if updatedColumns.Bio.Valid {
				cols[models.UserColumns.Bio] = updatedColumns.Bio.String
			}
			if updatedColumns.Birthdate.Valid {
				cols[models.UserColumns.Birthdate] = updatedColumns.Birthdate.Time
			}
			if updatedColumns.Joindate != (time.Time{}) {
				cols[models.UserColumns.Joindate] = updatedColumns.Joindate
			}

			err = r.UserUpdate(ctx, user.ID, cols)
			require.NoError(err)

			userID := user.ID
			if updatedColumns.ID != 0 {
				userID = updatedColumns.ID + i
			}
			fetchedUser, err := r.UserGet(ctx, userID)
			require.NoError(err)

			// assert changes
			if updatedColumns.ID != 0 {
				require.Equal(
					fetchedUser.ID,
					updatedColumns.ID+i,
				)
			} else {
				require.Equal(fetchedUser.ID, user.ID)
			}
			if updatedColumns.Email != "" {
				require.Equal(
					fetchedUser.Email,
					updatedColumns.Email,
				)
			} else {
				require.Equal(fetchedUser.Email, user.Email)
			}
			if updatedColumns.HashedPassword != "" {
				require.Equal(
					fetchedUser.HashedPassword,
					updatedColumns.HashedPassword,
				)
			} else {
				require.Equal(fetchedUser.HashedPassword, user.HashedPassword)
			}
			if updatedColumns.FirstName.Valid {
				require.Equal(
					fetchedUser.FirstName,
					updatedColumns.FirstName,
				)
			} else {
				require.Equal(fetchedUser.FirstName, user.FirstName)
			}
			if updatedColumns.LastName.Valid {
				require.Equal(
					fetchedUser.LastName,
					updatedColumns.LastName,
				)
			} else {
				require.Equal(fetchedUser.LastName, user.LastName)
			}
			if updatedColumns.Bio.Valid {
				require.Equal(fetchedUser.Bio, updatedColumns.Bio)
			} else {
				require.Equal(fetchedUser.Bio, user.Bio)
			}
			if updatedColumns.Birthdate.Valid {
				require.Equal(
					fetchedUser.Birthdate.Time.Format(time.RFC3339[:10]),
					updatedColumns.Birthdate.Time.Format(time.RFC3339[:10]),
				)
			} else {
				require.Equal(
					fetchedUser.Birthdate.Time.Format(time.RFC3339[:10]),
					user.Birthdate.Time.Format(time.RFC3339[:10]),
				)
			}
			if updatedColumns.Joindate != (time.Time{}) {
				require.Equal(
					fetchedUser.Joindate.Format(time.RFC3339[:10]),
					updatedColumns.Joindate.Format(time.RFC3339[:10]),
				)
			} else {
				require.Equal(
					fetchedUser.Joindate.Format(time.RFC3339[:10]),
					user.Joindate.Format(time.RFC3339[:10]),
				)
			}
		})
	}
}
