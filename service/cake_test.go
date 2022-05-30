package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	mock_logging "privy-test/mock/integration/logging"
	mock_repository "privy-test/mock/repository"
	"privy-test/model"
	"privy-test/param/cake"
	"testing"
)

var (
	cakeServiceMock CakeService
	mockLogger      *mock_logging.MockLogger
	mockCakeRepo    *mock_repository.MockCakeRepository
	errMock         = errors.New("mock error")
)

func provideCakeServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger = mock_logging.NewMockLogger(ctrl)
	mockCakeRepo = mock_repository.NewMockCakeRepository(ctrl)
	cakeServiceMock = NewCakeService(mockCakeRepo, mockLogger)
	return func() {}
}

func TestGetDetail(t *testing.T) {
	t.Run("TestGetDetail", func(t *testing.T) {
		Convey("TestGetDetail", t, FailureContinues, func(c C) {
			Convey("TestGetDetail", func(c C) {
				type (
					args struct {
						ctx    context.Context
						cakeID int64
					}
					MockGetDetailMockCake struct {
						cakeModel *model.Cake
						err       error
					}
				)

				testCases := []struct {
					testID                int
					testDesc              string
					args                  args
					mockGetDetailMockCake MockGetDetailMockCake
					wantErr               bool
				}{
					{
						testID:   1,
						testDesc: "Error Get Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   2,
						testDesc: "Error Not FoundGet Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: sql.ErrNoRows,
						},
						wantErr: true,
					},
					{
						testID:   3,
						testDesc: "Success get detail cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							cakeModel: &model.Cake{
								ID:          sql.NullInt64{Int64: 1, Valid: true},
								Title:       sql.NullString{String: "SS", Valid: true},
								Description: sql.NullString{String: "SS", Valid: true},
								Rating:      sql.NullFloat64{Float64: 1, Valid: true},
								Image:       sql.NullString{String: "SS", Valid: true},
							},
						},
					},
				}

				for _, tc := range testCases {
					close := provideCakeServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						mockCakeRepo.EXPECT().GetDetail(gomock.Any(), gomock.Any()).Return(tc.mockGetDetailMockCake.cakeModel, tc.mockGetDetailMockCake.err)
						mockLogger.EXPECT().ErrorWithContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

						message, err := cakeServiceMock.GetDetail(tc.args.ctx, tc.args.cakeID)
						if tc.wantErr {
							c.So(message, ShouldBeNil)
							c.So(err, ShouldNotBeNil)
						} else {
							c.So(err, ShouldBeNil)
							c.So(message, ShouldNotBeNil)
						}
					})
				}
			})
		})
	})
}

func TestCreate(t *testing.T) {
	t.Run("TestCreate", func(t *testing.T) {
		Convey("TestCreate", t, FailureContinues, func(c C) {
			Convey("TestCreate", func(c C) {
				type (
					args struct {
						ctx     context.Context
						request *cake.CreateUpdateRequest
					}
					MockCreateMockCake struct {
						cakeModel *model.Cake
						err       error
					}
				)

				testCases := []struct {
					testID             int
					testDesc           string
					args               args
					mockCreateMockCake MockCreateMockCake
					wantErr            bool
				}{
					{
						testID:   1,
						testDesc: "Error Create Cake",
						args: args{
							ctx: context.Background(),
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockCreateMockCake: MockCreateMockCake{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   3,
						testDesc: "Success get detail cake",
						args: args{
							ctx: context.Background(),
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockCreateMockCake: MockCreateMockCake{
							cakeModel: &model.Cake{
								ID:          sql.NullInt64{Int64: 1, Valid: true},
								Title:       sql.NullString{String: "SS", Valid: true},
								Description: sql.NullString{String: "SS", Valid: true},
								Rating:      sql.NullFloat64{Float64: 1, Valid: true},
								Image:       sql.NullString{String: "SS", Valid: true},
							},
						},
					},
				}

				for _, tc := range testCases {
					close := provideCakeServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						mockCakeRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(tc.mockCreateMockCake.cakeModel, tc.mockCreateMockCake.err)
						mockLogger.EXPECT().ErrorWithContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

						message, err := cakeServiceMock.Create(tc.args.ctx, tc.args.request)
						if tc.wantErr {
							c.So(message, ShouldBeNil)
							c.So(err, ShouldNotBeNil)
						} else {
							c.So(err, ShouldBeNil)
							c.So(message, ShouldNotBeNil)
						}
					})
				}
			})
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("TestDelete", func(t *testing.T) {
		Convey("TestDelete", t, FailureContinues, func(c C) {
			Convey("TestDelete", func(c C) {
				type (
					args struct {
						ctx    context.Context
						cakeID int64
					}
					MockGetDetailMockCake struct {
						cakeModel *model.Cake
						err       error
					}
					MockDeleteCakeRepo struct {
						err error
					}
				)

				testCases := []struct {
					testID                int
					testDesc              string
					args                  args
					mockGetDetailMockCake MockGetDetailMockCake
					mockDeleteCakeRepo    MockDeleteCakeRepo
					wantErr               bool
				}{
					{
						testID:   1,
						testDesc: "Error Get Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   2,
						testDesc: "Error Not FoundGet Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: sql.ErrNoRows,
						},
						wantErr: true,
					},
					{
						testID:   3,
						testDesc: "Error Delete Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
						mockDeleteCakeRepo: MockDeleteCakeRepo{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   4,
						testDesc: "Success Delete Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
						},
					},
				}

				for _, tc := range testCases {
					close := provideCakeServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						mockCakeRepo.EXPECT().GetDetail(gomock.Any(), gomock.Any()).Return(tc.mockGetDetailMockCake.cakeModel, tc.mockGetDetailMockCake.err)
						mockCakeRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(tc.mockDeleteCakeRepo.err)
						mockLogger.EXPECT().ErrorWithContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

						err := cakeServiceMock.Delete(tc.args.ctx, tc.args.cakeID)
						if tc.wantErr {
							c.So(err, ShouldNotBeNil)
						} else {
							c.So(err, ShouldBeNil)
						}
					})
				}
			})
		})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("TestUpdate", func(t *testing.T) {
		Convey("TestUpdate", t, FailureContinues, func(c C) {
			Convey("TestUpdate", func(c C) {
				type (
					args struct {
						ctx     context.Context
						cakeID  int64
						request *cake.CreateUpdateRequest
					}
					MockGetDetailMockCake struct {
						cakeModel *model.Cake
						err       error
					}
					MockUpdateCakeRepo struct {
						cakeModel *model.Cake
						err       error
					}
				)

				testCases := []struct {
					testID                int
					testDesc              string
					args                  args
					mockGetDetailMockCake MockGetDetailMockCake
					mockUpdateCakeRepo    MockUpdateCakeRepo
					wantErr               bool
				}{
					{
						testID:   1,
						testDesc: "Error Get Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   2,
						testDesc: "Error Not FoundGet Detail Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockGetDetailMockCake: MockGetDetailMockCake{
							err: sql.ErrNoRows,
						},
						wantErr: true,
					},
					{
						testID:   3,
						testDesc: "Error Update Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockUpdateCakeRepo: MockUpdateCakeRepo{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   4,
						testDesc: "Success Delete Cake",
						args: args{
							ctx:    context.Background(),
							cakeID: 1,
							request: &cake.CreateUpdateRequest{
								Title:       "AA",
								Description: "BB",
								Rating:      8,
								Image:       "CC",
							},
						},
						mockUpdateCakeRepo: MockUpdateCakeRepo{
							cakeModel: &model.Cake{
								ID:          sql.NullInt64{Int64: 1, Valid: true},
								Title:       sql.NullString{String: "SS", Valid: true},
								Description: sql.NullString{String: "SS", Valid: true},
								Rating:      sql.NullFloat64{Float64: 1, Valid: true},
								Image:       sql.NullString{String: "SS", Valid: true},
							},
						},
					},
				}

				for _, tc := range testCases {
					close := provideCakeServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						mockCakeRepo.EXPECT().GetDetail(gomock.Any(), gomock.Any()).Return(tc.mockGetDetailMockCake.cakeModel, tc.mockGetDetailMockCake.err)
						mockCakeRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockUpdateCakeRepo.cakeModel, tc.mockUpdateCakeRepo.err)
						mockLogger.EXPECT().ErrorWithContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

						_, err := cakeServiceMock.Update(tc.args.ctx, tc.args.cakeID, tc.args.request)
						if tc.wantErr {
							c.So(err, ShouldNotBeNil)
						} else {
							c.So(err, ShouldBeNil)
						}
					})
				}
			})
		})
	})
}

func TestGetList(t *testing.T) {
	t.Run("TestGetList", func(t *testing.T) {
		Convey("TestGetList", t, FailureContinues, func(c C) {
			Convey("TestGetList", func(c C) {
				type (
					args struct {
						ctx     context.Context
						request *cake.FindAllRequest
					}
					MockGetListCake struct {
						cakeModel []model.Cake
						err       error
					}
				)

				testCases := []struct {
					testID          int
					testDesc        string
					args            args
					mockGetListCake MockGetListCake
					wantErr         bool
				}{
					{
						testID:   1,
						testDesc: "Error Get List Cake",
						args: args{
							ctx:     context.Background(),
							request: &cake.FindAllRequest{MinRating: 0, MaxRating: 10},
						},
						mockGetListCake: MockGetListCake{
							err: errMock,
						},
						wantErr: true,
					},
					{
						testID:   2,
						testDesc: "Error Not Found Get List Cake",
						args: args{
							ctx:     context.Background(),
							request: &cake.FindAllRequest{MinRating: 0, MaxRating: 10},
						},
						mockGetListCake: MockGetListCake{
							err: sql.ErrNoRows,
						},
					},
					{
						testID:   2,
						testDesc: "Error Not Found Get List Cake",
						args: args{
							ctx:     context.Background(),
							request: &cake.FindAllRequest{MinRating: 0, MaxRating: 10},
						},
						mockGetListCake: MockGetListCake{
							cakeModel: []model.Cake{
								{
									ID:          sql.NullInt64{Int64: 1, Valid: true},
									Title:       sql.NullString{String: "SS", Valid: true},
									Description: sql.NullString{String: "SS", Valid: true},
									Rating:      sql.NullFloat64{Float64: 1, Valid: true},
									Image:       sql.NullString{String: "SS", Valid: true},
								},
							},
						},
					},
				}

				for _, tc := range testCases {
					close := provideCakeServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						mockCakeRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(tc.mockGetListCake.cakeModel, tc.mockGetListCake.err)
						mockLogger.EXPECT().ErrorWithContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

						message, err := cakeServiceMock.GetList(tc.args.ctx, tc.args.request)
						if tc.wantErr {
							c.So(message, ShouldBeNil)
							c.So(err, ShouldNotBeNil)
						} else {
							c.So(err, ShouldBeNil)
							c.So(message, ShouldNotBeNil)
						}
					})
				}
			})
		})
	})
}
