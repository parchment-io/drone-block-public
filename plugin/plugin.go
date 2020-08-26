// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/validator"
)

// New returns a new validator plugin.
func New() (validator.Plugin, error) {
	return new(plugin), nil
}

type plugin struct {
}

func (p *plugin) Validate(ctx context.Context, req *validator.Request) error {
	if req.Repo.Visibility == "public" {
		return validator.ErrBlock
	}

	// a nil error indicates the configuration is valid.
	return nil
}
