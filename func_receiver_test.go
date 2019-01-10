package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncReceiverCodeBeSuccessful(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "*Foo")
	gen, err := funcReceiver.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, "(f *Foo)", gen)
}

func TestShouldGeneratingFuncReceiverCodeBeSuccessfulWithEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("", "")
	gen, err := funcReceiver.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, "", gen)
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverNameIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("", "*Foo")
	_, err := funcReceiver.GenerateCode()
	assert.EqualError(t, err, errmsg.FuncReceiverNameIsEmptyError().Error())
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverTypeIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "")
	_, err := funcReceiver.GenerateCode()
	assert.EqualError(t, err, errmsg.FuncReceiverTypeIsEmptyError().Error())
}
