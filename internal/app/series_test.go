package app_test

import (
	"context"
	"errors"
	"testing"
	_ "unsafe"

	"github.com/aria3ppp/watch-server/internal/app"
	"github.com/aria3ppp/watch-server/internal/dto"
	"github.com/aria3ppp/watch-server/internal/models"
	"github.com/aria3ppp/watch-server/internal/repo"
	"github.com/aria3ppp/watch-server/internal/repo/mock_repo"
	"github.com/aria3ppp/watch-server/internal/testutils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestSeriesGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	var (
		id        = 1
		expError  = errors.New("error")
		expSeries = &models.Series{Title: "series"}
	)

	type GetExp struct {
		series *models.Series
		err    error
	}
	type Get struct {
		exp GetExp
	}
	type Exp struct {
		series *models.Series
		err    error
	}
	type TestCase struct {
		name string
		get  Get
		exp  Exp
	}

	testCases := []TestCase{
		{
			name: "error",
			get: Get{
				exp: GetExp{
					series: nil,
					err:    expError,
				},
			},
			exp: Exp{
				series: nil,
				err:    expError,
			},
		},
		{
			name: "not found",
			get: Get{
				exp: GetExp{
					series: nil,
					err:    repo.ErrNoRecord,
				},
			},
			exp: Exp{
				series: nil,
				err:    app.ErrNotFound,
			},
		},
		{
			name: "ok",
			get: Get{
				exp: GetExp{
					series: expSeries,
					err:    nil,
				},
			},
			exp: Exp{
				series: expSeries,
				err:    nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			mockRepo.EXPECT().
				SeriesGet(ctx, id).
				Return(tc.get.exp.series, tc.get.exp.err)

			app := app.NewApplication(mockRepo, nil, nil, nil)

			series, err := app.SeriesGet(ctx, id)
			require.Equal(tc.exp.err, err)
			require.Equal(tc.exp.series, series)
		})
	}
}

func TestSeriesesGetAll(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		offset = 0
		limit  = 50

		expSerieses            = []*models.Series{{Title: "series"}}
		expTotal               = 1000
		expSeriesesGetAllError = errors.New("SeriesesGetAll error")
		expSeriesesCountError  = errors.New("SeriesesCount error")
	)

	type GetAllExp struct {
		serieses []*models.Series
		err      error
	}
	type CountExp struct {
		total int
		err   error
	}
	type TxExp struct {
		err error
	}
	type Tx struct {
		exp TxExp
	}
	type GetAll struct {
		exp GetAllExp
	}
	type Count struct {
		exp CountExp
	}
	type Exp struct {
		serieses []*models.Series
		total    int
		err      error
	}
	type TestCase struct {
		name   string
		tx     Tx
		getAll GetAll
		count  Count
		exp    Exp
	}

	testCases := []TestCase{
		{
			name: "SeriesesGetAll error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesesGetAllError,
				},
			},
			getAll: GetAll{
				exp: GetAllExp{
					serieses: nil,
					err:      expSeriesesGetAllError,
				},
			},
			exp: Exp{
				serieses: nil,
				total:    0,
				err:      expSeriesesGetAllError,
			},
		},

		{
			name: "SeriesesCount error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesesCountError,
				},
			},
			getAll: GetAll{
				exp: GetAllExp{
					serieses: expSerieses,
					err:      nil,
				},
			},
			count: Count{
				exp: CountExp{
					total: 0,
					err:   expSeriesesCountError,
				},
			},
			exp: Exp{
				serieses: nil,
				total:    0,
				err:      expSeriesesCountError,
			},
		},

		{
			name: "ok",
			tx: Tx{
				exp: TxExp{
					err: nil,
				},
			},
			getAll: GetAll{
				exp: GetAllExp{
					serieses: expSerieses,
					err:      nil,
				},
			},
			count: Count{
				exp: CountExp{
					total: expTotal,
					err:   nil,
				},
			},
			exp: Exp{
				serieses: expSerieses,
				total:    expTotal,
				err:      nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			txCall := mockRepo.EXPECT().
				Transaction(ctx, gomock.Any()).
				Do(func(ctx context.Context, fn func(_ context.Context, _ repo.Service) error) {
					fn(ctx, mockRepo)
				}).
				Return(tc.tx.exp.err)

			getAllCall := mockRepo.EXPECT().
				SeriesesGetAll(ctx, offset, limit).
				Return(tc.getAll.exp.serieses, tc.getAll.exp.err).
				After(txCall)

			if tc.getAll.exp.err == nil {
				mockRepo.EXPECT().
					SeriesesCount(ctx).
					Return(tc.count.exp.total, tc.count.exp.err).
					After(getAllCall)
			}

			app := app.NewApplication(mockRepo, nil, nil, nil)

			serieses, total, err := app.SeriesesGetAll(ctx, offset, limit)
			require.Equal(tc.exp.err, err)
			require.Equal(tc.exp.serieses, serieses)
			require.Equal(tc.exp.total, total)
		})
	}
}

func TestSeriesCreate(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		seriesID      = 1
		contributorID = 1
		req           = &dto.SeriesCreateRequest{
			Title: "series",
		}
		expError = errors.New("error")
	)

	type UpsertExp struct {
		err error
	}
	type Upsert struct {
		exp UpsertExp
	}
	type Exp struct {
		seriesID int
		err      error
	}
	type TestCase struct {
		name   string
		upsert Upsert
		exp    Exp
	}

	testCases := []TestCase{
		{
			name: "error",
			upsert: Upsert{
				exp: UpsertExp{
					err: expError,
				},
			},
			exp: Exp{
				seriesID: 0,
				err:      expError,
			},
		},

		{
			name: "ok",
			upsert: Upsert{
				exp: UpsertExp{
					err: nil,
				},
			},
			exp: Exp{
				seriesID: seriesID,
				err:      nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			mockRepo.EXPECT().
				SeriesCreate(
					ctx,
					contributorID, &models.Series{
						Title:        req.Title,
						Descriptions: req.Descriptions,
						DateStarted:  req.DateStarted,
						DateEnded:    req.DateEnded,
					},
				).
				Do(func(_ context.Context, _ int, s *models.Series) {
					s.ID = seriesID
				}).
				Return(tc.upsert.exp.err)

			app := app.NewApplication(mockRepo, nil, nil, nil)

			id, err := app.SeriesCreate(ctx, contributorID, req)
			require.Equal(tc.exp.err, err)
			require.Equal(tc.exp.seriesID, id)
		})
	}
}

//go:linkname seriesUpdateRequestToValidMap github.com/aria3ppp/watch-server/internal/app.seriesUpdateRequestToValidMap
func seriesUpdateRequestToValidMap(
	req *dto.SeriesUpdateRequest,
) map[string]any

func TestSeriesUpdate(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		seriesID      = 1
		contributorID = 1
		req           = &dto.SeriesUpdateRequest{
			Title:        null.StringFrom("series"),
			Descriptions: null.StringFrom("descriptions"),
			DateStarted: null.TimeFrom(
				testutils.Date(1994, 6, 1),
			),
			DateEnded: null.TimeFrom(
				testutils.Date(2003, 9, 13),
			),
		}
		expError = errors.New("error")
	)

	type UpdateExp struct {
		err error
	}
	type Update struct {
		exp UpdateExp
	}
	type Exp struct {
		err error
	}
	type TestCase struct {
		name   string
		update Update
		exp    Exp
	}

	testCases := []TestCase{
		{
			name: "error",
			update: Update{
				exp: UpdateExp{
					err: expError,
				},
			},
			exp: Exp{
				err: expError,
			},
		},

		{
			name: "not found",
			update: Update{
				exp: UpdateExp{
					err: repo.ErrNoRecord,
				},
			},
			exp: Exp{
				err: app.ErrNotFound,
			},
		},

		{
			name: "ok",
			update: Update{
				exp: UpdateExp{
					err: nil,
				},
			},
			exp: Exp{
				err: nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			mockRepo.EXPECT().
				SeriesUpdate(ctx, seriesID, contributorID, seriesUpdateRequestToValidMap(req)).
				Return(tc.update.exp.err)

			app := app.NewApplication(mockRepo, nil, nil, nil)

			err := app.SeriesUpdate(
				ctx,
				seriesID,
				contributorID,
				req,
			)
			require.Equal(tc.exp.err, err)
		})
	}
}

func TestSeriesInvalidate(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		seriesID      = 1
		contributorID = 1
		req           = &dto.InvalidationRequest{
			Invalidation: "invalidation",
		}

		expEpisodesDeleteAllBySeriesError = errors.New(
			"EpisodesDeleteAllBySeries error",
		)
		expSeriesDeleteError = errors.New("SeriesDelete error")
		expNotFound          = app.ErrNotFound
	)

	type TxExp struct {
		err error
	}
	type Tx struct {
		exp TxExp
	}
	type SeriesInvalidateExp struct {
		err error
	}
	type SeriesInvalidate struct {
		exp SeriesInvalidateExp
	}
	type EpisodesDeleteAllExp struct {
		err error
	}
	type EpisodeDeleteAll struct {
		exp EpisodesDeleteAllExp
	}
	type Exp struct {
		err error
	}
	type TestCase struct {
		name                  string
		tx                    Tx
		seriesInvalidate      SeriesInvalidate
		episodesInvalidateAll EpisodeDeleteAll
		exp                   Exp
	}

	testCases := []TestCase{
		{
			name: "SeriesDelete error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesDeleteError,
				},
			},
			seriesInvalidate: SeriesInvalidate{
				exp: SeriesInvalidateExp{
					err: expSeriesDeleteError,
				},
			},
			exp: Exp{
				err: expSeriesDeleteError,
			},
		},

		{
			name: "not found",
			tx: Tx{
				exp: TxExp{
					err: expNotFound,
				},
			},
			seriesInvalidate: SeriesInvalidate{
				exp: SeriesInvalidateExp{
					err: repo.ErrNoRecord,
				},
			},
			exp: Exp{
				err: expNotFound,
			},
		},

		{
			name: "EpisodesInvalidateAllBySeries error",
			tx: Tx{
				exp: TxExp{
					err: expEpisodesDeleteAllBySeriesError,
				},
			},
			seriesInvalidate: SeriesInvalidate{
				exp: SeriesInvalidateExp{
					err: nil,
				},
			},
			episodesInvalidateAll: EpisodeDeleteAll{
				exp: EpisodesDeleteAllExp{
					err: expEpisodesDeleteAllBySeriesError,
				},
			},
			exp: Exp{
				err: expEpisodesDeleteAllBySeriesError,
			},
		},

		{
			name: "ok",
			tx: Tx{
				exp: TxExp{
					err: nil,
				},
			},
			seriesInvalidate: SeriesInvalidate{
				exp: SeriesInvalidateExp{
					err: nil,
				},
			},
			episodesInvalidateAll: EpisodeDeleteAll{
				exp: EpisodesDeleteAllExp{
					err: nil,
				},
			},
			exp: Exp{
				err: nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			txCall := mockRepo.EXPECT().
				Transaction(ctx, gomock.Any()).
				Do(func(ctx context.Context, fn func(_ context.Context, _ repo.Service) error) {
					fn(ctx, mockRepo)
				}).
				Return(tc.tx.exp.err)

			seriessInvalidate := mockRepo.EXPECT().
				SeriesInvalidate(ctx, seriesID, contributorID, req.Invalidation).
				Return(tc.seriesInvalidate.exp.err).
				After(txCall)

			if tc.seriesInvalidate.exp.err == nil {
				mockRepo.EXPECT().
					EpisodesInvalidateAllBySeries(ctx, seriesID, contributorID, req.Invalidation).
					Return(tc.episodesInvalidateAll.exp.err).
					After(seriessInvalidate)
			}

			app := app.NewApplication(mockRepo, nil, nil, nil)

			err := app.SeriesInvalidate(
				ctx,
				seriesID,
				contributorID,
				req,
			)
			require.Equal(tc.exp.err, err)
		})
	}
}

func TestSeriesAuditsGetAll(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		seriesID = 1
		offset   = 0
		limit    = 50

		expSeries = &models.Series{
			ID:    seriesID,
			Title: "series",
		}
		expSeriesGetError          = errors.New("SeriesGet error")
		expSeriesAuditsGetAllError = errors.New("SeriesAuditsGetAll error")
		expSeriesAuditsCountError  = errors.New("SeriesAuditsCount error")
		expAudits                  = []*models.SeriesesAudit{{Title: "audit"}}
		expTotal                   = 100
	)

	type SeriesGetExp struct {
		series *models.Series
		err    error
	}
	type SeriesAuditsCountExp struct {
		total int
		err   error
	}
	type SeriesAuditsGetAllExp struct {
		audits []*models.SeriesesAudit
		err    error
	}
	type TxExp struct {
		err error
	}
	type Tx struct {
		exp TxExp
	}
	type SeriesGet struct {
		exp SeriesGetExp
	}
	type SeriesAuditsGetAll struct {
		exp SeriesAuditsGetAllExp
	}
	type SeriesAuditsCount struct {
		exp SeriesAuditsCountExp
	}
	type Exp struct {
		audits []*models.SeriesesAudit
		total  int
		err    error
	}
	type TestCase struct {
		name               string
		tx                 Tx
		seriesGet          SeriesGet
		seriesAuditsGetAll SeriesAuditsGetAll
		seriesAuditsCount  SeriesAuditsCount
		exp                Exp
	}

	testCases := []TestCase{
		{
			name: "not found",
			tx: Tx{
				exp: TxExp{
					err: app.ErrNotFound,
				},
			},
			seriesGet: SeriesGet{
				exp: SeriesGetExp{
					series: nil,
					err:    repo.ErrNoRecord,
				},
			},
			exp: Exp{
				audits: nil,
				total:  0,
				err:    app.ErrNotFound,
			},
		},

		{
			name: "SeriesGet error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesGetError,
				},
			},
			seriesGet: SeriesGet{
				exp: SeriesGetExp{
					series: nil,
					err:    expSeriesGetError,
				},
			},
			exp: Exp{
				audits: nil,
				total:  0,
				err:    expSeriesGetError,
			},
		},

		{
			name: "SeriesAuditsGetAll error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesAuditsGetAllError,
				},
			},
			seriesGet: SeriesGet{
				exp: SeriesGetExp{
					series: expSeries,
					err:    nil,
				},
			},
			seriesAuditsGetAll: SeriesAuditsGetAll{
				exp: SeriesAuditsGetAllExp{
					audits: nil,
					err:    expSeriesAuditsGetAllError,
				},
			},
			exp: Exp{
				audits: nil,
				total:  0,
				err:    expSeriesAuditsGetAllError,
			},
		},

		{
			name: "SeriesAuditsCount error",
			tx: Tx{
				exp: TxExp{
					err: expSeriesAuditsCountError,
				},
			},
			seriesGet: SeriesGet{
				exp: SeriesGetExp{
					series: expSeries,
					err:    nil,
				},
			},
			seriesAuditsGetAll: SeriesAuditsGetAll{
				exp: SeriesAuditsGetAllExp{
					audits: expAudits,
					err:    nil,
				},
			},
			seriesAuditsCount: SeriesAuditsCount{
				exp: SeriesAuditsCountExp{
					total: 0,
					err:   expSeriesAuditsCountError,
				},
			},
			exp: Exp{
				audits: nil,
				total:  0,
				err:    expSeriesAuditsCountError,
			},
		},

		{
			name: "ok",
			tx: Tx{
				exp: TxExp{
					err: nil,
				},
			},
			seriesGet: SeriesGet{
				exp: SeriesGetExp{
					series: expSeries,
					err:    nil,
				},
			},
			seriesAuditsGetAll: SeriesAuditsGetAll{
				exp: SeriesAuditsGetAllExp{
					audits: expAudits,
					err:    nil,
				},
			},
			seriesAuditsCount: SeriesAuditsCount{
				exp: SeriesAuditsCountExp{
					total: expTotal,
					err:   nil,
				},
			},
			exp: Exp{
				audits: expAudits,
				total:  expTotal,
				err:    nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			controller := gomock.NewController(t)
			mockRepo := mock_repo.NewMockRepositoryTx(controller)

			txCall := mockRepo.EXPECT().
				Transaction(ctx, gomock.Any()).
				Do(func(ctx context.Context, fn func(_ context.Context, _ repo.Service) error) {
					fn(ctx, mockRepo)
				}).
				Return(tc.tx.exp.err)

			seriesGetCall := mockRepo.EXPECT().
				SeriesGet(ctx, seriesID).
				Return(tc.seriesGet.exp.series, tc.seriesGet.exp.err).
				After(txCall)

			if tc.seriesGet.exp.err == nil {
				seriesAuditsGetAllCall := mockRepo.EXPECT().
					SeriesAuditsGetAll(ctx, seriesID, offset, limit).
					Return(tc.seriesAuditsGetAll.exp.audits, tc.seriesAuditsGetAll.exp.err).
					After(seriesGetCall)

				if tc.seriesAuditsGetAll.exp.err == nil {
					mockRepo.EXPECT().
						SeriesAuditsCount(ctx, seriesID).
						Return(tc.seriesAuditsCount.exp.total, tc.seriesAuditsCount.exp.err).
						After(seriesAuditsGetAllCall)
				}
			}

			app := app.NewApplication(mockRepo, nil, nil, nil)

			audits, total, err := app.SeriesAuditsGetAll(
				ctx,
				seriesID,
				offset, limit,
			)
			require.Equal(tc.exp.err, err)
			require.Equal(tc.exp.audits, audits)
			require.Equal(tc.exp.total, total)
		})
	}
}
