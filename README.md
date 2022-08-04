# #1 Guessing Numbers Game

Guessing Numbers Game is a Golang project to deal with the computer in guessing random numbers. 

If the guess is incorrect, then the first player tells the second player whether the guess was too high or too low. The game loops on, until the guess is correct.

## Installation

[go.dev](https://go.dev/dl/go1.19.darwin-amd64.pkg) to install golang and check for go version.

```bash
go version
```

## Code 

```python
import( 
    "fmt" 
    "math/rand"
    "time"
)

# to see if the numbers match
func isMatching(secretNumber, guess int) bool

# to retain the user's input
func getUserInput() int

# to generate random numbers
func getRandomNumber() int
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[LYFEYTECH](https://github.com/lyfeytech)
