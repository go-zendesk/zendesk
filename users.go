package zendesk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
	"strings"
)

type usersRequest struct {
	*Request
}

type usersOut struct {
	Users []user `json:"users"`

	NextPage     string      `json:"next_page"`
	PreviousPage string `json:"previous_page"`
	Count        int         `json:"count"`
}

type user struct {
	ID                   int           `json:"id"`
	URL                  string        `json:"url"`
	Name                 string        `json:"name"`
	Email                string        `json:"email"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	TimeZone             string        `json:"time_zone"`
	Phone                interface{}   `json:"phone"`
	SharedPhoneNumber    interface{}   `json:"shared_phone_number"`
	Photo                interface{}   `json:"photo"`
	LocaleID             int           `json:"locale_id"`
	Locale               string        `json:"locale"`
	OrganizationID       interface{}   `json:"organization_id"`
	Role                 string        `json:"role"`
	Verified             bool          `json:"verified"`
	ExternalID           interface{}   `json:"external_id"`
	Tags                 []interface{} `json:"tags"`
	Alias                string        `json:"alias"`
	Active               bool          `json:"active"`
	Shared               bool          `json:"shared"`
	SharedAgent          bool          `json:"shared_agent"`
	LastLoginAt          interface{}   `json:"last_login_at"`
	TwoFactorAuthEnabled bool          `json:"two_factor_auth_enabled"`
	Signature            interface{}   `json:"signature"`
	Details              string        `json:"details"`
	Notes                string        `json:"notes"`
	RoleType             interface{}   `json:"role_type"`
	CustomRoleID         interface{}   `json:"custom_role_id"`
	Moderator            bool          `json:"moderator"`
	TicketRestriction    string        `json:"ticket_restriction"`
	OnlyPrivateComments  bool          `json:"only_private_comments"`
	RestrictedAgent      bool          `json:"restricted_agent"`
	Suspended            bool          `json:"suspended"`
	ChatOnly             bool          `json:"chat_only"`
	DefaultGroupID       interface{}   `json:"default_group_id"`
	UserFields           struct {
		Library                  interface{} `json:"library"`
		SystemEmbeddableLastSeen interface{} `json:"system::embeddable_last_seen"`
	} `json:"user_fields"`
}

// GET /api/v2/users.json
func (b *Request) Users() *usersRequest {
	url := fmt.Sprintf("%s/api/v2/users.json", b.subDomain)
	b.Get(url)
	return &usersRequest{Request: b}
}

//GET /api/v2/groups/{id}/users.json
func (b *Request) GroupUsers(id int) *usersRequest {
	url := fmt.Sprintf("%s/api/v2/groups/%d/users.json", b.subDomain, id)
	b.Get(url)
	return &usersRequest{Request: b}
}

//GET /api/v2/organizations/{id}/users.json
func (b *Request) OrganizationUsers(id int) *usersRequest {
	url := fmt.Sprintf("%s/api/v2/organizations/%d/users.json", b.subDomain, id)
	b.Get(url)
	return &usersRequest{Request: b}
}

//find all page
func (u *usersRequest) FindAll() ([]user, []error) {
	var users []user

	u.Request.setPage(1)
	u.Request.setPerPage(100)

	for {
		userOur, errs := u.Find()

		if errs != nil {
			u.Errors = append(u.Errors, errs...)
			break
		}

		users = append(users,userOur.Users...)

		if userOur.NextPage == "" {
			break
		} else {
			query:=strings.Split(userOur.NextPage,"?")[1]
			values, error := url.ParseQuery(query)
			if error != nil {
				u.Errors = append(u.Errors, error)
				break
			}

			page, err := strconv.Atoi(values.Get("page"))
			if err != nil {
				u.Errors = append(u.Errors, error)
				break
			}
			perPage, err := strconv.Atoi(values.Get("per_page"))
			if err != nil {
				u.Errors = append(u.Errors, error)
				break
			}

			u.Request.setPage(page)
			u.Request.setPerPage(perPage)
		}
	}

	return users, u.Errors
}

//fetch one page
func (u *usersRequest) Find() (*usersOut, []error) {
	var out usersOut

	resp, _, errs := u.EndStruct(&out)
	if errs != nil {
		return nil, errs
	}

	if resp.StatusCode == 200 {
		return &out, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	errOut := errorOut{}
	if err := json.Unmarshal(body, &errOut); err != nil {
		u.Errors = append(u.Errors, err)
		return nil, u.Errors
	}

	u.Errors = append(u.Errors, errors.New(errOut.Error))
	return nil, u.Errors

}

func (u *usersRequest) Page(page int) *usersRequest {
	u.Request.setPage(page)
	return u
}

func (u *usersRequest) PerPage(perPage int) *usersRequest {
	u.Request.setPerPage(perPage)
	return u
}

func (u *usersRequest) SortOrder(sortOrder string) *usersRequest {
	u.Request.setSortOrder(sortOrder)
	return u
}
