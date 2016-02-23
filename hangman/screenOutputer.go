package hangman

import (
  "fmt"
)

/*
 * Outputs to the console to allow users to play the game
 */
type ScreenOutputer struct {
	Images []string
}

/* Selects the correct image and outputs to console */
func (screenOutputer ScreenOutputer) selectImage(guesses int) string {
	return screenOutputer.Images[guesses]
}

/*
 * Outputs the basic game question to user
 */
func (ScreenOutputer ScreenOutputer) OutputChooseLetter(characterFromUser *string, wordToFind string, charactersFromUser []string) {
  for _, character := range wordToFind {
    if (stringInSlice(string(character), charactersFromUser)) {
        fmt.Printf(string(character))
    } else {
        fmt.Printf("_ ")
    }
  }

  fmt.Println("\n\n")
  fmt.Printf("Please select a letter: ")
  fmt.Scan(characterFromUser)
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

/*
 * Factory for screenOutputer
 * Sets array of images to each iteration of hangman image
 */
func screenOutputerFactory() *ScreenOutputer {
	return &ScreenOutputer{Images: []string{
		`





         _______
        `,
		`
        |
        |
        |
        |
        |
        |_______
        `,
		`	  ________
        |
        |
        |
        |
        |
        |_______
        `,
		`	  ________
        |/
        |
        |
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |
        |
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |      O
        |
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |      O
        |      |
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |      O
        |     _|
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |      O
        |     _|_
        |
        |
        |_______
        `,
		`	  ________
        |/      |
        |      O
        |     _|_
        |      /
        |
        |_______
        `,
		`	  ________
    	  |/      |
    	  |      O
    	  |     _|_
    	  |      /\
    	  |
    	  |_______
    	  `},
	}
}
