package demo

import (
	"strings"

	"github.com/fengdotdev/golibs-traits/trait"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/def"
	"github.com/fengdotdev/golibs-utilitycss/sandbox/draf2/helpers"
)

var _ Component = (*ContainerComponent)(nil)

var _ trait.DataTransferObject[ContainerDTO] = (*ContainerComponent)(nil)

type ContainerDTO struct {
	Id      string
	Class   string
	Content string // every children must be render to string for content
}

func Container(children []Component, class string, id ...string) Component {

	s, haveId := helpers.GetFirstString(id)
	haveClass := len(class) > 0
	haveContent := false
	haveChildren := haveContent || len(children) > 0

	return &ContainerComponent{
		id:           helpers.Ternary(haveId, s, helpers.GenerateId()),
		class:        class,
		children:     children,
		haveId:       haveId,
		haveClass:    haveClass,
		haveContent:  haveContent,  // false content is not directly set, but derived from children
		haveChildren: haveChildren, // container may have children
	}

}

type ContainerComponent struct {
	id           string
	class        string
	haveId       bool
	haveClass    bool
	haveContent  bool
	haveChildren bool
	children     []Component
}

// FromDTO implements trait.DataTransferObject.
func (c *ContainerComponent) FromDTO(dto ContainerDTO) error {
	panic("unimplemented")
}

// ToDTO implements trait.DataTransferObject.
func (c *ContainerComponent) ToDTO() (ContainerDTO, error) {
	children, err := c.renderedChildren()
	if err != nil {
		return ContainerDTO{}, err
	}
	dto := ContainerDTO{
		Id:      helpers.Ternary(c.haveId, c.id, helpers.DeterministicId("ContainerComponent", c.id)),
		Class:   helpers.Ternary(c.haveClass, c.class, ""),
		Content: children,
	}

	return dto, nil
}

func (c *ContainerComponent) renderedChildren() (string, error) {
	if c.haveChildren {

		lines := make([]string, len(c.children))
		for _, child := range c.children {
			rendered, err := child.Render()
			if err != nil {
				return "", err
			}
			lines = append(lines, rendered)

		}
		content := strings.Join(lines, def.Newline)
		return content, nil
	}
	return "", nil
}

// Children implements Component.
func (c *ContainerComponent) Children() []Component {
	if c.haveChildren {
		return c.children
	}
	return nil
}

// Class implements Component.
func (c *ContainerComponent) Class() string {
	if c.haveClass {
		return c.class
	}
	return ""
}

// Content implements Component.
func (c *ContainerComponent) Content() string {
	return ""
}

// Render implements Component.
func (c *ContainerComponent) Render() (string, error) {
	dto, err := c.ToDTO()
	if err != nil {
		return "", err
	}

	return helpers.HTMLTemplateGenerator(div, dto)
}

var div = `
<div class="{{.Class}}" id="{{.Id}}">
	{{.Content}}
</div>
`
