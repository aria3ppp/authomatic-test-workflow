package server_test

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/aria3ppp/watch-server/internal/config"
	"github.com/aria3ppp/watch-server/internal/dto"
	"github.com/aria3ppp/watch-server/internal/models"
	"github.com/aria3ppp/watch-server/internal/server/response"
	"github.com/aria3ppp/watch-server/internal/testutils"
	"github.com/gavv/httpexpect/v2"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	prequire "github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestHandleMovieGet(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultUser,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/{id}/"
	method := http.MethodGet

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// movie not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	createTime := time.Now()

	// add a new movie
	movieCreateReq := &dto.MovieCreateRequest{
		Title:        "movie",
		DateReleased: testutils.Date(2000, 1, 1),
	}
	movieID, err := appInstance.MovieCreate(
		ctx,
		defaults.user.id,
		movieCreateReq,
	)
	require.NoError(err)

	gotMovie, err := appInstance.MovieGet(ctx, movieID)
	require.NoError(err)

	require.GreaterOrEqual(gotMovie.ContributedAt, createTime)

	payload := &models.Film{
		ID:            movieID,
		Title:         movieCreateReq.Title,
		Descriptions:  movieCreateReq.Descriptions,
		DateReleased:  movieCreateReq.DateReleased,
		Duration:      movieCreateReq.Duration,
		SeriesID:      null.Int{},
		SeasonNumber:  null.Int{},
		EpisodeNumber: null.Int{},
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotMovie.ContributedAt,
	}

	// get movie
	e.Request(method, path).
		WithPath("id", movieID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(payload))
}

func TestHandleMoviesGetAll(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultUser,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/"
	method := http.MethodGet

	// no movie
	e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	// add movies
	movieCreateReqs := []*dto.MovieCreateRequest{
		{
			Title:        "m1",
			DateReleased: testutils.Date(2000, 1, 1),
		},
		{
			Title:        "m2",
			DateReleased: testutils.Date(2001, 1, 1),
		},
		{
			Title:        "m3",
			DateReleased: testutils.Date(2002, 1, 1),
		},
		{
			Title:        "m4",
			DateReleased: testutils.Date(2003, 1, 1),
		},
		{
			Title:        "m5",
			DateReleased: testutils.Date(2004, 1, 1),
		},
	}
	movieIDs := make([]int, len(movieCreateReqs))

	createTime := time.Now()

	for i, req := range movieCreateReqs {
		movieIDs[i], err = appInstance.MovieCreate(
			ctx,
			defaults.user.id,
			req,
		)
		require.NoError(err)
	}

	gotMovies, total, err := appInstance.MoviesGetAll(
		ctx,
		0,
		config.Config.Pagination.PageSize.MaxValue,
	)
	require.NoError(err)

	items := make([]*models.Film, len(gotMovies))

	for i := range gotMovies {
		require.GreaterOrEqual(gotMovies[i].ContributedAt, createTime)

		items[i] = &models.Film{
			ID:            movieIDs[i],
			Title:         movieCreateReqs[i].Title,
			Descriptions:  movieCreateReqs[i].Descriptions,
			DateReleased:  movieCreateReqs[i].DateReleased,
			Duration:      movieCreateReqs[i].Duration,
			SeriesID:      null.Int{},
			SeasonNumber:  null.Int{},
			EpisodeNumber: null.Int{},
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotMovies[i].ContributedAt,
		}
	}

	// get all movies
	e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, items, total))
}

func TestHandleMovieCreate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultUser,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/"
	method := http.MethodPost

	// create movie
	movieCreateReq := &dto.MovieCreateRequest{
		Title:        "movie",
		DateReleased: testutils.Date(2000, 1, 1),
	}

	createTime := time.Now()

	rawMovieID := e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(movieCreateReq).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("status", response.StatusOK.String()).
		Value("payload").Number().Ge(0).Raw()

	movieID := int(rawMovieID)

	// check movie created

	gotMovie, err := appInstance.MovieGet(ctx, movieID)
	require.NoError(err)

	testutils.SetTimeLocation(
		&movieCreateReq.DateReleased,
		gotMovie.DateReleased.Location(),
	)

	require.GreaterOrEqual(gotMovie.ContributedAt, createTime)

	require.Equal(
		&models.Film{
			ID:            movieID,
			Title:         movieCreateReq.Title,
			Descriptions:  movieCreateReq.Descriptions,
			DateReleased:  movieCreateReq.DateReleased,
			Duration:      movieCreateReq.Duration,
			SeriesID:      null.Int{},
			SeasonNumber:  null.Int{},
			EpisodeNumber: null.Int{},
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotMovie.ContributedAt,
		},
		gotMovie,
	)
}

func TestHandleMovieCreate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/movie/"
	method := http.MethodPost

	timeNow := time.Now()

	type Length struct {
		shorterThanMin string
		longerThanMax  string
	}
	testDatas := struct {
		title struct {
			length Length
		}
		descriptions struct {
			length Length
		}
		dateReleased struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
		duration struct {
			lesserThanMinValue  int
			greaterThanMaxValue int
		}
	}{
		title: struct{ length Length }{
			length: Length{
				shorterThanMin: "t",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Film.Title.MaxLength,
				),
			},
		},
		descriptions: struct{ length Length }{
			length: Length{
				shorterThanMin: "d",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Film.Descriptions.MaxLength,
				),
			},
		},
		dateReleased: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Film.DateReleased.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Film.DateReleased.MinValue.Month,
				),
				config.Config.Validation.Film.DateReleased.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
		duration: struct {
			lesserThanMinValue  int
			greaterThanMaxValue int
		}{
			lesserThanMinValue:  config.Config.Validation.Film.Duraion.MinLength - 1,
			greaterThanMaxValue: config.Config.Validation.Film.Duraion.MaxLength + 1,
		},
	}

	testCases := []struct {
		name      string
		req       dto.MovieCreateRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req:  dto.MovieCreateRequest{},
			expErrors: validation.Errors{
				"title":         validation.ErrRequired,
				"date_released": validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.MovieCreateRequest{
				Title:        "",
				Descriptions: null.StringFrom(""),
				DateReleased: time.Time{},
				Duration:     null.IntFrom(0),
			},
			// reauired if submitted (null.Valid == true)
			expErrors: validation.Errors{
				"title":         validation.ErrRequired,
				"descriptions":  validation.ErrRequired,
				"date_released": validation.ErrRequired,
				"duration":      validation.ErrRequired,
			},
		},

		{
			name: "tc3",
			req: dto.MovieCreateRequest{
				Title: testDatas.title.length.shorterThanMin,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.shorterThanMin,
				),
				DateReleased: testDatas.dateReleased.lesserThanMinValue,
				Duration: null.IntFrom(
					testDatas.duration.lesserThanMinValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Title.MinLength,
						"max": config.Config.Validation.Film.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Descriptions.MinLength,
						"max": config.Config.Validation.Film.Descriptions.MaxLength,
					},
				),
				"date_released": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Film.DateReleased.MinValue.Year,
							time.Month(
								config.Config.Validation.Film.DateReleased.MinValue.Month,
							),
							config.Config.Validation.Film.DateReleased.MinValue.Day,
						),
					},
				),
				"duration": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": config.Config.Validation.Film.Duraion.MinLength,
					},
				),
			},
		},

		{
			name: "tc4",
			req: dto.MovieCreateRequest{
				Title: testDatas.title.length.longerThanMax,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.longerThanMax,
				),
				DateReleased: testDatas.dateReleased.greaterThanMaxValue,
				Duration: null.IntFrom(
					testDatas.duration.greaterThanMaxValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Title.MinLength,
						"max": config.Config.Validation.Film.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Descriptions.MinLength,
						"max": config.Config.Validation.Film.Descriptions.MaxLength,
					},
				),
				"date_released": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
					},
				),
				"duration": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": config.Config.Validation.Film.Duraion.MaxLength,
					},
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			e.Request(method, path).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(tc.req).
				Expect().
				Status(http.StatusBadRequest).
				JSON().
				Equal(response.Error(
					response.StatusInvalidRequest,
					tc.expErrors.Error(),
				))
		})
	}
}

func TestHandleMovieUpdate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/{id}"
	method := http.MethodPatch

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.MovieUpdateRequest{}).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// movie not found
	e.Request(method, path).
		WithPath("id", 9999999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.MovieUpdateRequest{}).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	updates := []struct {
		name string
		req  *dto.MovieUpdateRequest
	}{
		{
			name: "u1",
			req:  &dto.MovieUpdateRequest{},
		},
		{
			name: "u2",
			req: &dto.MovieUpdateRequest{
				Title: null.StringFrom("updated_title"),
			},
		},
		{
			name: "u3",
			req: &dto.MovieUpdateRequest{
				Descriptions: null.StringFrom("updated_description"),
			},
		},
		{
			name: "u4",
			req: &dto.MovieUpdateRequest{
				DateReleased: null.TimeFrom(testutils.Date(2000, 1, 1)),
			},
		},
		{
			name: "u5",
			req: &dto.MovieUpdateRequest{
				Duration: null.IntFrom(10 * 60),
			},
		},
		{
			name: "u6",
			req: &dto.MovieUpdateRequest{
				Title:        null.StringFrom("updated_title"),
				Descriptions: null.StringFrom("updated_description"),
				DateReleased: null.TimeFrom(testutils.Date(2000, 1, 1)),
				Duration:     null.IntFrom(10 * 60),
			},
		},
	}

	for i, u := range updates {
		u := u
		i := i
		t.Run(u.name, func(t *testing.T) {
			require := prequire.New(t)

			// insert movie
			movieID, err := appInstance.MovieCreate(
				ctx,
				defaults.user.id,
				&dto.MovieCreateRequest{
					Title:        "movie" + strconv.Itoa(i),
					DateReleased: testutils.Date(1900, 3, 14),
				},
			)
			require.NoError(err)

			gotMovieBeforeUpdate, err := appInstance.MovieGet(ctx, movieID)
			require.NoError(err)

			updateTime := time.Now()

			// update
			e.Request(method, path).
				WithPath("id", movieID).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(u.req).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Equal(response.OK(nil))

			// check updated fields
			gotMovieAfterUpdate, err := appInstance.MovieGet(ctx, movieID)
			require.NoError(err)

			require.GreaterOrEqual(
				gotMovieAfterUpdate.ContributedAt,
				updateTime,
			)

			updatedMovie := &models.Film{}
			updatedMovie.ID = movieID
			if u.req.Title.Valid {
				updatedMovie.Title = u.req.Title.String
			} else {
				updatedMovie.Title = gotMovieBeforeUpdate.Title
			}
			if u.req.Descriptions.Valid {
				updatedMovie.Descriptions = u.req.Descriptions
			} else {
				updatedMovie.Descriptions = gotMovieBeforeUpdate.Descriptions
			}
			if u.req.DateReleased.Valid {
				updatedMovie.DateReleased = u.req.DateReleased.Time
			} else {
				updatedMovie.DateReleased = gotMovieBeforeUpdate.DateReleased
			}
			if u.req.Duration.Valid {
				updatedMovie.Duration = u.req.Duration
			} else {
				updatedMovie.Duration = gotMovieBeforeUpdate.Duration
			}
			updatedMovie.Invalidation = null.String{}
			updatedMovie.ContributedBy = defaults.user.id
			updatedMovie.ContributedAt = gotMovieAfterUpdate.ContributedAt

			testutils.SetTimeLocation(
				&updatedMovie.DateReleased,
				gotMovieAfterUpdate.DateReleased.Location(),
			)

			require.Equal(updatedMovie, gotMovieAfterUpdate)
		})
	}
}

func TestHandleMovieUpdate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/movie/{id}"
	method := http.MethodPatch

	timeNow := time.Now()

	type Length struct {
		shorterThanMin string
		longerThanMax  string
	}
	testDatas := struct {
		title struct {
			length Length
		}
		descriptions struct {
			length Length
		}
		dateReleased struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
		duration struct {
			lesserThanMinValue  int
			greaterThanMaxValue int
		}
	}{
		title: struct{ length Length }{
			length: Length{
				shorterThanMin: "t",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Film.Title.MaxLength,
				),
			},
		},
		descriptions: struct{ length Length }{
			length: Length{
				shorterThanMin: "d",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Film.Descriptions.MaxLength,
				),
			},
		},
		dateReleased: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Film.DateReleased.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Film.DateReleased.MinValue.Month,
				),
				config.Config.Validation.Film.DateReleased.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
		duration: struct {
			lesserThanMinValue  int
			greaterThanMaxValue int
		}{
			lesserThanMinValue:  config.Config.Validation.Film.Duraion.MinLength - 1,
			greaterThanMaxValue: config.Config.Validation.Film.Duraion.MaxLength + 1,
		},
	}

	testCases := []struct {
		name      string
		req       dto.MovieUpdateRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req: dto.MovieUpdateRequest{
				Title:        null.StringFrom(""),
				Descriptions: null.StringFrom(""),
				DateReleased: null.TimeFrom(time.Time{}),
				Duration:     null.IntFrom(0),
			},
			// reauired if submitted (null.Valid == true)
			expErrors: validation.Errors{
				"title":         validation.ErrRequired,
				"descriptions":  validation.ErrRequired,
				"date_released": validation.ErrRequired,
				"duration":      validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.MovieUpdateRequest{
				Title: null.StringFrom(
					testDatas.title.length.shorterThanMin,
				),
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.shorterThanMin,
				),
				DateReleased: null.TimeFrom(
					testDatas.dateReleased.lesserThanMinValue,
				),
				Duration: null.IntFrom(
					testDatas.duration.lesserThanMinValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Title.MinLength,
						"max": config.Config.Validation.Film.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Descriptions.MinLength,
						"max": config.Config.Validation.Film.Descriptions.MaxLength,
					},
				),
				"date_released": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Film.DateReleased.MinValue.Year,
							time.Month(
								config.Config.Validation.Film.DateReleased.MinValue.Month,
							),
							config.Config.Validation.Film.DateReleased.MinValue.Day,
						),
					},
				),
				"duration": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": config.Config.Validation.Film.Duraion.MinLength,
					},
				),
			},
		},

		{
			name: "tc3",
			req: dto.MovieUpdateRequest{
				Title: null.StringFrom(testDatas.title.length.longerThanMax),
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.longerThanMax,
				),
				DateReleased: null.TimeFrom(
					testDatas.dateReleased.greaterThanMaxValue,
				),
				Duration: null.IntFrom(
					testDatas.duration.greaterThanMaxValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Title.MinLength,
						"max": config.Config.Validation.Film.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Film.Descriptions.MinLength,
						"max": config.Config.Validation.Film.Descriptions.MaxLength,
					},
				),
				"date_released": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
					},
				),
				"duration": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": config.Config.Validation.Film.Duraion.MaxLength,
					},
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			e.Request(method, path).
				WithPath("id", 1).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(tc.req).
				Expect().
				Status(http.StatusBadRequest).
				JSON().
				Equal(response.Error(
					response.StatusInvalidRequest,
					tc.expErrors.Error(),
				))
		})
	}
}

func TestHandleMovieInvalidate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/{id}"
	method := http.MethodDelete

	invalidationRequest := &dto.InvalidationRequest{
		Invalidation: "invalidation",
	}

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// movie not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// invalidate movie
	movieCreateReq := &dto.MovieCreateRequest{
		Title:        "movie",
		DateReleased: testutils.Date(1900, 3, 14),
	}
	movieID, err := appInstance.MovieCreate(
		ctx,
		defaults.user.id,
		movieCreateReq,
	)
	require.NoError(err)

	e.Request(method, path).
		WithPath("id", movieID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check movie invalidated
	gotInvalidatedMovie, err := appInstance.MovieGet(ctx, movieID)
	require.NoError(err)

	testutils.SetTimeLocation(
		&movieCreateReq.DateReleased,
		gotInvalidatedMovie.DateReleased.Location(),
	)

	require.Equal(
		&models.Film{
			ID:            movieID,
			Title:         movieCreateReq.Title,
			Descriptions:  movieCreateReq.Descriptions,
			DateReleased:  movieCreateReq.DateReleased,
			Duration:      movieCreateReq.Duration,
			SeriesID:      null.Int{},
			SeasonNumber:  null.Int{},
			EpisodeNumber: null.Int{},
			Invalidation:  null.StringFrom(invalidationRequest.Invalidation),
			ContributedBy: defaults.user.id,
			ContributedAt: gotInvalidatedMovie.ContributedAt,
		},
		gotInvalidatedMovie,
	)
}

func TestHandleMovieInvalidate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/movie/{id}"
	method := http.MethodDelete

	testCases := []struct {
		name      string
		req       dto.InvalidationRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req:  dto.InvalidationRequest{},
			expErrors: validation.Errors{
				"invalidation": validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.InvalidationRequest{
				Invalidation: "i",
			},
			expErrors: validation.Errors{
				"invalidation": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Request.Invalidation.MinLength,
						"max": config.Config.Validation.Request.Invalidation.MaxLength,
					},
				),
			},
		},

		{
			name: "tc3",
			req: dto.InvalidationRequest{
				Invalidation: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Request.Invalidation.MaxLength,
				),
			},
			expErrors: validation.Errors{
				"invalidation": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Request.Invalidation.MinLength,
						"max": config.Config.Validation.Request.Invalidation.MaxLength,
					},
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.New(t, server.URL)

			e.Request(method, path).
				WithPath("id", 1).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(tc.req).
				Expect().
				Status(http.StatusBadRequest).
				JSON().
				Object().
				Equal(response.Error(
					response.StatusInvalidRequest,
					tc.expErrors.Error(),
				))
		})
	}
}

func TestHandleMovieAuditsGetAll(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/movie/{id}/audits/"
	method := http.MethodGet

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// movie not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	createTime := time.Now()

	// add movie
	movieCreateReq := &dto.MovieCreateRequest{
		Title:        "movie",
		DateReleased: testutils.Date(2000, 1, 1),
	}
	movieID, err := appInstance.MovieCreate(
		ctx,
		defaults.user.id,
		movieCreateReq,
	)
	require.NoError(err)

	gotMovie, err := appInstance.MovieGet(ctx, movieID)
	require.NoError(err)

	require.GreaterOrEqual(gotMovie.ContributedAt, createTime)

	expMovieUpdateAudit := &models.FilmsAudit{
		ID:            movieID,
		Title:         movieCreateReq.Title,
		Descriptions:  movieCreateReq.Descriptions,
		DateReleased:  movieCreateReq.DateReleased,
		Duration:      movieCreateReq.Duration,
		SeriesID:      null.Int{},
		SeasonNumber:  null.Int{},
		EpisodeNumber: null.Int{},
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotMovie.ContributedAt,
	}

	// no audits
	e.Request(method, path).
		WithPath("id", movieID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	updateTime := time.Now()

	// update the movie
	movieUpdateReq := &dto.MovieUpdateRequest{
		Title:        null.StringFrom("updated title"),
		Descriptions: null.StringFrom("updated descriptions"),
		DateReleased: null.TimeFrom(
			testutils.Date(2005, 11, 14),
		),
		Duration: null.IntFrom(10 * 60),
	}
	err = appInstance.MovieUpdate(
		ctx,
		movieID,
		defaults.user.id,
		movieUpdateReq,
	)
	require.NoError(err)

	gotMovie, err = appInstance.MovieGet(ctx, movieID)
	require.NoError(err)

	require.GreaterOrEqual(gotMovie.ContributedAt, updateTime)

	expMovieInvalidationAudit := &models.FilmsAudit{
		ID:            movieID,
		Title:         movieUpdateReq.Title.String,
		Descriptions:  movieUpdateReq.Descriptions,
		DateReleased:  movieUpdateReq.DateReleased.Time,
		Duration:      movieUpdateReq.Duration,
		SeriesID:      null.Int{},
		SeasonNumber:  null.Int{},
		EpisodeNumber: null.Int{},
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotMovie.ContributedAt,
	}

	// get update audits
	e.Request(method, path).
		WithPath("id", movieID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.FilmsAudit{expMovieUpdateAudit}, 1))

	// invalidate the movie
	err = appInstance.MovieInvalidate(
		ctx,
		movieID,
		defaults.user.id,
		&dto.InvalidationRequest{Invalidation: "invalidation"},
	)
	require.NoError(err)

	// get invalidation audits
	e.Request(method, path).
		WithPath("id", movieID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.FilmsAudit{expMovieInvalidationAudit, expMovieUpdateAudit}, 2))
}
