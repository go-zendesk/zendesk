# zendesk

zendesk api wrapper for golang.

[![Travis branch](https://img.shields.io/travis/go-zendesk/zendesk/master.svg)](https://travis-ci.org/go-zendesk/zendesk)
[![Codecov branch](https://img.shields.io/codecov/c/github/go-zendesk/zendesk/master.svg)](https://codecov.io/gh/go-zendesk/zendesk)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zendesk/zendesk)](https://goreportcard.com/report/github.com/go-zendesk/zendesk)
[![GoDoc](https://godoc.org/github.com/go-zendesk/zendesk?status.svg)](https://godoc.org/github.com/go-zendesk/zendesk)

## Installation

```bash
$ go get -u github.com/go-zendesk/zendesk
```

## API Status

### List Users
- [x] GET /api/v2/users.json
- [x] GET /api/v2/groups/{id}/users.json
- [x] GET /api/v2/organizations/{id}/users.json

### Show User
- [] GET /api/v2/users/{id}.json

### Show Many Users
- [] GET /api/v2/users/show_many.json?ids={ids}

- [] GET /api/v2/users/show_many.json?external_ids={external_ids}

### User related information
- [] GET /api/v2/users/{id}/related.json

### Create User
- [] POST /api/v2/users.json

### Create Or Update Many Users
- [] POST /api/v2/users/create_or_update_many.json

### Merge Self With Another User
- []PUT /api/v2/users/me/merge.json

### Merge End Users
- [] PUT /api/v2/users/{id}/merge.json

### Create Many Users
- [] POST /api/v2/users/create_many.json

### Update User
- [] PUT /api/v2/users/{id}.json

### Update Many Users
- [] PUT /api/v2/users/update_many.json
- [] PUT /api/v2/users/update_many.json?ids={ids}
- [] PUT /api/v2/users/update_many.json?external_ids={external_ids}

### Bulk Deleting Users
- [] DELETE /api/v2/users/destroy_many.json?ids={ids}
- [] DELETE /api/v2/users/destroy_many.json?external_ids={external_ids}

### Delete User
- [] DELETE /api/v2/users/{id}.json

### Search Users
- [] GET /api/v2/users/search.json?query={query}
- [] GET /api/v2/users/search.json?external_id={external_id}

### Autocomplete Users
- [] GET /api/v2/users/autocomplete.json?name={name}

### Request User Create
- [] POST /api/v2/users/request_create.json

### Show the Currently Authenticated User
- [] GET /api/v2/users/me.json

### Set a User's Password
- [] POST /api/v2/users/{user_id}/password.json

### Change Your Password
- [] PUT /api/v2/users/{user_id}/password.json

### Get a list of password requirements
- [] GET /api/v2/users/{user_id}/password/requirements.json