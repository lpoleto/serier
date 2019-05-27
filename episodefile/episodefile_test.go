// Copyright (c) 2019 Luiz Poleto
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package episodefile

import (
	"testing"
)

func (f *FileInfo) Equals(to *FileInfo) bool {
	return (f.Extension == to.Extension) &&
		(f.EpisodeNumber == to.EpisodeNumber) &&
		(f.Filename == to.Filename) &&
		(f.Path == to.Path) &&
		(f.SeasonNumber == to.SeasonNumber)
}

func TestParseEpisodeWithFormatXXYY(t *testing.T) {
	expected := &FileInfo{
		Path:          "/tmp/season",
		Filename:      "Fake Episode 01.04 With Title.mkv",
		Extension:     ".mkv",
		SeasonNumber:  1,
		EpisodeNumber: 4}

	actual := parse("/tmp/season", "Fake Episode 01.04 With Title.mkv")

	if !actual.Equals(expected) {
		t.Errorf("Assertion failed.\nExpected: %+v.\nActual:   %+v", expected, actual)
	}
}

func TestParseEpisodeWithFormatSXXEYY(t *testing.T) {
	expected := &FileInfo{
		Path:          "/tmp/season",
		Filename:      "Series S02E18 Title.mp4",
		Extension:     ".mp4",
		SeasonNumber:  2,
		EpisodeNumber: 18}

	actual := parse("/tmp/season", "Series S02E18 Title.mp4")
	if !actual.Equals(expected) {
		t.Errorf("Assertion failed.\nExpected: %+v.\nActual:   %+v", expected, actual)
	}
}

func TestParseEpisodeWithFormatsXXeYY(t *testing.T) {
	expected := &FileInfo{
		Path:          "/tmp/season",
		Filename:      "Series s02e18 Title.mp4",
		Extension:     ".mp4",
		SeasonNumber:  2,
		EpisodeNumber: 18}

	actual := parse("/tmp/season", "Series s02e18 Title.mp4")
	if !actual.Equals(expected) {
		t.Errorf("Assertion failed.\nExpected: %+v.\nActual:   %+v", expected, actual)
	}
}

func TestParseEpisodeWithInvalidFormat(t *testing.T) {
	actual := parse("/tmp/season", "Series Title.mp4")

	if actual != nil {
		t.Errorf("Assertion failed. Expected is not nil: %+v.", actual)
	}
}

func BenchmarkParseEpisode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse("/tmp/season", "Series S02E18 Title.mp4")
	}
}
