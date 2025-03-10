// Copyright 2025 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	proto "github.com/bucketeer-io/bucketeer/proto/tag"
)

func TestNewTag(t *testing.T) {
	t.Parallel()
	tag, err := NewTag("tag", "envID", proto.Tag_FEATURE_FLAG)
	assert.NoError(t, err)
	assert.IsType(t, &Tag{}, tag)
	assert.NotEmpty(t, tag.Id)
	assert.Equal(t, "tag", tag.Name)
	assert.True(t, tag.CreatedAt > 0)
	assert.True(t, tag.UpdatedAt > 0)
	assert.Equal(t, proto.Tag_FEATURE_FLAG, tag.EntityType)
	assert.Equal(t, "envID", tag.EnvironmentId)
}
