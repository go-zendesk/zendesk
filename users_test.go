package zendesk

import (
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

var testUrl = "https://my-api"

func TestUsersRequest_Find(t *testing.T) {
	gorequest.DisableTransportSwap = true

	defer gock.Off()

	gock.New(testUrl).
		Get("/api/v2/users.json").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  363031767271,
					URL: "https://my-api/api/v2/users/363065702632.json",
				},
			}, Count: 1,
		},
		)

	userOut, err := New(testUrl).Users().Page(1).PerPage(10).Find()

	assert.Nil(t, err)

	assert.Exactly(t, &usersOut{
		Users: []user{
			{
				ID:  363031767271,
				URL: "https://my-api/api/v2/users/363065702632.json",
			},
		}, Count: 1,
	}, userOut)

}

func TestUsersRequest_FindAll(t *testing.T) {
	gorequest.DisableTransportSwap = true

	defer gock.Off()

	gock.New(testUrl).
		Get("/api/v2/users.json").
		MatchParam("page", "1").
		MatchParam("per_page", "100").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  1,
					URL: "https://my-api/api/v2/users/1.json",
				},
			}, Count: 100,
			NextPage: "https://my-api/api/v2/users.json?page=2&per_page=100",
		},
		)

	gock.New(testUrl).
		Get("/api/v2/users.json").
		MatchParam("page", "2").
		MatchParam("per_page", "100").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  2,
					URL: "https://my-api/api/v2/users/2.json",
				},
			}, Count: 100,
		},
		)

	users, err := New(testUrl).Users().FindAll()

	assert.Nil(t, err)

	assert.Exactly(t, []user{
		{
			ID:  1,
			URL: "https://my-api/api/v2/users/1.json",
		},
		{
			ID:  2,
			URL: "https://my-api/api/v2/users/2.json",
		},
	}, users)

}

func TestRequest_GroupUsers(t *testing.T) {
	gorequest.DisableTransportSwap = true

	defer gock.Off()

	gock.New(testUrl).
		Get("/api/v2/groups/1/users.json").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  363031767271,
					URL: "https://my-api/api/v2/users/363065702632.json",
				},
			}, Count: 1,
		},
		)

	userOut, err := New(testUrl).GroupUsers(1).Find()

	assert.Nil(t, err)

	assert.Exactly(t, &usersOut{
		Users: []user{
			{
				ID:  363031767271,
				URL: "https://my-api/api/v2/users/363065702632.json",
			},
		}, Count: 1,
	}, userOut)
}

func TestRequest_OrganizationUsers(t *testing.T) {
	gorequest.DisableTransportSwap = true

	defer gock.Off()

	gock.New(testUrl).
		Get("/api/v2/organizations/1/users.json").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  363031767271,
					URL: "https://my-api/api/v2/users/363065702632.json",
				},
			}, Count: 1,
		},
		)

	userOut, err := New(testUrl).OrganizationUsers(1).Find()

	assert.Nil(t, err)

	assert.Exactly(t, &usersOut{
		Users: []user{
			{
				ID:  363031767271,
				URL: "https://my-api/api/v2/users/363065702632.json",
			},
		}, Count: 1,
	}, userOut)
}

func TestRequest_Users(t *testing.T) {
	gorequest.DisableTransportSwap = true

	defer gock.Off()

	gock.New(testUrl).
		Get("/api/v2/users/123.json").
		Reply(200).
		JSON(&usersOut{
			Users: []user{
				{
					ID:  123,
					URL: "https://my-api/api/v2/users/123.json",
				},
			},
		},
		)

	userOut, err := New(testUrl).User(123).Find()

	assert.Nil(t, err)

	assert.Exactly(t, &usersOut{
		Users: []user{
			{
				ID:  123,
				URL: "https://my-api/api/v2/users/123.json",
			},
		},
	}, userOut)
}
