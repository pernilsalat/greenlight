package mocks

import "greenlight.net/internal/data"

type MockMovieModel struct{}

func (m MockMovieModel) Insert(movie *data.Movie) error {
	// Mock the action...
	return nil
}
func (m MockMovieModel) Get(id int64) (*data.Movie, error) {
	// Mock the action...
	return nil, nil
}
func (m MockMovieModel) Update(movie *data.Movie) error {
	// Mock the action...
	return nil
}
func (m MockMovieModel) Delete(id int64) error {
	// Mock the action...
	return nil
}
