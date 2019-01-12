package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncReceiverCodeBeSuccessful(t *testing.T) {
	funcReceiver := NewFuncReceiverGenerator("f", "*Foo")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(f *Foo)", gen)
}

func TestShouldGeneratingFuncReceiverCodeBeSuccessfulWithEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiverGenerator("", "")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "", gen)
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverNameIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiverGenerator("", "*Foo")
	_, err := funcReceiver.Generate(0)
	assert.EqualError(t, err, errmsg.FuncReceiverNameIsEmptyError().Error())
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverTypeIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiverGenerator("f", "")
	_, err := funcReceiver.Generate(0)
	assert.EqualError(t, err, errmsg.FuncReceiverTypeIsEmptyError().Error())
}
