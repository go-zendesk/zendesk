package zendesk

import (
	"github.com/parnurzeal/gorequest"
	"strconv"
)

type Request struct {
	subDomain string

	pagination
	*gorequest.SuperAgent

	Errors []error
}

type errorOut struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

func New(subDomain string) *Request {
	return &Request{
		subDomain:  subDomain,
		SuperAgent: gorequest.New(),
	}
}

func (b *Request) Debug(enable bool) *Request {
	b.SuperAgent.SetDebug(enable)
	return b
}

func (b *Request) Oauth2Auth(accessToken string) *Request {
	b.Header.Add("Authorization", "Bearer "+accessToken)
	return b
}

func (b *Request) BasicAuth(emailAddress, password string) *Request {
	b.SetBasicAuth(emailAddress, password)
	return b
}

func (b *Request) ApiTokenAuth(emailAddress, apiToken string) *Request {
	b.SetBasicAuth(emailAddress+"/token", apiToken)
	return b
}

func (b *Request) setPage(page int) {
	b.page = page
	b.Param("page", strconv.Itoa(page))
}

func (b *Request) setPerPage(perPage int) {
	b.perPage = perPage
	b.Param("per_page", strconv.Itoa(perPage))
}

func (b *Request) setSortOrder(sortOrder string) {
	b.sortOrder = sortOrder
	b.Param("sort_order", sortOrder)
}
