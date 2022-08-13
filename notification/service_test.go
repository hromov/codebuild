package notification_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hromov/notification/mocks"
	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	ctrl := gomock.NewController(t)

	ts := mocks.NewMockTopicSender(ctrl)
	ts.EXPECT().Send(gomock.Any()).Return(nil)

	require.NoError(t, ts.Send("some msg"))

	ts.EXPECT().Send(gomock.Any()).Return(errors.New("some error"))
	require.EqualError(t, ts.Send("some msg"), "some error")
}
