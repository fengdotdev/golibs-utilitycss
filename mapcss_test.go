package golibsutilitycss_test

import (
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	utilitycss "github.com/fengdotdev/golibs-utilitycss"
)

func TestMapCssInterface(t *testing.T) {

	var mycss utilitycss.MapCSSInterface = utilitycss.NewMapCSS()

	assert.True(t, mycss != nil && mycss.GetCSS() == "")
}

func TestMapCssAdd(t *testing.T) {

	var mycss utilitycss.MapCSSInterface = utilitycss.NewMapCSS()

	selector := utilitycss.NewSelector("body")
	declaration := utilitycss.NewDeclaration("background-color", "red")

	err := mycss.Add(selector, declaration)
	assert.True(t, err == nil)

	assert.True(t, mycss.GetCSS() != "")
	assert.True(t, mycss.GetContent() != nil)
	result, err := mycss.Get(selector)
	assert.NoError(t, err)
	assert.True(t, result != nil)
	assert.True(t, len(result) == 1)
	assert.True(t, result[0].GetProperty() == "background-color")
	assert.True(t, result[0].GetValue() == "red")
	assert.Equal(t, "body{background-color:red;}", mycss.GetCSS())

	var mycss2 utilitycss.MapCSSInterface = utilitycss.NewMapCSS()

	selector2 := utilitycss.NewSelector("body")
	declaration1 := utilitycss.NewDeclaration("background-color", "blue")
	declaration2 := utilitycss.NewDeclaration("color", "white")
	declaration3 := utilitycss.NewDeclaration("font-size", "12px")

	err = mycss2.Add(selector2, declaration1, declaration2, declaration3)
	assert.True(t, err == nil)
	assert.True(t, mycss.GetCSS() != "")

	result2, err := mycss2.Get(selector)
	assert.NoError(t, err)
	assert.True(t, result != nil)
	assert.TrueWithMessage(t, len(result2) == 3, fmt.Sprintf("len(result) = %d", len(result2)))
	assert.True(t, result2[0].GetProperty() == "background-color")
	assert.True(t, result2[0].GetValue() == "blue")
	assert.True(t, result2[1].GetProperty() == "color")
	assert.True(t, result2[1].GetValue() == "white")
	assert.True(t, result2[2].GetProperty() == "font-size")
	assert.True(t, result2[2].GetValue() == "12px")
	assert.Equal(t, "body{background-color:blue;color:white;font-size:12px;}", mycss2.GetCSS())
}

func TestMapCssAddOverride(t *testing.T) {

	var mycss utilitycss.MapCSSInterface = utilitycss.NewMapCSS()


	class:= "bg-red-500 text-white lg"

	selector := utilitycss.NewSelector("body")
	declaration := utilitycss.NewDeclaration("background-color", "red")

	mycss.Add(selector, declaration)

	declaration = utilitycss.NewDeclaration("background-color", "blue")

	mycss.AddOverride(selector, declaration)
}
