// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aria3ppp/watch-server/internal/repo (interfaces: ServiceTx)

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	reflect "reflect"

	models "github.com/aria3ppp/watch-server/internal/models"
	repo "github.com/aria3ppp/watch-server/internal/repo"
	gomock "github.com/golang/mock/gomock"
)

// MockServiceTx is a mock of ServiceTx interface.
type MockServiceTx struct {
	ctrl     *gomock.Controller
	recorder *MockServiceTxMockRecorder
}

// MockServiceTxMockRecorder is the mock recorder for MockServiceTx.
type MockServiceTxMockRecorder struct {
	mock *MockServiceTx
}

// NewMockServiceTx creates a new mock instance.
func NewMockServiceTx(ctrl *gomock.Controller) *MockServiceTx {
	mock := &MockServiceTx{ctrl: ctrl}
	mock.recorder = &MockServiceTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceTx) EXPECT() *MockServiceTxMockRecorder {
	return m.recorder
}

// EpisodeAuditsCount mocks base method.
func (m *MockServiceTx) EpisodeAuditsCount(arg0 context.Context, arg1, arg2, arg3 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodeAuditsCount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodeAuditsCount indicates an expected call of EpisodeAuditsCount.
func (mr *MockServiceTxMockRecorder) EpisodeAuditsCount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodeAuditsCount", reflect.TypeOf((*MockServiceTx)(nil).EpisodeAuditsCount), arg0, arg1, arg2, arg3)
}

// EpisodeAuditsGetAll mocks base method.
func (m *MockServiceTx) EpisodeAuditsGetAll(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 int) ([]*models.FilmsAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodeAuditsGetAll", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*models.FilmsAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodeAuditsGetAll indicates an expected call of EpisodeAuditsGetAll.
func (mr *MockServiceTxMockRecorder) EpisodeAuditsGetAll(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodeAuditsGetAll", reflect.TypeOf((*MockServiceTx)(nil).EpisodeAuditsGetAll), arg0, arg1, arg2, arg3, arg4, arg5)
}

// EpisodeGet mocks base method.
func (m *MockServiceTx) EpisodeGet(arg0 context.Context, arg1, arg2, arg3 int) (*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodeGet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodeGet indicates an expected call of EpisodeGet.
func (mr *MockServiceTxMockRecorder) EpisodeGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodeGet", reflect.TypeOf((*MockServiceTx)(nil).EpisodeGet), arg0, arg1, arg2, arg3)
}

// EpisodeInvalidate mocks base method.
func (m *MockServiceTx) EpisodeInvalidate(arg0 context.Context, arg1, arg2, arg3, arg4 int, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodeInvalidate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// EpisodeInvalidate indicates an expected call of EpisodeInvalidate.
func (mr *MockServiceTxMockRecorder) EpisodeInvalidate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodeInvalidate", reflect.TypeOf((*MockServiceTx)(nil).EpisodeInvalidate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// EpisodePut mocks base method.
func (m *MockServiceTx) EpisodePut(arg0 context.Context, arg1, arg2, arg3, arg4 int, arg5 *models.Film) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodePut", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// EpisodePut indicates an expected call of EpisodePut.
func (mr *MockServiceTxMockRecorder) EpisodePut(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodePut", reflect.TypeOf((*MockServiceTx)(nil).EpisodePut), arg0, arg1, arg2, arg3, arg4, arg5)
}

// EpisodeUpdate mocks base method.
func (m *MockServiceTx) EpisodeUpdate(arg0 context.Context, arg1, arg2, arg3, arg4 int, arg5 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodeUpdate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// EpisodeUpdate indicates an expected call of EpisodeUpdate.
func (mr *MockServiceTxMockRecorder) EpisodeUpdate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodeUpdate", reflect.TypeOf((*MockServiceTx)(nil).EpisodeUpdate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// EpisodesAuditsCountBySeason mocks base method.
func (m *MockServiceTx) EpisodesAuditsCountBySeason(arg0 context.Context, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesAuditsCountBySeason", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesAuditsCountBySeason indicates an expected call of EpisodesAuditsCountBySeason.
func (mr *MockServiceTxMockRecorder) EpisodesAuditsCountBySeason(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesAuditsCountBySeason", reflect.TypeOf((*MockServiceTx)(nil).EpisodesAuditsCountBySeason), arg0, arg1, arg2)
}

// EpisodesAuditsCountBySeries mocks base method.
func (m *MockServiceTx) EpisodesAuditsCountBySeries(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesAuditsCountBySeries", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesAuditsCountBySeries indicates an expected call of EpisodesAuditsCountBySeries.
func (mr *MockServiceTxMockRecorder) EpisodesAuditsCountBySeries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesAuditsCountBySeries", reflect.TypeOf((*MockServiceTx)(nil).EpisodesAuditsCountBySeries), arg0, arg1)
}

// EpisodesAuditsGetAllBySeason mocks base method.
func (m *MockServiceTx) EpisodesAuditsGetAllBySeason(arg0 context.Context, arg1, arg2, arg3, arg4 int) ([]*models.FilmsAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesAuditsGetAllBySeason", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*models.FilmsAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesAuditsGetAllBySeason indicates an expected call of EpisodesAuditsGetAllBySeason.
func (mr *MockServiceTxMockRecorder) EpisodesAuditsGetAllBySeason(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesAuditsGetAllBySeason", reflect.TypeOf((*MockServiceTx)(nil).EpisodesAuditsGetAllBySeason), arg0, arg1, arg2, arg3, arg4)
}

// EpisodesAuditsGetAllBySeries mocks base method.
func (m *MockServiceTx) EpisodesAuditsGetAllBySeries(arg0 context.Context, arg1, arg2, arg3 int) ([]*models.FilmsAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesAuditsGetAllBySeries", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.FilmsAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesAuditsGetAllBySeries indicates an expected call of EpisodesAuditsGetAllBySeries.
func (mr *MockServiceTxMockRecorder) EpisodesAuditsGetAllBySeries(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesAuditsGetAllBySeries", reflect.TypeOf((*MockServiceTx)(nil).EpisodesAuditsGetAllBySeries), arg0, arg1, arg2, arg3)
}

// EpisodesCountBySeason mocks base method.
func (m *MockServiceTx) EpisodesCountBySeason(arg0 context.Context, arg1, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesCountBySeason", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesCountBySeason indicates an expected call of EpisodesCountBySeason.
func (mr *MockServiceTxMockRecorder) EpisodesCountBySeason(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesCountBySeason", reflect.TypeOf((*MockServiceTx)(nil).EpisodesCountBySeason), arg0, arg1, arg2)
}

// EpisodesCountBySeries mocks base method.
func (m *MockServiceTx) EpisodesCountBySeries(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesCountBySeries", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesCountBySeries indicates an expected call of EpisodesCountBySeries.
func (mr *MockServiceTxMockRecorder) EpisodesCountBySeries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesCountBySeries", reflect.TypeOf((*MockServiceTx)(nil).EpisodesCountBySeries), arg0, arg1)
}

// EpisodesGetAllBySeason mocks base method.
func (m *MockServiceTx) EpisodesGetAllBySeason(arg0 context.Context, arg1, arg2, arg3, arg4 int) ([]*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesGetAllBySeason", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesGetAllBySeason indicates an expected call of EpisodesGetAllBySeason.
func (mr *MockServiceTxMockRecorder) EpisodesGetAllBySeason(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesGetAllBySeason", reflect.TypeOf((*MockServiceTx)(nil).EpisodesGetAllBySeason), arg0, arg1, arg2, arg3, arg4)
}

// EpisodesGetAllBySeries mocks base method.
func (m *MockServiceTx) EpisodesGetAllBySeries(arg0 context.Context, arg1, arg2, arg3 int) ([]*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesGetAllBySeries", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpisodesGetAllBySeries indicates an expected call of EpisodesGetAllBySeries.
func (mr *MockServiceTxMockRecorder) EpisodesGetAllBySeries(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesGetAllBySeries", reflect.TypeOf((*MockServiceTx)(nil).EpisodesGetAllBySeries), arg0, arg1, arg2, arg3)
}

// EpisodesInvalidateAllBySeason mocks base method.
func (m *MockServiceTx) EpisodesInvalidateAllBySeason(arg0 context.Context, arg1, arg2, arg3 int, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesInvalidateAllBySeason", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// EpisodesInvalidateAllBySeason indicates an expected call of EpisodesInvalidateAllBySeason.
func (mr *MockServiceTxMockRecorder) EpisodesInvalidateAllBySeason(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesInvalidateAllBySeason", reflect.TypeOf((*MockServiceTx)(nil).EpisodesInvalidateAllBySeason), arg0, arg1, arg2, arg3, arg4)
}

// EpisodesInvalidateAllBySeries mocks base method.
func (m *MockServiceTx) EpisodesInvalidateAllBySeries(arg0 context.Context, arg1, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpisodesInvalidateAllBySeries", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EpisodesInvalidateAllBySeries indicates an expected call of EpisodesInvalidateAllBySeries.
func (mr *MockServiceTxMockRecorder) EpisodesInvalidateAllBySeries(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpisodesInvalidateAllBySeries", reflect.TypeOf((*MockServiceTx)(nil).EpisodesInvalidateAllBySeries), arg0, arg1, arg2, arg3)
}

// MovieAuditsCount mocks base method.
func (m *MockServiceTx) MovieAuditsCount(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieAuditsCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieAuditsCount indicates an expected call of MovieAuditsCount.
func (mr *MockServiceTxMockRecorder) MovieAuditsCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieAuditsCount", reflect.TypeOf((*MockServiceTx)(nil).MovieAuditsCount), arg0, arg1)
}

// MovieAuditsGetAll mocks base method.
func (m *MockServiceTx) MovieAuditsGetAll(arg0 context.Context, arg1, arg2, arg3 int) ([]*models.FilmsAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieAuditsGetAll", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.FilmsAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieAuditsGetAll indicates an expected call of MovieAuditsGetAll.
func (mr *MockServiceTxMockRecorder) MovieAuditsGetAll(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieAuditsGetAll", reflect.TypeOf((*MockServiceTx)(nil).MovieAuditsGetAll), arg0, arg1, arg2, arg3)
}

// MovieCreate mocks base method.
func (m *MockServiceTx) MovieCreate(arg0 context.Context, arg1 int, arg2 *models.Film) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MovieCreate indicates an expected call of MovieCreate.
func (mr *MockServiceTxMockRecorder) MovieCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieCreate", reflect.TypeOf((*MockServiceTx)(nil).MovieCreate), arg0, arg1, arg2)
}

// MovieGet mocks base method.
func (m *MockServiceTx) MovieGet(arg0 context.Context, arg1 int) (*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieGet", arg0, arg1)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MovieGet indicates an expected call of MovieGet.
func (mr *MockServiceTxMockRecorder) MovieGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieGet", reflect.TypeOf((*MockServiceTx)(nil).MovieGet), arg0, arg1)
}

// MovieInvalidate mocks base method.
func (m *MockServiceTx) MovieInvalidate(arg0 context.Context, arg1, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieInvalidate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MovieInvalidate indicates an expected call of MovieInvalidate.
func (mr *MockServiceTxMockRecorder) MovieInvalidate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieInvalidate", reflect.TypeOf((*MockServiceTx)(nil).MovieInvalidate), arg0, arg1, arg2, arg3)
}

// MovieUpdate mocks base method.
func (m *MockServiceTx) MovieUpdate(arg0 context.Context, arg1, arg2 int, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MovieUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MovieUpdate indicates an expected call of MovieUpdate.
func (mr *MockServiceTxMockRecorder) MovieUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MovieUpdate", reflect.TypeOf((*MockServiceTx)(nil).MovieUpdate), arg0, arg1, arg2, arg3)
}

// MoviesCount mocks base method.
func (m *MockServiceTx) MoviesCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoviesCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoviesCount indicates an expected call of MoviesCount.
func (mr *MockServiceTxMockRecorder) MoviesCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoviesCount", reflect.TypeOf((*MockServiceTx)(nil).MoviesCount), arg0)
}

// MoviesGetAll mocks base method.
func (m *MockServiceTx) MoviesGetAll(arg0 context.Context, arg1, arg2 int) ([]*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoviesGetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoviesGetAll indicates an expected call of MoviesGetAll.
func (mr *MockServiceTxMockRecorder) MoviesGetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoviesGetAll", reflect.TypeOf((*MockServiceTx)(nil).MoviesGetAll), arg0, arg1, arg2)
}

// SeriesAuditsCount mocks base method.
func (m *MockServiceTx) SeriesAuditsCount(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesAuditsCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesAuditsCount indicates an expected call of SeriesAuditsCount.
func (mr *MockServiceTxMockRecorder) SeriesAuditsCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesAuditsCount", reflect.TypeOf((*MockServiceTx)(nil).SeriesAuditsCount), arg0, arg1)
}

// SeriesAuditsGetAll mocks base method.
func (m *MockServiceTx) SeriesAuditsGetAll(arg0 context.Context, arg1, arg2, arg3 int) ([]*models.SeriesesAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesAuditsGetAll", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.SeriesesAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesAuditsGetAll indicates an expected call of SeriesAuditsGetAll.
func (mr *MockServiceTxMockRecorder) SeriesAuditsGetAll(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesAuditsGetAll", reflect.TypeOf((*MockServiceTx)(nil).SeriesAuditsGetAll), arg0, arg1, arg2, arg3)
}

// SeriesCreate mocks base method.
func (m *MockServiceTx) SeriesCreate(arg0 context.Context, arg1 int, arg2 *models.Series) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeriesCreate indicates an expected call of SeriesCreate.
func (mr *MockServiceTxMockRecorder) SeriesCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesCreate", reflect.TypeOf((*MockServiceTx)(nil).SeriesCreate), arg0, arg1, arg2)
}

// SeriesGet mocks base method.
func (m *MockServiceTx) SeriesGet(arg0 context.Context, arg1 int) (*models.Series, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesGet", arg0, arg1)
	ret0, _ := ret[0].(*models.Series)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesGet indicates an expected call of SeriesGet.
func (mr *MockServiceTxMockRecorder) SeriesGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesGet", reflect.TypeOf((*MockServiceTx)(nil).SeriesGet), arg0, arg1)
}

// SeriesInvalidate mocks base method.
func (m *MockServiceTx) SeriesInvalidate(arg0 context.Context, arg1, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesInvalidate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeriesInvalidate indicates an expected call of SeriesInvalidate.
func (mr *MockServiceTxMockRecorder) SeriesInvalidate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesInvalidate", reflect.TypeOf((*MockServiceTx)(nil).SeriesInvalidate), arg0, arg1, arg2, arg3)
}

// SeriesUpdate mocks base method.
func (m *MockServiceTx) SeriesUpdate(arg0 context.Context, arg1, arg2 int, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeriesUpdate indicates an expected call of SeriesUpdate.
func (mr *MockServiceTxMockRecorder) SeriesUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesUpdate", reflect.TypeOf((*MockServiceTx)(nil).SeriesUpdate), arg0, arg1, arg2, arg3)
}

// SeriesesCount mocks base method.
func (m *MockServiceTx) SeriesesCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesesCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesesCount indicates an expected call of SeriesesCount.
func (mr *MockServiceTxMockRecorder) SeriesesCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesesCount", reflect.TypeOf((*MockServiceTx)(nil).SeriesesCount), arg0)
}

// SeriesesGetAll mocks base method.
func (m *MockServiceTx) SeriesesGetAll(arg0 context.Context, arg1, arg2 int) ([]*models.Series, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesesGetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Series)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesesGetAll indicates an expected call of SeriesesGetAll.
func (mr *MockServiceTxMockRecorder) SeriesesGetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesesGetAll", reflect.TypeOf((*MockServiceTx)(nil).SeriesesGetAll), arg0, arg1, arg2)
}

// Transaction mocks base method.
func (m *MockServiceTx) Transaction(arg0 context.Context, arg1 func(context.Context, repo.Service) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockServiceTxMockRecorder) Transaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockServiceTx)(nil).Transaction), arg0, arg1)
}

// UserCreate mocks base method.
func (m *MockServiceTx) UserCreate(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserCreate indicates an expected call of UserCreate.
func (mr *MockServiceTxMockRecorder) UserCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockServiceTx)(nil).UserCreate), arg0, arg1)
}

// UserDelete mocks base method.
func (m *MockServiceTx) UserDelete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserDelete indicates an expected call of UserDelete.
func (mr *MockServiceTxMockRecorder) UserDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*MockServiceTx)(nil).UserDelete), arg0, arg1)
}

// UserGet mocks base method.
func (m *MockServiceTx) UserGet(arg0 context.Context, arg1 int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGet", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGet indicates an expected call of UserGet.
func (mr *MockServiceTxMockRecorder) UserGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockServiceTx)(nil).UserGet), arg0, arg1)
}

// UserGetByEmail mocks base method.
func (m *MockServiceTx) UserGetByEmail(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGetByEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGetByEmail indicates an expected call of UserGetByEmail.
func (mr *MockServiceTxMockRecorder) UserGetByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGetByEmail", reflect.TypeOf((*MockServiceTx)(nil).UserGetByEmail), arg0, arg1)
}

// UserUpdate mocks base method.
func (m *MockServiceTx) UserUpdate(arg0 context.Context, arg1 int, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserUpdate indicates an expected call of UserUpdate.
func (mr *MockServiceTxMockRecorder) UserUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*MockServiceTx)(nil).UserUpdate), arg0, arg1, arg2)
}

// UsersCount mocks base method.
func (m *MockServiceTx) UsersCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersCount indicates an expected call of UsersCount.
func (mr *MockServiceTxMockRecorder) UsersCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersCount", reflect.TypeOf((*MockServiceTx)(nil).UsersCount), arg0)
}
