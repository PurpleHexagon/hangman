package hangman

type ScreenOutputerInterface interface {
  selectImage(guesses int) string
  OutputChooseLetter(characterFromUser *string, wordToFind string, charactersFromUser []string)
}
