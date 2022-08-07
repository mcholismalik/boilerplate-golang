package dto

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SearchGetResponseDoc struct {
	Datas          []interface{} `json:"data"`
	PaginationInfo abstraction.PaginationInfo
}

type ByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type cUUID struct {
	uuid.UUID
}

func CUUIDInit(value uuid.UUID) cUUID {
	return cUUID{value}
}

func (c *cUUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		c.UUID = uuid.Nil
		return nil
	}
	uuid, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	c.UUID = uuid
	return nil
}

func (c *cUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.UUID.String())
}

func (c *cUUID) ToUUID() uuid.UUID {
	return c.UUID
}

type cTime struct {
	time.Time
}

func (c *cTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		c.Time = time.Time{}
		return nil
	}
	var format = "2006-01-02"
	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c *cTime) ToTime() time.Time {
	return c.Time
}

func (r *ByIDRequest) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}

	var err error
	r.ID, err = uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id is not valid")
	}

	return nil
}
