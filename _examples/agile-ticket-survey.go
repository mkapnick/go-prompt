package main

import (
	"fmt"
	"os"
)

var communicationOptions = []string{
	"Good - no issues",
	"Mostly good - some miscommunication occurred along the way",
	"Somewhat good - too many issues occurred",
	"Not good - no one was talking to eachother",
	"N/A",
}

var effortOptions = []string{
	"Accurate - no unforseen issues came up",
	"Mostly accurate - there were some things that came up that made this either more or less work",
	"Somewhat accurate - didn't think through every possible use case for this ticket",
	"Not accurate - completely miscalculated",
	"N/A",
}

var organizationOptions = []string{
	"Organized - workflow was smooth",
	"Mostly organized - workflow was smooth with a hiccup",
	"Somewhat organized - workflow could have been a lot smoother",
	"Not organized - workflow was confusing",
	"N/A",
}

func main() {

	if _, err := os.Stat("/tmp/agileStats.txt"); err != nil {
		file, err := os.Create("/tmp/agileStats.txt")
		if err != nil {
			fmt.Println("Could not create /tmp/agileStats.txt file")
			panic(err)
		}
	}

	println(file)

	/*ticket := prompt.String("Ticket reference number (CW-112):")
	i := prompt.Choose("How well was communication on this ticket? ", communicationOptions)
	j := prompt.Choose("How accurate was the effort on this ticket? ", effortOptions)
	k := prompt.Choose("How organized was the workflow of this ticket? ", organizationOptions)
	*/
}
