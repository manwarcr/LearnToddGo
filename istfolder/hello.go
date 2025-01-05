package main

import (
	"fmt"

	"example.com/greetings"
	//"/home/chris/mygo/greetings" //the greetings directory is 1 level above the hello directory
)

func main() {
	// get a greeting message and print it
	message := greetings.Hello2("Chris")
	fmt.Println(message)
}

/*
go mod init example/home/chris/mygo/hello //creates go.mod (depandacies file)
go mod edit -replace example.com/greetings=../greetings // gets rid of
                                                        example.com as we haven'y published yet
go mod tidy //checks all the packages in your project
              and their dependencies to clean up and
			  optimize your go. mod file.
go run .


*/
