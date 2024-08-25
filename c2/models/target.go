package models

import (
	"errors"
	"strings"
)

type Target struct {
	Parent string `json:"parentId"`
	Name   string `json:"name"`
	Value  string `json:"value" gorm:"primaryKey"`
	Type   string `json:"type"`
}

func (t *Target) Validate() error {
	if t.Value == "" {
		return errors.New("target.Value cannot be empty")
	}
	if strings.HasPrefix(t.Value, "http://") || strings.HasPrefix(t.Value, "https://") {
		t.Type = "URL"
	}
	// TODO: Add extra analysis with AI.
	if t.Name == "" {
		t.Name = t.Value
	}
	return nil
}
