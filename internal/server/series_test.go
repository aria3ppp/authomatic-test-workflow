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

func TestHandleSeriesGet(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}"
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

	// series not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	upsertTime := time.Now()

	// add a new series
	seriesUpsertReq := &dto.SeriesCreateRequest{
		Title:       "series",
		DateStarted: testutils.Date(2000, 1, 1),
	}
	seriesID, err := appInstance.SeriesCreate(
		ctx,
		defaults.user.id,
		seriesUpsertReq,
	)
	require.NoError(err)

	gotSeries, err := appInstance.SeriesGet(ctx, seriesID)
	require.NoError(err)

	require.GreaterOrEqual(gotSeries.ContributedAt, upsertTime)

	payload := &models.Series{
		ID:            seriesID,
		Title:         seriesUpsertReq.Title,
		Descriptions:  seriesUpsertReq.Descriptions,
		DateStarted:   seriesUpsertReq.DateStarted,
		DateEnded:     seriesUpsertReq.DateEnded,
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotSeries.ContributedAt,
	}

	// get series
	e.Request(method, path).
		WithPath("id", seriesID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(payload))
}

func TestHandleSeriesesGetAll(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/"
	method := http.MethodGet

	// no serieses
	e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	// add serieses
	seriesUpsertReqs := []*dto.SeriesCreateRequest{
		{
			Title:       "s1",
			DateStarted: testutils.Date(2000, 1, 1),
		},
		{
			Title:       "s2",
			DateStarted: testutils.Date(2001, 1, 1),
		},
		{
			Title:       "s3",
			DateStarted: testutils.Date(2002, 1, 1),
		},
		{
			Title:       "s4",
			DateStarted: testutils.Date(2003, 1, 1),
		},
		{
			Title:       "s5",
			DateStarted: testutils.Date(2004, 1, 1),
		},
	}
	seriesIDs := make([]int, len(seriesUpsertReqs))

	upsertTime := time.Now()

	for i, req := range seriesUpsertReqs {
		seriesIDs[i], err = appInstance.SeriesCreate(
			ctx,
			defaults.user.id,
			req,
		)
		require.NoError(err)
	}

	gotSerieses, total, err := appInstance.SeriesesGetAll(
		ctx,
		0,
		config.Config.Pagination.PageSize.MaxValue,
	)
	require.NoError(err)

	items := make([]*models.Series, len(gotSerieses))

	for i := range gotSerieses {
		require.GreaterOrEqual(gotSerieses[i].ContributedAt, upsertTime)

		items[i] = &models.Series{
			ID:            seriesIDs[i],
			Title:         seriesUpsertReqs[i].Title,
			Descriptions:  seriesUpsertReqs[i].Descriptions,
			DateStarted:   seriesUpsertReqs[i].DateStarted,
			DateEnded:     seriesUpsertReqs[i].DateEnded,
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotSerieses[i].ContributedAt,
		}
	}

	// get all series
	e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, items, total))
}

func TestHandleSeriesCreate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/"
	method := http.MethodPost

	// create series
	seriesCreateReq := &dto.SeriesCreateRequest{
		Title:       "series",
		DateStarted: testutils.Date(2000, 1, 1),
	}

	createTime := time.Now()

	rawSeriesID := e.Request(method, path).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(seriesCreateReq).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("status", response.StatusOK.String()).
		Value("payload").Number().Ge(0).Raw()

	seriesID := int(rawSeriesID)

	// check series created

	gotSeries, err := appInstance.SeriesGet(ctx, seriesID)
	require.NoError(err)

	testutils.SetTimeLocation(
		&seriesCreateReq.DateStarted,
		gotSeries.DateStarted.Location(),
	)
	testutils.SetTimeLocation(
		&seriesCreateReq.DateEnded.Time,
		gotSeries.DateEnded.Time.Location(),
	)

	require.GreaterOrEqual(gotSeries.ContributedAt, createTime)

	require.Equal(
		&models.Series{
			ID:            seriesID,
			Title:         seriesCreateReq.Title,
			Descriptions:  seriesCreateReq.Descriptions,
			DateStarted:   seriesCreateReq.DateStarted,
			DateEnded:     seriesCreateReq.DateEnded,
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotSeries.ContributedAt,
		},
		gotSeries,
	)
}

func TestHandleSeriesCreate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/"
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
		dateStarted struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
		dateEnded struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
	}{
		title: struct{ length Length }{
			length: Length{
				shorterThanMin: "t",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Series.Title.MaxLength,
				),
			},
		},
		descriptions: struct{ length Length }{
			length: Length{
				shorterThanMin: "d",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Series.Descriptions.MaxLength,
				),
			},
		},
		dateStarted: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Series.DateStarted.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Series.DateStarted.MinValue.Month,
				),
				config.Config.Validation.Series.DateStarted.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
		dateEnded: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Series.DateEnded.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Series.DateEnded.MinValue.Month,
				),
				config.Config.Validation.Series.DateEnded.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
	}

	testCases := []struct {
		name      string
		req       dto.SeriesCreateRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req:  dto.SeriesCreateRequest{},
			expErrors: validation.Errors{
				"title":        validation.ErrRequired,
				"date_started": validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.SeriesCreateRequest{
				Title:        "",
				Descriptions: null.StringFrom(""),
				DateStarted:  time.Time{},
				DateEnded:    null.TimeFrom(time.Time{}),
			},
			// reauired if submitted (null.Valid == true)
			expErrors: validation.Errors{
				"title":        validation.ErrRequired,
				"descriptions": validation.ErrRequired,
				"date_started": validation.ErrRequired,
				"date_ended":   validation.ErrRequired,
			},
		},

		{
			name: "tc3",
			req: dto.SeriesCreateRequest{
				Title: testDatas.title.length.shorterThanMin,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.shorterThanMin,
				),
				DateStarted: testDatas.dateStarted.lesserThanMinValue,
				DateEnded: null.TimeFrom(
					testDatas.dateEnded.lesserThanMinValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Title.MinLength,
						"max": config.Config.Validation.Series.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Descriptions.MinLength,
						"max": config.Config.Validation.Series.Descriptions.MaxLength,
					},
				),
				"date_started": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Series.DateStarted.MinValue.Year,
							time.Month(
								config.Config.Validation.Series.DateStarted.MinValue.Month,
							),
							config.Config.Validation.Series.DateStarted.MinValue.Day,
						),
					},
				),
				"date_ended": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Series.DateEnded.MinValue.Year,
							time.Month(
								config.Config.Validation.Series.DateEnded.MinValue.Month,
							),
							config.Config.Validation.Series.DateEnded.MinValue.Day,
						),
					},
				),
			},
		},

		{
			name: "tc4",
			req: dto.SeriesCreateRequest{
				Title: testDatas.title.length.longerThanMax,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.longerThanMax,
				),
				DateStarted: testDatas.dateStarted.greaterThanMaxValue,
				DateEnded: null.TimeFrom(
					testDatas.dateEnded.greaterThanMaxValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Title.MinLength,
						"max": config.Config.Validation.Series.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Descriptions.MinLength,
						"max": config.Config.Validation.Series.Descriptions.MaxLength,
					},
				),
				"date_started": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
					},
				),
				"date_ended": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
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

func TestHandleSeriesUpdate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}"
	method := http.MethodPatch

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.SeriesUpdateRequest{}).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// series not found
	e.Request(method, path).
		WithPath("id", 9999999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.SeriesUpdateRequest{}).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	updates := []struct {
		name string
		req  *dto.SeriesUpdateRequest
	}{
		{
			name: "u1",
			req:  &dto.SeriesUpdateRequest{},
		},
		{
			name: "u2",
			req: &dto.SeriesUpdateRequest{
				Title: null.StringFrom("updated_title"),
			},
		},
		{
			name: "u3",
			req: &dto.SeriesUpdateRequest{
				Descriptions: null.StringFrom("updated_description"),
			},
		},
		{
			name: "u4",
			req: &dto.SeriesUpdateRequest{
				DateStarted: null.TimeFrom(testutils.Date(2000, 1, 1)),
			},
		},
		{
			name: "u5",
			req: &dto.SeriesUpdateRequest{
				DateEnded: null.TimeFrom(testutils.Date(2008, 8, 11)),
			},
		},
		{
			name: "u6",
			req: &dto.SeriesUpdateRequest{
				Title:        null.StringFrom("updated_title"),
				Descriptions: null.StringFrom("updated_description"),
				DateStarted:  null.TimeFrom(testutils.Date(2000, 1, 1)),
				DateEnded:    null.TimeFrom(testutils.Date(2008, 8, 11)),
			},
		},
	}

	for i, u := range updates {
		u := u
		i := i
		t.Run(u.name, func(t *testing.T) {
			require := prequire.New(t)

			// insert series
			seriesID, err := appInstance.SeriesCreate(
				ctx,
				defaults.user.id,
				&dto.SeriesCreateRequest{
					Title:       "series" + strconv.Itoa(i),
					DateStarted: testutils.Date(1900, 3, 14),
				},
			)
			require.NoError(err)

			gotSeriesBeforeUpdate, err := appInstance.SeriesGet(ctx, seriesID)
			require.NoError(err)

			updateTime := time.Now()

			// update
			e.Request(method, path).
				WithPath("id", seriesID).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(u.req).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Equal(response.OK(nil))

			// check updated fields
			gotSeriesAfterUpdate, err := appInstance.SeriesGet(ctx, seriesID)
			require.NoError(err)

			require.GreaterOrEqual(
				gotSeriesAfterUpdate.ContributedAt,
				updateTime,
			)

			updatedSeries := &models.Series{}
			updatedSeries.ID = seriesID
			if u.req.Title.Valid {
				updatedSeries.Title = u.req.Title.String
			} else {
				updatedSeries.Title = gotSeriesBeforeUpdate.Title
			}
			if u.req.Descriptions.Valid {
				updatedSeries.Descriptions = u.req.Descriptions
			} else {
				updatedSeries.Descriptions = gotSeriesBeforeUpdate.Descriptions
			}
			if u.req.DateStarted.Valid {
				updatedSeries.DateStarted = u.req.DateStarted.Time
			} else {
				updatedSeries.DateStarted = gotSeriesBeforeUpdate.DateStarted
			}
			if u.req.DateEnded.Valid {
				updatedSeries.DateEnded = u.req.DateEnded
			} else {
				updatedSeries.DateEnded = gotSeriesBeforeUpdate.DateEnded
			}
			updatedSeries.Invalidation = null.String{}
			updatedSeries.ContributedBy = defaults.user.id
			updatedSeries.ContributedAt = gotSeriesAfterUpdate.ContributedAt

			testutils.SetTimeLocation(
				&updatedSeries.DateStarted,
				gotSeriesAfterUpdate.DateStarted.Location(),
			)
			testutils.SetTimeLocation(
				&updatedSeries.DateEnded.Time,
				gotSeriesAfterUpdate.DateEnded.Time.Location(),
			)

			require.Equal(updatedSeries, gotSeriesAfterUpdate)
		})
	}
}

func TestHandleSeriesUpdate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/{id}"
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
		dateStarted struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
		dateEnded struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}
	}{
		title: struct{ length Length }{
			length: Length{
				shorterThanMin: "t",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Series.Title.MaxLength,
				),
			},
		},
		descriptions: struct{ length Length }{
			length: Length{
				shorterThanMin: "d",
				longerThanMax: testutils.GenerateStringLongerThanMaxLength(
					config.Config.Validation.Series.Descriptions.MaxLength,
				),
			},
		},
		dateStarted: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Series.DateStarted.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Series.DateStarted.MinValue.Month,
				),
				config.Config.Validation.Series.DateStarted.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
		dateEnded: struct {
			lesserThanMinValue  time.Time
			greaterThanMaxValue time.Time
		}{
			lesserThanMinValue: testutils.Date(
				config.Config.Validation.Series.DateEnded.MinValue.Year-1,
				time.Month(
					config.Config.Validation.Series.DateEnded.MinValue.Month,
				),
				config.Config.Validation.Series.DateEnded.MinValue.Day,
			),
			greaterThanMaxValue: testutils.Date(
				timeNow.Year(), timeNow.Month(), timeNow.Day(),
			).Add(time.Hour),
		},
	}

	testCases := []struct {
		name      string
		req       dto.SeriesUpdateRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req: dto.SeriesUpdateRequest{
				Title:        null.StringFrom(""),
				Descriptions: null.StringFrom(""),
				DateStarted:  null.TimeFrom(time.Time{}),
				DateEnded:    null.TimeFrom(time.Time{}),
			},
			// reauired if submitted (null.Valid == true)
			expErrors: validation.Errors{
				"title":        validation.ErrRequired,
				"descriptions": validation.ErrRequired,
				"date_started": validation.ErrRequired,
				"date_ended":   validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.SeriesUpdateRequest{
				Title: null.StringFrom(
					testDatas.title.length.shorterThanMin,
				),
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.shorterThanMin,
				),
				DateStarted: null.TimeFrom(
					testDatas.dateStarted.lesserThanMinValue,
				),
				DateEnded: null.TimeFrom(
					testDatas.dateEnded.lesserThanMinValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Title.MinLength,
						"max": config.Config.Validation.Series.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Descriptions.MinLength,
						"max": config.Config.Validation.Series.Descriptions.MaxLength,
					},
				),
				"date_started": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Series.DateStarted.MinValue.Year,
							time.Month(
								config.Config.Validation.Series.DateStarted.MinValue.Month,
							),
							config.Config.Validation.Series.DateStarted.MinValue.Day,
						),
					},
				),
				"date_ended": validation.ErrMinGreaterEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							config.Config.Validation.Series.DateEnded.MinValue.Year,
							time.Month(
								config.Config.Validation.Series.DateEnded.MinValue.Month,
							),
							config.Config.Validation.Series.DateEnded.MinValue.Day,
						),
					},
				),
			},
		},

		{
			name: "tc3",
			req: dto.SeriesUpdateRequest{
				Title: null.StringFrom(testDatas.title.length.longerThanMax),
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.longerThanMax,
				),
				DateStarted: null.TimeFrom(
					testDatas.dateStarted.greaterThanMaxValue,
				),
				DateEnded: null.TimeFrom(
					testDatas.dateEnded.greaterThanMaxValue,
				),
			},
			expErrors: validation.Errors{
				"title": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Title.MinLength,
						"max": config.Config.Validation.Series.Title.MaxLength,
					},
				),
				"descriptions": validation.ErrLengthOutOfRange.SetParams(
					map[string]any{
						"min": config.Config.Validation.Series.Descriptions.MinLength,
						"max": config.Config.Validation.Series.Descriptions.MaxLength,
					},
				),
				"date_started": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
					},
				),
				"date_ended": validation.ErrMaxLessEqualThanRequired.SetParams(
					map[string]any{
						"threshold": testutils.Date(
							time.Now().Year(),
							time.Now().Month(),
							time.Now().Day(),
						),
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

func TestHandleSeriesInvalidate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}"
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

	// series not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// invalidate series
	seriesCreateReq := &dto.SeriesCreateRequest{
		Title:       "series",
		DateStarted: testutils.Date(1900, 3, 14),
	}
	seriesID, err := appInstance.SeriesCreate(
		ctx,
		defaults.user.id,
		seriesCreateReq,
	)
	require.NoError(err)

	e.Request(method, path).
		WithPath("id", seriesID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check series invalidated
	gotInvalidatedSeries, err := appInstance.SeriesGet(ctx, seriesID)
	require.NoError(err)

	testutils.SetTimeLocation(
		&seriesCreateReq.DateStarted,
		gotInvalidatedSeries.DateStarted.Location(),
	)
	testutils.SetTimeLocation(
		&seriesCreateReq.DateEnded.Time,
		gotInvalidatedSeries.DateEnded.Time.Location(),
	)

	require.Equal(
		&models.Series{
			ID:            seriesID,
			Title:         seriesCreateReq.Title,
			Descriptions:  seriesCreateReq.Descriptions,
			DateStarted:   seriesCreateReq.DateStarted,
			DateEnded:     seriesCreateReq.DateEnded,
			Invalidation:  null.StringFrom(invalidationRequest.Invalidation),
			ContributedBy: defaults.user.id,
			ContributedAt: gotInvalidatedSeries.ContributedAt,
		},
		gotInvalidatedSeries,
	)
}

func TestHandleSeriesInvalidate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/{id}"
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

func TestHandleSeriesAuditsGetAll(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(OptEnableDefaultUser)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/audits"
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

	// series not found
	e.Request(method, path).
		WithPath("id", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	createTime := time.Now()

	// add series
	seriesCreateReq := &dto.SeriesCreateRequest{
		Title:       "series",
		DateStarted: testutils.Date(2000, 1, 1),
	}
	seriesID, err := appInstance.SeriesCreate(
		ctx,
		defaults.user.id,
		seriesCreateReq,
	)
	require.NoError(err)

	gotSeries, err := appInstance.SeriesGet(ctx, seriesID)
	require.NoError(err)

	require.GreaterOrEqual(gotSeries.ContributedAt, createTime)

	expSeriesUpdateAudit := &models.SeriesesAudit{
		ID:            seriesID,
		Title:         seriesCreateReq.Title,
		Descriptions:  seriesCreateReq.Descriptions,
		DateStarted:   seriesCreateReq.DateStarted,
		DateEnded:     seriesCreateReq.DateEnded,
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotSeries.ContributedAt,
	}

	// no audits
	e.Request(method, path).
		WithPath("id", seriesID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	updateTime := time.Now()

	// update the series
	seriesUpdateReq := &dto.SeriesUpdateRequest{
		Title:        null.StringFrom("updated title"),
		Descriptions: null.StringFrom("updated descriptions"),
		DateStarted: null.TimeFrom(
			testutils.Date(2005, 11, 14),
		),
		DateEnded: null.TimeFrom(
			testutils.Date(2015, 3, 26),
		),
	}
	err = appInstance.SeriesUpdate(
		ctx,
		seriesID,
		defaults.user.id,
		seriesUpdateReq,
	)
	require.NoError(err)

	gotSeries, err = appInstance.SeriesGet(ctx, seriesID)
	require.NoError(err)

	require.GreaterOrEqual(gotSeries.ContributedAt, updateTime)

	expSeriesInvalidationAudit := &models.SeriesesAudit{
		ID:            seriesID,
		Title:         seriesUpdateReq.Title.String,
		Descriptions:  seriesUpdateReq.Descriptions,
		DateStarted:   seriesUpdateReq.DateStarted.Time,
		DateEnded:     seriesUpdateReq.DateEnded,
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotSeries.ContributedAt,
	}

	// get update audits
	e.Request(method, path).
		WithPath("id", seriesID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.SeriesesAudit{expSeriesUpdateAudit}, 1))

	// invalidate the series
	err = appInstance.SeriesInvalidate(
		ctx,
		seriesID,
		defaults.user.id,
		&dto.InvalidationRequest{Invalidation: "invalidation"},
	)
	require.NoError(err)

	// get invalidation audits
	e.Request(method, path).
		WithPath("id", seriesID).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.SeriesesAudit{expSeriesInvalidationAudit, expSeriesUpdateAudit}, 2))
}
