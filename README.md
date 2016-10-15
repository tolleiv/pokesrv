# Pokesrv demo app.

This web app is used as a demo application which brings up a HTTP service with some Pokemon content.

The actual content can be influenced by ane nvironment variable, a command line flag or the hostname.

## Usage

    # Use the hostname to pick a pseudo-random pokemon
    go run main.go
    
    # Bypassing the randomness with an environment variable 
    POKEMON=gengar go run main.go
    
    # Bypoassing the randomness with a command line parameter
    go run main.go -pokemon pikachu
    
    
## References

  The app makes use of other projects. The [Gin](https://github.com/gin-gonic/gin) HTTP wrapper is used to provide the HTTP parser, The [Pokemon Hollow](http://www.dafont.com/pokemon.font) font is used along with the great data from the [PokeAPI](https://pokeapi.co) project.
   
## License

MIT License
    
    