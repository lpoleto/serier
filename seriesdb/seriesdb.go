// Copyright (c) 2019 Luiz Poleto
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package seriesdb

import (
	"errors"
	"fmt"

	"github.com/garfunkel/go-tvdb"
)

// EpisodeInfo hold resumed information about an episode.
type EpisodeInfo struct {
	SeriesName    string
	Season        uint64
	EpisodeNumber uint64
	Name          string
}

// Holdd a map containing the season number as key and another map with the episode number as key and EpisodeInfo as value.
type episodeMap struct {
	episodes map[uint64][]EpisodeInfo
}

// SeriesDB is a wrapper for TVDB functionality.
type SeriesDB struct {
	SeriesName   string
	seriesList   []*tvdb.Series
	activeSeries *tvdb.Series
	episodeMap   *episodeMap
}

// New returns an instance of SeriesDB.
func New(name string) (s *SeriesDB, e error) {
	seriesList, err := findByName(name)

	if err != nil {
		e = errors.New("Failed to initialize SeriesDB")
	}

	return &SeriesDB{SeriesName: name, seriesList: seriesList}, e
}

// NumSeries returns the number of series found on TVDB.
func (s *SeriesDB) NumSeries() int {
	return len(s.seriesList)
}

// ListSeries prints a list of series found.
func (s *SeriesDB) ListSeries() {
	for i, series := range s.seriesList {
		fmt.Printf("[%d] - %s\n", i, series.SeriesName)
	}
}

// LoadEpisodes fetch additional series data from TVDB
func (s *SeriesDB) LoadEpisodes(i int) error {
	s.setActiveSeries(i)
	err := s.activeSeries.GetDetail()

	if err != nil {
		return err
	}

	// Initialize the map here for now until I find something better
	s.episodeMap = &episodeMap{episodes: make(map[uint64][]EpisodeInfo)}

	for k, v := range s.activeSeries.Seasons {
		episodeInfo := make([]EpisodeInfo, 0)

		for _, episode := range v {
			e := EpisodeInfo{SeriesName: s.activeSeries.SeriesName, Season: k, EpisodeNumber: episode.EpisodeNumber, Name: episode.EpisodeName}
			episodeInfo = append(episodeInfo, e)
		}

		s.episodeMap.episodes[k] = episodeInfo
	}

	return nil
}

// GetEpisodeInfo returns information about an episode on a given season.
func (s *SeriesDB) GetEpisodeInfo(season, episode uint64) EpisodeInfo {
	episodes := s.episodeMap.episodes[season]

	if episodes == nil {
		fmt.Printf("Could not find episodes for season %d\n", season)
	}

	for _, episodeInfo := range episodes {
		if episodeInfo.EpisodeNumber == episode {
			return episodeInfo
		}
	}

	fmt.Printf("Error. Could not find episode number %d\n", episode)

	return EpisodeInfo{}
}

func findByName(name string) ([]*tvdb.Series, error) {
	seriesList, err := tvdb.GetSeries(name)

	return seriesList.Series, err
}

func (s *SeriesDB) setActiveSeries(i int) {
	s.activeSeries = s.seriesList[i]
}
