package hangman

import (
  "testing"
)

/*
 * Outputs to the console to allow users to play the game
 */
type ScreenOutputerMock struct {
	Images []string
}

/* Selects the correct image and outputs to console */
func (screenOutputer ScreenOutputerMock) selectImage() string {
	return "test"
}

/*
 * Outputs the basic game question to user
 */
func (ScreenOutputer ScreenOutputerMock) OutputChooseLetter(characterFromUser *string, wordToFind string) {

}

func TestTranslate(testHelper *testing.T) {
  hangman := Hangman{}
  hangman.displayResult(ScreenOutputerMock{}, true)
}
