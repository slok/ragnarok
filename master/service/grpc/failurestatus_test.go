package grpc_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	"github.com/slok/ragnarok/api"
	chaosv1 "github.com/slok/ragnarok/api/chaos/v1"
	"github.com/slok/ragnarok/apimachinery/serializer"
	"github.com/slok/ragnarok/clock"
	pbfs "github.com/slok/ragnarok/grpc/failurestatus"
	"github.com/slok/ragnarok/log"
	"github.com/slok/ragnarok/master/service/grpc"
	mclock "github.com/slok/ragnarok/mocks/clock"
	mpbfs "github.com/slok/ragnarok/mocks/grpc/failurestatus"
	mservice "github.com/slok/ragnarok/mocks/master/service"
	testpb "github.com/slok/ragnarok/test/pb"
)

func TestFailureStatusGRPCGetFailureOK(t *testing.T) {
	assert := assert.New(t)

	stubF := &chaosv1.Failure{
		Metadata: api.ObjectMeta{
			ID: "test1",
		},
		Spec: chaosv1.FailureSpec{},
		Status: chaosv1.FailureStatus{
			CurrentState:  chaosv1.EnabledFailureState,
			ExpectedState: chaosv1.DisabledFailureState,
		},
	}
	expF := testpb.CreatePBFailure(stubF, t)

	// Create mocks.
	mfss := &mservice.FailureStatusService{}
	mfss.On("GetFailure", mock.Anything).Once().Return(stubF, nil)

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(0, serializer.PBSerializerDefault, mfss, clock.Base(), log.Dummy)

	// Get the failure and check.
	fID := &pbfs.FailureId{Id: stubF.Metadata.ID}
	gotF, err := fs.GetFailure(context.Background(), fID)
	if assert.NoError(err) {
		assert.Equal(expF, gotF)
	}
}

func TestFailureStatusGRPCGetFailureError(t *testing.T) {
	assert := assert.New(t)

	// Create mocks.
	mfss := &mservice.FailureStatusService{}
	mfss.On("GetFailure", mock.AnythingOfType("string")).Once().Return(nil, errors.New("wanted error"))

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(0, serializer.PBSerializerDefault, mfss, clock.Base(), log.Dummy)

	// Get the failure and check.
	fID := &pbfs.FailureId{Id: "test"}
	_, err := fs.GetFailure(context.Background(), fID)
	assert.Error(err)
}

func TestFailureStatusGRPCGetFailureCtxCanceled(t *testing.T) {
	assert := assert.New(t)

	// Create mocks.
	mfss := &mservice.FailureStatusService{}
	mfss.On("GetFailure", mock.AnythingOfType("string")).Once().Return(nil, nil)

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(0, serializer.PBSerializerDefault, mfss, clock.Base(), log.Dummy)

	// Cancel context.
	ctx, cncl := context.WithCancel(context.Background())
	cncl()

	// Get the failure and check.
	fID := &pbfs.FailureId{Id: "test"}
	_, err := fs.GetFailure(ctx, fID)
	assert.Error(err)
}

func TestFailureStatusGRPCFailureStateListOK(t *testing.T) {
	assert := assert.New(t)

	nodeID := &pbfs.NodeId{Id: "test1"}
	times := 5
	fss := []*chaosv1.Failure{
		&chaosv1.Failure{
			Metadata: api.ObjectMeta{
				ID: "test11",
			},
			Status: chaosv1.FailureStatus{
				CurrentState:  chaosv1.EnabledFailureState,
				ExpectedState: chaosv1.EnabledFailureState,
			},
		},
		&chaosv1.Failure{
			Metadata: api.ObjectMeta{
				ID: "test12",
			},
			Status: chaosv1.FailureStatus{
				CurrentState:  chaosv1.EnabledFailureState,
				ExpectedState: chaosv1.EnabledFailureState,
			},
		},
		&chaosv1.Failure{
			Metadata: api.ObjectMeta{
				ID: "test13",
			},
			Status: chaosv1.FailureStatus{
				CurrentState:  chaosv1.EnabledFailureState,
				ExpectedState: chaosv1.DisabledFailureState,
			},
		},
	}

	// Create mocks.
	mfss := &mservice.FailureStatusService{}
	mfss.On("GetNodeFailures", nodeID.GetId()).Times(times).Return(fss)

	mstream := &mpbfs.FailureStatus_FailureStateListServer{}
	mstream.On("Context").Return(context.Background())
	mstream.On("Send", mock.Anything).Return(nil)

	mtime := &mclock.Clock{}
	tC := make(chan time.Time)
	mtime.On("NewTicker", mock.Anything).Once().Return(&time.Ticker{C: tC})

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(1, serializer.PBSerializerDefault, mfss, mtime, log.Dummy)

	// Simulate the ticker that triggers the update.
	go func() {
		for i := 0; i < times; i++ {
			tC <- time.Now()
		}
		close(tC)
	}()

	// Run the failure state refresh in background.
	err := fs.FailureStateList(nodeID, mstream)
	assert.NoError(err)

	time.Sleep(5 * time.Millisecond) // Used to wait for the final calls and have a real assert.
	mfss.AssertExpectations(t)
	mstream.AssertExpectations(t)
}

func TestFailureStatusGRPCFailureStateListContextClosed(t *testing.T) {
	assert := assert.New(t)

	nodeID := &pbfs.NodeId{Id: "test1"}

	// Create mocks.
	mfss := &mservice.FailureStatusService{}

	mstream := &mpbfs.FailureStatus_FailureStateListServer{}
	ctx, clfn := context.WithCancel(context.Background())
	clfn() // Cancel the context.
	mstream.On("Context").Return(ctx)

	mtime := &mclock.Clock{}
	tC := make(chan time.Time)
	mtime.On("NewTicker", mock.Anything).Once().Return(&time.Ticker{C: tC})

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(1, serializer.PBSerializerDefault, mfss, mtime, log.Dummy)

	// Trigger one round on the update loop.
	go func() {
		tC <- time.Now()
	}()

	// Check.
	err := fs.FailureStateList(nodeID, mstream)
	if assert.NoError(err) {
		mfss.AssertExpectations(t)
		mstream.AssertExpectations(t)
	}

	close(tC)
}

func TestFailureStatusGRPCFailureStateListErr(t *testing.T) {
	assert := assert.New(t)

	nodeID := &pbfs.NodeId{Id: "test1"}

	// Create mocks.
	mfss := &mservice.FailureStatusService{}
	mfss.On("GetNodeFailures", nodeID.GetId()).Once().Return(nil)

	mstream := &mpbfs.FailureStatus_FailureStateListServer{}
	mstream.On("Context").Return(context.Background())
	mstream.On("Send", mock.Anything).Return(errors.New("wanted error"))

	mtime := &mclock.Clock{}
	tC := make(chan time.Time)
	mtime.On("NewTicker", mock.Anything).Once().Return(&time.Ticker{C: tC})

	// Create the GRPC service.
	fs := grpc.NewFailureStatus(1, serializer.PBSerializerDefault, mfss, mtime, log.Dummy)

	// Trigger one round on the update loop.
	go func() {
		tC <- time.Now()
	}()

	// Check.
	err := fs.FailureStateList(nodeID, mstream)
	if assert.Error(err) {
		mfss.AssertExpectations(t)
		mstream.AssertExpectations(t)
	}

	close(tC)
}
