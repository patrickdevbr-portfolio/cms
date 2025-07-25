package page

import "github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"

type PageService interface {
	CreateDraftPage(title string) (*Page, error)
	PublishPage(p *Page) error
	GetPages(filter GetPages) ([]*Page, error)
	GetPageById(id PageID) (*Page, error)
	AddComponent(p *Page, comp *component.Component) error
	EditComponent(p *Page, componentID component.ComponentID, updatedComponent *component.Component) error
}

type GetPages struct {
	Title string
}
