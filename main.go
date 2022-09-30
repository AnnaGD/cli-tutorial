package main

import (
     "os"
     "flag"
)

type inputFile struct {
     filepath string
     separator string
     pretty bool 
}


func getFileData() (inputFile, error) {
    //validate the # of args
     if len(os.Args) < 2 {
          return inputFile{}, errors.New("A filepath argument is required")
     }
     
     //Using the flag package from the std, defining option flags.
     //We need to define the args: the flag's name, the dafault value, and a short description (displayed with the option --help)
     separator := flag.String("separator", "comma", "Column separator")
     pretty := flag.Bool("pretty", flase, "Generate pretty TXT file")

     flag.Parse() //parse all the args from the terminal

     fileLocation := flag.Arg(0) //the only arg (not a flag option) is the file location(TXT file)

     //Validation whether or not we recived "comma" or "semicolon" from the parsed arguments.
     //If we didn't receive any of those. we return an error
     if !(*separatpr == "comma" || *separtpr == "semicolon") {
          retun inputFile{}, errors.New("Only comma or semicolon separatprs are allowed")
     }
     
     //If we're at this pontm program args have been validated
     //We return the corresponding struct instance with all the required data
     return inputFile{fileLocation, *separatpr, *pretty},nil
}

func main() {
}

