package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncReceiverCodeBeSuccessful(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "*Foo")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(f *Foo)", gen)
}

func TestShouldGeneratingFuncReceiverCodeBeSuccessfulWithEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("", "")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "", gen)
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverNameIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("", "*Foo")
	_, err := funcReceiver.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncReceiverNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGeneratingFuncReceiverRaisesErrorWhenFuncReceiverTypeIsEmpty(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "")
	_, err := funcReceiver.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncReceiverTypeIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGeneratingFuncReceiverCodeWithGenericsTypeParamNameSuccessfully(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "*Foo", "T")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(f *Foo[T])", gen)
}

func TestShouldGeneratingFuncReceiverCodeWithGenericsTypeParamNamesSuccessfully(t *testing.T) {
	funcReceiver := NewFuncReceiver("f", "*Foo", "T", "U")
	gen, err := funcReceiver.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(f *Foo[T, U])", gen)
}
