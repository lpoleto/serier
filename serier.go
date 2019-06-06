// Copyright (c) 2019 Luiz Poleto
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lpoleto/serier/episodefile"
	"github.com/lpoleto/serier/seriesdb"
)

const finalFilename = "%s - S%02dE%02d - %s%s"

// FileRename contains a mapping of a file to be renamed.
type FileRename struct {
	Old string
	New string
}

// ReadSeries orchestrates the renaming of the files.
func ReadSeries(name, dir string) error {
	seriesDB, err := seriesdb.New(name)

	if err != nil {
		return err
	}

	if seriesDB.NumSeries() <= 0 {
		fmt.Printf("No series found that match the search term \"%s\"\n", seriesDB.SeriesName)
		os.Exit(0)
	}

	fmt.Printf("Found %d serie(s) that match the search term %s\n\n", seriesDB.NumSeries(), seriesDB.SeriesName)
	seriesDB.ListSeries()

	fmt.Printf("\nType the number from the list above for the series you want and press ENTER. Press CTRL + C to abort.\n")
	var selection int
	_, err = fmt.Scanf("%d", &selection)

	if (err != nil) || (selection < 0) || (selection > seriesDB.NumSeries()-1) {
		fmt.Println("Invalid selection")
		os.Exit(1)
	}

	err = seriesDB.LoadEpisodes(selection)

	if err != nil {
		fmt.Printf("Failed to load series data.\n")
		os.Exit(-1)
	}

	files, err := episodefile.ListFiles(dir)

	if err != nil {
		return err
	}

	renameMap := make([]FileRename, 0)

	for _, file := range files {
		episodeInfo := seriesDB.GetEpisodeInfo(file.SeasonNumber, file.EpisodeNumber)
		formattedFileName := fmt.Sprintf(finalFilename, episodeInfo.SeriesName, episodeInfo.Season, episodeInfo.EpisodeNumber, episodeInfo.Name, file.Extension)

		oldPath := filepath.Join(file.Path, file.Filename)
		newPath := filepath.Join(file.Path, formattedFileName)
		toRename := FileRename{Old: oldPath, New: newPath}
		renameMap = append(renameMap, toRename)
	}

	if len(renameMap) <= 0 {
		fmt.Println("No files eligible to rename. Aborting.")
		os.Exit(0)
	}

	fmt.Println("\nThe following table shows the names the old files will be renamed to.")
	for _, f := range renameMap {
		fmt.Printf("From:\t%s\nTo:\t%s\n\n", f.Old, f.New)
	}

	fmt.Println("\nIf OK, press ENTER to continue. To cancel, press CTRL + C")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	for _, f := range renameMap {
		if episodefile.FileExists(f.New) {
			fmt.Printf("File '%s' already exists. Skipping.\n", f.New)
		} else {
			episodefile.RenameFile(f.Old, f.New)
		}
	}

	return nil
}
