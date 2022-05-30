package service

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"privy-test/infra"
	"testing"
)

var healthServiceMock HealthCheckService

func provideHealthCheckServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	healthServiceMock = NewHealthCheckService(infra.MergeConfig{AppConfig: infra.AppConfig{Environment: "Development"}})

	return func() {}
}

func TestHealthCheck(t *testing.T) {
	t.Run("TestHealthCheck", func(t *testing.T) {
		Convey("TestHealthCheck", t, FailureContinues, func(c C) {
			Convey("TestHealthCheck", func(c C) {
				type (
					args struct {
						ctx context.Context
					}
				)

				testCases := []struct {
					testID   int
					testDesc string
					args     args
				}{
					{
						testID:   1,
						testDesc: "Seccess API Health",
						args: args{
							ctx: context.Background(),
						},
					},
				}
				for _, tc := range testCases {
					close := provideHealthCheckServiceTest(t)
					defer close()
					t.Run(fmt.Sprintf("%d: %s", tc.testID, tc.testDesc), func(t *testing.T) {
						response := healthServiceMock.HealthCheck(tc.args.ctx)
						c.So(response.Data.Environment, ShouldEqual, "Development")
					})
				}
			})
		})
	})
}
