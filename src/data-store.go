package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type DataStore interface {
	getAllAlbums() ([]album, error)
	getAlbumById(id string) (album, error)
	addAlbum(album) error
}

type fileDataStore struct {
	pathToData string
}

func NewFileDataStore(pathToData string) *fileDataStore {
	fileDataStore := new(fileDataStore)

	fileDataStore.pathToData = pathToData
	return fileDataStore
}

func (fileDataStore *fileDataStore) getAllAlbums() ([]album, error) {
	albums := make([]album, 0)
	file, err := os.Open(fileDataStore.pathToData)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentAlbum := new(album)
		err := json.Unmarshal(scanner.Bytes(), currentAlbum)
		if err != nil {
			fmt.Printf("Error parsing data")
			return nil, err
		}
		albums = append(albums, *currentAlbum)
	}

	return albums, nil
}

func (fileDataStore *fileDataStore) getAlbumById(id string) (album, error) {

	albums, err := fileDataStore.getAllAlbums()
	if err != nil {
		return album{}, err
	}

	// search for specific album
	for i := 0; i < len(albums); i++ {
		if albums[i].ID == id {
			return albums[i], nil
		}
	}

	return album{}, fmt.Errorf("album with id %s not found", id)
}

func (fileDataStore *fileDataStore) addAlbum(album album) error {
	file, _ := os.OpenFile(fileDataStore.pathToData, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	bytes, err := json.Marshal(album)
	if err != nil {
		fmt.Printf("Error writing data")
		return err
	}

	_, _ = file.WriteString("\n");
	_, err = file.Write(bytes)
	return err
}
