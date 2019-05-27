// Copyright (c) 2019 Luiz Poleto
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package episodefile

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	relist = [2]*regexp.Regexp{
		regexp.MustCompile(`.*[Ss](\d{2})[Ee](\d{2})`),
		regexp.MustCompile(`.*(\d{2})\.(\d{2})`)}
)

// FileInfo stores information about a file and its parsed information: season and episode number.
type FileInfo struct {
	Path          string
	Filename      string
	Extension     string
	SeasonNumber  uint64
	EpisodeNumber uint64
	NewName       string
}

// Parse an episode name and tries to extract the season and episode numbers.
func parse(path, name string) *FileInfo {
	var seasonNumber, episodeNumber uint64

	ext := filepath.Ext(name)

	for _, re := range relist {
		result := re.FindStringSubmatch(name)

		if result == nil || result[0] == "" {
			continue
		}

		seasonNumber, _ = strconv.ParseUint(result[1], 10, 64)
		episodeNumber, _ = strconv.ParseUint(result[2], 10, 64)

		return &FileInfo{Path: path, Filename: name, Extension: ext, SeasonNumber: seasonNumber, EpisodeNumber: episodeNumber}
	}

	// If we get here it means we couldn't parse the file
	return nil
}

// ListFiles lists all files in the given path.
func ListFiles(path string) ([]*FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	fileInfo := make([]*FileInfo, 0)

	for _, file := range files {
		if !file.IsDir() {
			info := parse(path, file.Name())

			// Ignore the files that failed to parse for now.
			if info != nil {
				fileInfo = append(fileInfo, info)
			}
		}
	}

	return fileInfo, nil
}

// FileExists checks if the file specified in filename exists.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// RenameFile renames a file from old to new.
func RenameFile(oldpath, newpath string) {
	os.Rename(oldpath, newpath)
}
