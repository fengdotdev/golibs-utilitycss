package demo

import (
	"github.com/fengdotdev/golibs-traits/trait"
	"github.com/fengdotdev/golibs-utilitycss/v0/helpers"
)

var _ trait.DataTransferObject[TextDTO] = (*TextComponent)(nil)
var _ Component = (*TextComponent)(nil)

func Text(text string, class string, id ...string) Component {
	s, haveId := helpers.GetFirstString(id)
	haveClass := len(class) > 0
	haveContent := len(text) > 0
	return &TextComponent{
		id:           helpers.Ternary(haveId, s, helpers.GenerateId()),
		class:        class,
		content:      text,
		haveId:       haveId,
		haveClass:    haveClass,
		haveContent:  haveContent,
		haveChildren: false, // Text component does not have children
	}
}

var txt = `
<div class="{{.Class}}" id="{{.Id}}">
	{{.Content}}
</div>
`

type TextDTO struct {
	Id      string
	Class   string
	Content string
}

type TextComponent struct {
	id           string
	content      string
	class        string
	haveId       bool
	haveClass    bool
	haveContent  bool
	haveChildren bool
}

// FromDTO implements trait.DataTransferObject.
func (t *TextComponent) FromDTO(dto TextDTO) error {
	panic("unimplemented")
}

// ToDTO implements trait.DataTransferObject.
func (t *TextComponent) ToDTO() (TextDTO, error) {
	dto := TextDTO{
		Id:      helpers.Ternary(t.haveId, t.id, helpers.DeterministicId("TextComponent", t.id)),
		Class:   helpers.Ternary(t.haveClass, t.class, ""),
		Content: helpers.Ternary(t.haveContent, t.content, ""),
	}
	return dto, nil
}

// Class implements Component.
func (t *TextComponent) Class() string {
	return t.class
}

// Content implements Component.
func (t *TextComponent) Content() string {
	return t.content
}

// Children implements Component.
func (t *TextComponent) Children() []Component {
	return nil
}

// Render implements Component.
func (t *TextComponent) Render() (string, error) {

	dto, err := t.ToDTO()
	if err != nil {
		return "", err
	}
	result, err := helpers.TemplateGenerator(txt, dto, true)

	if err != nil {
		return "", err
	}
	return result, nil

}
