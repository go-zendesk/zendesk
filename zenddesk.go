package zendesk

type Builder struct {
	subdomain string

	accessToken string // for Oauth2 auth use

	emailAddress string //for basic auth and api toke use
	password     string // for basic auth use

	apiToken string // for api token auth
}

func New(subdomain string) *Builder {
	return &Builder{
		subdomain: subdomain,
	}
}

func (b *Builder) Oauth2Auth(accessToken string) *Builder {
	b.accessToken =accessToken
	//TODO:
}

func (b *Builder) BasicAuth(emailAddress, password string) *Builder {
	b.emailAddress = emailAddress
	b.password = password
	//TODO:
}

func (b *Builder) ApiTokenAuth(emailAddress, apiToken string) *Builder {
	b.emailAddress = emailAddress
	b.apiToken = apiToken
	//TODO:
}

func (b *Builder) FindAll() {
	//TODO:
}

/**
The allowed options are

per_page
page
sort_order
 */
func (b *Builder) Find() {

}
