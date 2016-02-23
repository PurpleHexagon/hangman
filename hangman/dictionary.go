package hangman

import (
  "math/rand"
	"time"
  "strings"
)

/*
 * Contains all words used in the game play
 * Provides API to query dictionary
 */
type dictionary struct {
  words []string
  maxCount int
}

/*
 * Initialises the dictionary setting default values
 */
func NewDictionary() *dictionary {
  dictionaryInstance := new(dictionary)
  dictionaryInstance.words = []string{"magic", "pies", "random", "cats", "cabbage"}
  dictionaryInstance.maxCount = len(dictionaryInstance.words)

  return dictionaryInstance
}

/*
 * Fetches a random word from the dictionary
 */
func (dictionary dictionary) FetchRandomWord() string {
	rand.Seed(time.Now().UnixNano())

	return dictionary.words[rand.Intn(dictionary.maxCount)]
}

/*
 * Checks to see if the character from the user is in the current games word
 */
func (dictionary dictionary) IsCharacterInWord(character string, word string) bool {
  if (strings.Contains(word, character)) {
      return true
    }

  return false
}
