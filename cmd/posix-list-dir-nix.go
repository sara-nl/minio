// +build linux darwin freebsd netbsd openbsd

/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"io/ioutil"
	"os"
	"path"
)

// Return all the entries at the directory dirPath.
func readDir(dirPath string) (entries []string, err error) {
	err = nil
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filename := file.Name()
		fullpath := path.Join(dirPath, filename)
		info, err := os.Lstat(fullpath)
		if err != nil {
			// no access to this file; skip it
			continue
		}

		mode := info.Mode()
		if mode.IsRegular() {
			entries = append(entries, filename)
		} else if mode.IsDir() {
			entries = append(entries, filename+slashSeparator)
		}

		// else it might be a symlink or a special file
		// those are skipped
	}

	return
}

// EOB
