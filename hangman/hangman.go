package hangman

import (
  "html/template"
  "os"
  "fmt"
  "sort"
  "strings"
)

/* Marshalls required Structs and outputs to console */
type Hangman struct {
  guesses int
  userCharactersList []string
  userCorrectCharactersList []string
}

type Person struct {
    UserName string
}

/*
 * Starts the application loop
 * Continues until the solution is found and the player wins
 * [or]
 * The hangman is complete and the player looses
 */
func (hangman Hangman) Run() {
  fmt.Println("Welcome to Hangman!")

  hangman.guesses = 0
  characterFromUser := ""

  dictionary := NewDictionary()
  wordToFind := dictionary.FetchRandomWord()
	screenOutputer := screenOutputerFactory()

  complete := false
  for (complete == false) {
    if (hangman.guesses == 11) {
      fmt.Println("Game over!")
      break
    }

    screenOutputer.OutputChooseLetter(&characterFromUser, wordToFind, hangman.userCorrectCharactersList)

    result := false

    if (len(characterFromUser) != 1) {
      result = false
    } else if (stringInSlice(characterFromUser, hangman.userCharactersList)) {
      result = false
    } else {
      hangman.userCharactersList = append(hangman.userCharactersList, characterFromUser)
      result = dictionary.IsCharacterInWord(characterFromUser, wordToFind)
    }

    if (result == true) {
      hangman.userCorrectCharactersList = append(hangman.userCorrectCharactersList, characterFromUser)
      arrayChars := strings.Split(wordToFind, "")

      sort.Strings(hangman.userCorrectCharactersList)
      sort.Strings(arrayChars)
      complete = IsArrayEqual(arrayChars, hangman.userCorrectCharactersList)
    }

    hangman.displayResult(screenOutputer, result)
  }
}

func IsArrayEqual(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for index, value := range a {
        if value != b[index] {
            return false
        }
    }
    return true
}

/*
 * displays the result using ScreenOutputer to the user
 */
func (hangman *Hangman) displayResult(screenOutputer ScreenOutputerInterface, result bool) {
  notFoundTemplate := template.New("Not Found Template")
  notFoundTemplate, _ = notFoundTemplate.Parse("Sorry letter not found you have {{.Guesses}} guesses left \n")

  if (result == false) {
    fmt.Println(screenOutputer.selectImage(hangman.guesses))
    hangman.guesses++

    notFoundTemplate.Execute(os.Stdout, struct {
      Guesses int
    }{
      11 - hangman.guesses,
    })


  } else {
    fmt.Println("Well done you guessed correctly")
  }
}
