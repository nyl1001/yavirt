package model

import (
	"testing"
	"time"

	"github.com/projecteru2/yavirt/internal/errors"
	"github.com/projecteru2/yavirt/idgen"
	"github.com/projecteru2/yavirt/store/mocks"
	"github.com/projecteru2/yavirt/test/assert"
	"github.com/projecteru2/yavirt/test/mock"
)

func init() {
	idgen.Setup(0, time.Now())
}

func TestCreateHost(t *testing.T) {
	var meta, cancel = mocks.Mock()
	defer cancel()
	defer meta.AssertExpectations(t)

	var host = NewHost()

	meta.On("IncrUint32", mock.Anything, mock.Anything).Return(uint32(1), nil).Once()
	meta.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
	assert.NilErr(t, host.Create())
	assert.Equal(t, StatusRunning, host.Status)
}

func TestCreateHostFailedAsNameExists(t *testing.T) {
	var meta, cancel = mocks.Mock()
	defer cancel()
	defer meta.AssertExpectations(t)

	var host = NewHost()

	meta.On("IncrUint32", mock.Anything, mock.Anything).Return(uint32(1), nil).Once()
	meta.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(errors.ErrKeyExists).Once()
	assert.Err(t, host.Create())
}