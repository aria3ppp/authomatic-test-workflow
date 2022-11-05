package server_test

import (
	"context"
	"encoding/json"
	"math"
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

func TestHandleEpisodeGet(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}"
	method := http.MethodGet

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

	// invalid id/numbers
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithPath("ep", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// episode not found
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	putTime := time.Now()

	// add a new episode
	episodePutReq := &dto.EpisodePutRequest{
		Title:        "episode",
		DateReleased: testutils.Date(2000, 1, 1),
	}
	err = appInstance.EpisodePut(
		ctx,
		defaults.series.id, seasonNumber, episodeNumber,
		defaults.user.id,
		episodePutReq,
	)
	require.NoError(err)

	gotEpisode, err := appInstance.EpisodeGet(
		ctx,
		defaults.series.id,
		seasonNumber,
		episodeNumber,
	)
	require.NoError(err)

	require.GreaterOrEqual(gotEpisode.ContributedAt, putTime)

	payload := &models.Film{
		ID:            gotEpisode.ID,
		Title:         episodePutReq.Title,
		Descriptions:  episodePutReq.Descriptions,
		DateReleased:  episodePutReq.DateReleased,
		Duration:      episodePutReq.Duration,
		SeriesID:      null.IntFrom(defaults.series.id),
		SeasonNumber:  null.IntFrom(seasonNumber),
		EpisodeNumber: null.IntFrom(episodeNumber),
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotEpisode.ContributedAt,
	}

	// get episode
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(payload))
}

func TestHandleEpisodesGetAllBySeries(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/episode/"
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

	// no episodes
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	// add episodes
	episodePutReqs := [][]*dto.EpisodePutRequest{
		{
			{
				Title:        "s1e1",
				DateReleased: testutils.Date(2000, 1, 1),
			},
			{
				Title:        "s1e2",
				DateReleased: testutils.Date(2001, 1, 1),
			},
			{
				Title:        "s1e3",
				DateReleased: testutils.Date(2002, 1, 1),
			},
		},
		{
			{
				Title:        "s2e1",
				DateReleased: testutils.Date(2005, 1, 1),
			},
			{
				Title:        "s2e2",
				DateReleased: testutils.Date(2006, 1, 1),
			},
			{
				Title:        "s2e3",
				DateReleased: testutils.Date(2007, 1, 1),
			},
			{
				Title:        "s2e4",
				DateReleased: testutils.Date(2008, 1, 1),
			},
			{
				Title:        "s2e5",
				DateReleased: testutils.Date(2009, 1, 1),
			},
		},
		{
			{
				Title:        "s3e1",
				DateReleased: testutils.Date(2010, 1, 1),
			},
			{
				Title:        "s3e2",
				DateReleased: testutils.Date(2011, 1, 1),
			},
			{
				Title:        "s3e3",
				DateReleased: testutils.Date(2012, 1, 1),
			},
			{
				Title:        "s3e4",
				DateReleased: testutils.Date(2013, 1, 1),
			},
		},
	}
	seasonEpisodeNumbers := make([][]struct{ se, ep int }, len(episodePutReqs))

	putTime := time.Now()

	for s, sreq := range episodePutReqs {
		seasonEpisodeNumbers[s] = make([]struct{ se, ep int }, len(sreq))
		for e, ereq := range sreq {
			se := s + 1
			ep := e + 1
			err = appInstance.EpisodePut(
				ctx,
				defaults.series.id,
				se,
				ep,
				defaults.user.id,
				ereq,
			)
			require.NoError(err)
			seasonEpisodeNumbers[s][e].se = se
			seasonEpisodeNumbers[s][e].ep = ep
		}
	}

	gotEpisodes, total, err := appInstance.EpisodesGetAllBySeries(
		ctx,
		defaults.series.id,
		0,
		config.Config.Pagination.PageSize.MaxValue,
	)
	require.NoError(err)

	items := make([]*models.Film, len(gotEpisodes))

	i := 0
	for s, sreq := range seasonEpisodeNumbers {
		for e := range sreq {
			require.GreaterOrEqual(gotEpisodes[i].ContributedAt, putTime)

			items[i] = &models.Film{
				ID:            gotEpisodes[i].ID,
				Title:         episodePutReqs[s][e].Title,
				Descriptions:  episodePutReqs[s][e].Descriptions,
				DateReleased:  episodePutReqs[s][e].DateReleased,
				Duration:      episodePutReqs[s][e].Duration,
				SeriesID:      null.IntFrom(defaults.series.id),
				SeasonNumber:  null.IntFrom(seasonEpisodeNumbers[s][e].se),
				EpisodeNumber: null.IntFrom(seasonEpisodeNumbers[s][e].ep),
				Invalidation:  null.String{},
				ContributedBy: defaults.user.id,
				ContributedAt: gotEpisodes[i].ContributedAt,
			}

			i++
		}
	}

	// get all episodes
	rawResp := e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, items, total)).
		Raw()

	t.Log("[DEBUG] all episodes from all seasons:")
	json, err := json.MarshalIndent(rawResp, "", "\t")
	require.NoError(err)
	t.Log(string(json))
}

func TestHandleEpisodesGetAllBySeason(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/"
	method := http.MethodGet

	seasonNumber := 1

	// invalid id/number
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// no episodes
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	// add episodes
	episodePutReqs := []*dto.EpisodePutRequest{
		{
			Title:        "e1",
			DateReleased: testutils.Date(2000, 1, 1),
		},
		{
			Title:        "e2",
			DateReleased: testutils.Date(2001, 1, 1),
		},
		{
			Title:        "e3",
			DateReleased: testutils.Date(2002, 1, 1),
		},
		{
			Title:        "e4",
			DateReleased: testutils.Date(2003, 1, 1),
		},
		{
			Title:        "e5",
			DateReleased: testutils.Date(2004, 1, 1),
		},
	}
	episodeNumbers := make([]int, len(episodePutReqs))

	putTime := time.Now()

	for i, req := range episodePutReqs {
		ep := i + 1
		err = appInstance.EpisodePut(
			ctx,
			defaults.series.id,
			seasonNumber,
			ep,
			defaults.user.id,
			req,
		)
		require.NoError(err)
		episodeNumbers[i] = ep
	}

	gotEpisodes, total, err := appInstance.EpisodesGetAllBySeason(
		ctx,
		defaults.series.id,
		seasonNumber,
		0,
		config.Config.Pagination.PageSize.MaxValue,
	)
	require.NoError(err)

	items := make([]*models.Film, len(gotEpisodes))

	for i := range gotEpisodes {
		require.GreaterOrEqual(gotEpisodes[i].ContributedAt, putTime)

		items[i] = &models.Film{
			ID:            gotEpisodes[i].ID,
			Title:         episodePutReqs[i].Title,
			Descriptions:  episodePutReqs[i].Descriptions,
			DateReleased:  episodePutReqs[i].DateReleased,
			Duration:      episodePutReqs[i].Duration,
			SeriesID:      null.IntFrom(defaults.series.id),
			SeasonNumber:  null.IntFrom(seasonNumber),
			EpisodeNumber: null.IntFrom(episodeNumbers[i]),
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotEpisodes[i].ContributedAt,
		}
	}

	// get all episodes
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, items, total))
}

func TestHandleEpisodePut(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
	method := http.MethodPut

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

	episodePutReq := &dto.EpisodePutRequest{
		Title:        "episode",
		DateReleased: testutils.Date(2000, 1, 1),
	}

	// invalid id/number
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithPath("ep", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutReq).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// series not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 999).
		WithPath("ep", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutReq).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// put episode
	putTime := time.Now()

	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutReq).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check episode put in place
	gotEpisode, err := appInstance.EpisodeGet(
		ctx,
		defaults.series.id,
		seasonNumber,
		episodeNumber,
	)
	require.NoError(err)

	testutils.SetTimeLocation(
		&episodePutReq.DateReleased,
		gotEpisode.DateReleased.Location(),
	)

	require.GreaterOrEqual(gotEpisode.ContributedAt, putTime)

	require.Equal(
		&models.Film{
			ID:            gotEpisode.ID,
			Title:         episodePutReq.Title,
			Descriptions:  episodePutReq.Descriptions,
			DateReleased:  episodePutReq.DateReleased,
			Duration:      episodePutReq.Duration,
			SeriesID:      null.IntFrom(defaults.series.id),
			SeasonNumber:  null.IntFrom(seasonNumber),
			EpisodeNumber: null.IntFrom(episodeNumber),
			Invalidation:  null.String{},
			ContributedBy: defaults.user.id,
			ContributedAt: gotEpisode.ContributedAt,
		},
		gotEpisode,
	)
}

func TestHandleEpisodePut_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultSeries)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
	method := http.MethodPut

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

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
		req       dto.EpisodePutRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req:  dto.EpisodePutRequest{},
			expErrors: validation.Errors{
				"title":         validation.ErrRequired,
				"date_released": validation.ErrRequired,
			},
		},

		{
			name: "tc2",
			req: dto.EpisodePutRequest{
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
			req: dto.EpisodePutRequest{
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
			req: dto.EpisodePutRequest{
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
				WithPath("id", defaults.series.id).
				WithPath("se", seasonNumber).
				WithPath("ep", episodeNumber).
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

func TestHandleEpisodesPutAllBySeason(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/"
	method := http.MethodPut

	seasonNumber := 1

	episodePutAllReq := &dto.EpisodesPutAllBySeasonRequest{
		Episodes: []*dto.EpisodePutRequest{
			{
				Title:        "episode1",
				DateReleased: testutils.Date(2000, 1, 1),
			},
			{
				Title:        "episode2",
				DateReleased: testutils.Date(2001, 1, 1),
			},
			{
				Title:        "episode3",
				DateReleased: testutils.Date(2002, 1, 1),
			},
			{
				Title:        "episode4",
				DateReleased: testutils.Date(2003, 1, 1),
			},
			{
				Title:        "episode5",
				DateReleased: testutils.Date(2004, 1, 1),
			},
		},
	}

	// invalid id/number
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutAllReq).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// series not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutAllReq).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// put episodes
	putTime := time.Now()

	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(episodePutAllReq).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check episode put in place
	gotEpisodes, total, err := appInstance.EpisodesGetAllBySeason(
		ctx,
		defaults.series.id,
		seasonNumber,
		0,
		math.MaxInt,
	)
	require.NoError(err)
	require.Equal(len(episodePutAllReq.Episodes), total)

	for i, re := range episodePutAllReq.Episodes {
		testutils.SetTimeLocation(
			&re.DateReleased,
			gotEpisodes[i].DateReleased.Location(),
		)

		require.GreaterOrEqual(gotEpisodes[i].ContributedAt, putTime)

		episodeNumber := i + 1

		require.Equal(
			&models.Film{
				ID:            gotEpisodes[i].ID,
				Title:         re.Title,
				Descriptions:  re.Descriptions,
				DateReleased:  re.DateReleased,
				Duration:      re.Duration,
				SeriesID:      null.IntFrom(defaults.series.id),
				SeasonNumber:  null.IntFrom(seasonNumber),
				EpisodeNumber: null.IntFrom(episodeNumber),
				Invalidation:  null.String{},
				ContributedBy: defaults.user.id,
				ContributedAt: gotEpisodes[i].ContributedAt,
			},
			gotEpisodes[i],
		)
	}
}

func TestHandleEpisodesPutAllBySeason_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(OptEnableDefaultSeries)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/"
	method := http.MethodPut

	seasonNumber := 1

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

	req := &dto.EpisodesPutAllBySeasonRequest{
		Episodes: []*dto.EpisodePutRequest{
			{},
			{
				Title:        "",
				Descriptions: null.StringFrom(""),
				DateReleased: time.Time{},
				Duration:     null.IntFrom(0),
			},
			{
				Title: testDatas.title.length.shorterThanMin,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.shorterThanMin,
				),
				DateReleased: testDatas.dateReleased.lesserThanMinValue,
				Duration: null.IntFrom(
					testDatas.duration.lesserThanMinValue,
				),
			},
			{
				Title: testDatas.title.length.longerThanMax,
				Descriptions: null.StringFrom(
					testDatas.descriptions.length.longerThanMax,
				),
				DateReleased: testDatas.dateReleased.greaterThanMaxValue,
				Duration: null.IntFrom(
					testDatas.duration.greaterThanMaxValue,
				),
			},
		},
	}

	expError := validation.Errors{
		"episodes": validation.Errors{
			"0": validation.Errors{
				"title":         validation.ErrRequired,
				"date_released": validation.ErrRequired,
			},
			"1": validation.Errors{
				"title":         validation.ErrRequired,
				"descriptions":  validation.ErrRequired,
				"date_released": validation.ErrRequired,
				"duration":      validation.ErrRequired,
			},
			"2": validation.Errors{
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
			"3": validation.Errors{
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

	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(req).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Equal(response.Error(
			response.StatusInvalidRequest,
			expError.Error(),
		))
}

func TestHandleEpisodeUpdate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
	method := http.MethodPatch

	seasonNumber := 1

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithPath("ep", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.EpisodeUpdateRequest{}).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// episode not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 999).
		WithPath("ep", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(&dto.EpisodeUpdateRequest{}).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	updates := []struct {
		name string
		req  *dto.EpisodeUpdateRequest
	}{
		{
			name: "u1",
			req:  &dto.EpisodeUpdateRequest{},
		},
		{
			name: "u2",
			req: &dto.EpisodeUpdateRequest{
				Title: null.StringFrom("updated_title"),
			},
		},
		{
			name: "u3",
			req: &dto.EpisodeUpdateRequest{
				Descriptions: null.StringFrom("updated_description"),
			},
		},
		{
			name: "u4",
			req: &dto.EpisodeUpdateRequest{
				DateReleased: null.TimeFrom(
					testutils.Date(2000, 1, 1),
				),
			},
		},
		{
			name: "u5",
			req: &dto.EpisodeUpdateRequest{
				DateReleased: null.TimeFrom(
					testutils.Date(2008, 8, 11),
				),
			},
		},
		{
			name: "u6",
			req: &dto.EpisodeUpdateRequest{
				Title:        null.StringFrom("updated_title"),
				Descriptions: null.StringFrom("updated_description"),
				DateReleased: null.TimeFrom(
					testutils.Date(2000, 1, 1),
				),
				Duration: null.IntFrom(7 * 60),
			},
		},
	}

	for i, u := range updates {
		u := u
		i := i
		// change episode number as underlying db instance is shared between all parallel test cases
		episodeNumber := i + 1
		t.Run(u.name, func(t *testing.T) {
			require := prequire.New(t)

			// insert episode
			err := appInstance.EpisodePut(
				ctx,
				defaults.series.id, seasonNumber, episodeNumber,
				defaults.user.id,
				&dto.EpisodePutRequest{
					Title:        "episode" + strconv.Itoa(i),
					DateReleased: testutils.Date(1900, 3, 14),
				},
			)
			require.NoError(err)

			gotEpisodeBeforeUpdate, err := appInstance.EpisodeGet(
				ctx,
				defaults.series.id,
				seasonNumber,
				episodeNumber,
			)
			require.NoError(err)

			updateTime := time.Now()

			// update
			e.Request(method, path).
				WithPath("id", defaults.series.id).
				WithPath("se", seasonNumber).
				WithPath("ep", episodeNumber).
				WithHeader(echo.HeaderAuthorization, defaults.user.auth).
				WithJSON(u.req).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Equal(response.OK(nil))

			// check updated fields
			gotEpisodeAfterUpdate, err := appInstance.EpisodeGet(
				ctx,
				defaults.series.id,
				seasonNumber,
				episodeNumber,
			)
			require.NoError(err)

			require.GreaterOrEqual(
				gotEpisodeAfterUpdate.ContributedAt,
				updateTime,
			)

			updatedEpisode := &models.Film{}
			updatedEpisode.ID = gotEpisodeAfterUpdate.ID
			updatedEpisode.SeriesID = null.IntFrom(defaults.series.id)
			updatedEpisode.SeasonNumber = null.IntFrom(seasonNumber)
			updatedEpisode.EpisodeNumber = null.IntFrom(episodeNumber)
			if u.req.Title.Valid {
				updatedEpisode.Title = u.req.Title.String
			} else {
				updatedEpisode.Title = gotEpisodeBeforeUpdate.Title
			}
			if u.req.Descriptions.Valid {
				updatedEpisode.Descriptions = u.req.Descriptions
			} else {
				updatedEpisode.Descriptions = gotEpisodeBeforeUpdate.Descriptions
			}
			if u.req.DateReleased.Valid {
				updatedEpisode.DateReleased = u.req.DateReleased.Time
			} else {
				updatedEpisode.DateReleased = gotEpisodeBeforeUpdate.DateReleased
			}
			if u.req.Duration.Valid {
				updatedEpisode.Duration = u.req.Duration
			} else {
				updatedEpisode.Duration = gotEpisodeBeforeUpdate.Duration
			}
			updatedEpisode.Invalidation = null.String{}
			updatedEpisode.ContributedBy = defaults.user.id
			updatedEpisode.ContributedAt = gotEpisodeAfterUpdate.ContributedAt

			testutils.SetTimeLocation(
				&updatedEpisode.DateReleased,
				gotEpisodeAfterUpdate.DateReleased.Location(),
			)

			require.Equal(updatedEpisode, gotEpisodeAfterUpdate)
		})
	}
}

func TestHandleEpisodeUpdate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
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
		req       dto.EpisodeUpdateRequest
		expErrors validation.Errors
	}{
		{
			name: "tc1",
			req: dto.EpisodeUpdateRequest{
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
			req: dto.EpisodeUpdateRequest{
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
			req: dto.EpisodeUpdateRequest{
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
				WithPath("id", defaults.series.id).
				WithPath("se", seasonNumber).
				WithPath("ep", episodeNumber).
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

func TestHandleEpisodeInvalidate(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
	method := http.MethodDelete

	invalidationRequest := &dto.InvalidationRequest{
		Invalidation: "invalidation",
	}

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithPath("ep", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// episode not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 999).
		WithPath("ep", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// invalidate episode
	episodeCreateReq := &dto.EpisodePutRequest{
		Title:        "episode",
		DateReleased: testutils.Date(1900, 3, 14),
	}
	err = appInstance.EpisodePut(
		ctx,
		defaults.series.id, seasonNumber, episodeNumber,
		defaults.user.id,
		episodeCreateReq,
	)
	require.NoError(err)

	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check episode invalidated
	gotInvalidatedEpisode, err := appInstance.EpisodeGet(
		ctx,
		defaults.series.id,
		seasonNumber,
		episodeNumber,
	)
	require.NoError(err)

	testutils.SetTimeLocation(
		&episodeCreateReq.DateReleased,
		gotInvalidatedEpisode.DateReleased.Location(),
	)

	require.Equal(
		&models.Film{
			ID:            gotInvalidatedEpisode.ID,
			Title:         episodeCreateReq.Title,
			Descriptions:  episodeCreateReq.Descriptions,
			DateReleased:  episodeCreateReq.DateReleased,
			Duration:      episodeCreateReq.Duration,
			SeriesID:      null.IntFrom(defaults.series.id),
			SeasonNumber:  null.IntFrom(seasonNumber),
			EpisodeNumber: null.IntFrom(episodeNumber),
			Invalidation:  null.StringFrom(invalidationRequest.Invalidation),
			ContributedBy: defaults.user.id,
			ContributedAt: gotInvalidatedEpisode.ContributedAt,
		},
		gotInvalidatedEpisode,
	)
}

func TestHandleEpisodeInvalidate_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/"
	method := http.MethodDelete

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

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
				WithPath("id", defaults.series.id).
				WithPath("se", seasonNumber).
				WithPath("ep", episodeNumber).
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

func TestHandleEpisodesInvalidateAllBySeason(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/"
	method := http.MethodDelete

	invalidationRequest := &dto.InvalidationRequest{
		Invalidation: "invalidation",
	}

	seasonNumber := 1

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	// invalidate all episodes
	episodePutAllReq := &dto.EpisodesPutAllBySeasonRequest{
		Episodes: []*dto.EpisodePutRequest{
			{
				Title:        "episode1",
				DateReleased: testutils.Date(2000, 1, 2),
			},
			{
				Title:        "episode2",
				DateReleased: testutils.Date(2001, 1, 2),
			},
			{
				Title:        "episode3",
				DateReleased: testutils.Date(2002, 1, 2),
			},
			{
				Title:        "episode4",
				DateReleased: testutils.Date(2003, 1, 2),
			},
			{
				Title:        "episode5",
				DateReleased: testutils.Date(2004, 1, 2),
			},
		},
	}
	err = appInstance.EpisodesPutAllBySeason(
		ctx,
		defaults.series.id, seasonNumber,
		defaults.user.id,
		episodePutAllReq,
	)
	require.NoError(err)

	invalidationTime := time.Now()

	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		WithJSON(invalidationRequest).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.OK(nil))

	// check episodes invalidated
	gotInvalidatedEpisodes, total, err := appInstance.EpisodesGetAllBySeason(
		ctx,
		defaults.series.id,
		seasonNumber,
		0,
		math.MaxInt,
	)
	require.NoError(err)
	require.Equal(len(episodePutAllReq.Episodes), total)

	for i, ep := range episodePutAllReq.Episodes {
		testutils.SetTimeLocation(
			&ep.DateReleased,
			gotInvalidatedEpisodes[i].DateReleased.Location(),
		)

		require.GreaterOrEqual(
			gotInvalidatedEpisodes[i].ContributedAt,
			invalidationTime,
		)

		episodeNumber := i + 1

		require.Equal(
			&models.Film{
				ID:            gotInvalidatedEpisodes[i].ID,
				Title:         ep.Title,
				Descriptions:  ep.Descriptions,
				DateReleased:  ep.DateReleased,
				Duration:      ep.Duration,
				SeriesID:      null.IntFrom(defaults.series.id),
				SeasonNumber:  null.IntFrom(seasonNumber),
				EpisodeNumber: null.IntFrom(episodeNumber),
				Invalidation: null.StringFrom(
					invalidationRequest.Invalidation,
				),
				ContributedBy: defaults.user.id,
				ContributedAt: gotInvalidatedEpisodes[i].ContributedAt,
			},
			gotInvalidatedEpisodes[i],
		)
	}
}

func TestHandleEpisodesInvalidateAllBySeason_ValidateRequest(t *testing.T) {
	require := require.New(t)

	server, _, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	path := "/v1/authorized/series/{id}/season/{se}/episode/"
	method := http.MethodDelete

	seasonNumber := 1

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
				WithPath("id", defaults.series.id).
				WithPath("se", seasonNumber).
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

func TestHandleEpisodeAuditsGetAll(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	server, appInstance, defaults, teardown, err := setup(
		OptEnableDefaultSeries,
	)
	require.NoError(err)
	t.Cleanup(func() { teardown() })

	e := httpexpect.New(t, server.URL)
	path := "/v1/authorized/series/{id}/season/{se}/episode/{ep}/audits/"
	method := http.MethodGet

	var (
		seasonNumber  = 1
		episodeNumber = 1
	)

	// invalid id
	e.Request(method, path).
		WithPath("id", -1).
		WithPath("se", -1).
		WithPath("ep", -1).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		Equal(response.Error(response.StatusInvalidURLParameter))

	// episode not found
	e.Request(method, path).
		WithPath("id", 999).
		WithPath("se", 999).
		WithPath("ep", 999).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		Equal(response.Error(response.StatusNotFound))

	putTime := time.Now()

	// add episode
	episodePutReq := &dto.EpisodePutRequest{
		Title:        "episode",
		DateReleased: testutils.Date(2000, 1, 1),
	}
	err = appInstance.EpisodePut(
		ctx,
		defaults.series.id, seasonNumber, episodeNumber,
		defaults.user.id,
		episodePutReq,
	)
	require.NoError(err)

	gotEpisode, err := appInstance.EpisodeGet(
		ctx,
		defaults.series.id,
		seasonNumber,
		episodeNumber,
	)
	require.NoError(err)

	require.GreaterOrEqual(gotEpisode.ContributedAt, putTime)

	expEpisodeUpdateAudit := &models.FilmsAudit{
		ID:            gotEpisode.ID,
		Title:         episodePutReq.Title,
		Descriptions:  episodePutReq.Descriptions,
		DateReleased:  episodePutReq.DateReleased,
		Duration:      episodePutReq.Duration,
		SeriesID:      null.IntFrom(defaults.series.id),
		SeasonNumber:  null.IntFrom(seasonNumber),
		EpisodeNumber: null.IntFrom(episodeNumber),
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotEpisode.ContributedAt,
	}

	// no audits
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, nil, 0))

	updateTime := time.Now()

	// update the episode
	episodeUpdateReq := &dto.EpisodeUpdateRequest{
		Title:        null.StringFrom("updated title"),
		Descriptions: null.StringFrom("updated descriptions"),
		DateReleased: null.TimeFrom(
			testutils.Date(2005, 11, 14),
		),
		Duration: null.IntFrom(10 * 60),
	}
	err = appInstance.EpisodeUpdate(
		ctx,
		defaults.series.id, seasonNumber, episodeNumber,
		defaults.user.id,
		episodeUpdateReq,
	)
	require.NoError(err)

	gotEpisode, err = appInstance.EpisodeGet(
		ctx,
		defaults.series.id,
		seasonNumber,
		episodeNumber,
	)
	require.NoError(err)

	require.GreaterOrEqual(gotEpisode.ContributedAt, updateTime)

	expEpisodeInvalidationAudit := &models.FilmsAudit{
		ID:            gotEpisode.ID,
		Title:         episodeUpdateReq.Title.String,
		Descriptions:  episodeUpdateReq.Descriptions,
		DateReleased:  episodeUpdateReq.DateReleased.Time,
		Duration:      episodeUpdateReq.Duration,
		SeriesID:      null.IntFrom(defaults.series.id),
		SeasonNumber:  null.IntFrom(seasonNumber),
		EpisodeNumber: null.IntFrom(episodeNumber),
		Invalidation:  null.String{},
		ContributedBy: defaults.user.id,
		ContributedAt: gotEpisode.ContributedAt,
	}

	// get update audits
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.FilmsAudit{expEpisodeUpdateAudit}, 1))

	// invalidate the episode
	err = appInstance.EpisodeInvalidate(
		ctx,
		defaults.series.id, seasonNumber, episodeNumber,
		defaults.user.id,
		&dto.InvalidationRequest{Invalidation: "invalidation"},
	)
	require.NoError(err)

	// get invalidation audits
	e.Request(method, path).
		WithPath("id", defaults.series.id).
		WithPath("se", seasonNumber).
		WithPath("ep", episodeNumber).
		WithHeader(echo.HeaderAuthorization, defaults.user.auth).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Equal(response.Paginated(config.Config.Pagination.Page.MinValue, config.Config.Pagination.PageSize.DefaultValue, []*models.FilmsAudit{expEpisodeInvalidationAudit, expEpisodeUpdateAudit}, 2))
}
