/*
 *     Copyright 2025 The CNCF ModelPack Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schema_test

import (
	"testing"

	"github.com/modelpack/model-spec/schema"
)

func TestFileSystem(t *testing.T) {
	fs := schema.FileSystem()
	if fs == nil {
		t.Fatal("FileSystem() returned nil")
	}

	// Test that we can open the config-schema.json file
	file, err := fs.Open("config-schema.json")
	if err != nil {
		t.Fatalf("Failed to open config-schema.json: %v", err)
	}
	defer file.Close()

	// Verify the file is not empty
	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to stat config-schema.json: %v", err)
	}
	if stat.Size() == 0 {
		t.Error("config-schema.json is empty")
	}
}