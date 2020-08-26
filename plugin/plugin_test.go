// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"testing"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/validator"
)

var noContext = context.Background()

func TestPlugin_Block(t *testing.T) {
	req := &validator.Request{
		Repo: drone.Repo{Visibility: "public"},
		Build: drone.Build{},
	}

	p, err := New()
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	v := p.Validate(noContext, req)

	if got, want := v, validator.ErrBlock; got != want {
		t.Errorf("Expected validate decision to be %s, got %s", want, got)
	}
}

func TestPlugin_Allow(t *testing.T) {
	req := &validator.Request{
		Repo: drone.Repo{Visibility: "private"},
		Build: drone.Build{},
	}

	p, err := New()
	if err != nil {
		t.Errorf("Unexpected error: %q", err)
	}

	v := p.Validate(noContext, req)

	var noError error;

	if got, want := v, noError; got != want {
		t.Errorf("Expected validate decision to be %s, got %s", want, got)
	}
}
