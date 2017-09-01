package main

import "github.com/mkapnick/go-prompt"

var communicationOptions = []string{
	"Good - no issues",
	"Mostly good - some miscommunication occurred along the way",
	"Somewhat good - too many issues occurred",
	"Not good - no one was talking to eachother",
}

var effortOptions = []string{
	"Accurate - no unforseen issues came up",
	"Mostly accurate - there were some things that came up that made this either more or less work",
	"Somewhat accurate - didn't think through every possible use case for this ticket",
	"Not accurate - completely miscalculated",
}

var organizationOptions = []string{
	"Organized - workflow was smooth",
	"Mostly organized - workflow was smooth with a hiccup",
	"Somewhat organized - workflow could have been a lot smoother",
	"Not organized - workflow was confusing",
}

func main() {
	i := prompt.Choose("How well was communication on this ticket? ", communicationOptions)
	println("==> You chose: " + communicationOptions[i])

	j := prompt.Choose("How accurate was the effort on this ticket? ", effortOptions)
	println("==> You chose: " + effortOptions[j])

	k := prompt.Choose("How organized was the workflow of this ticket? ", organizationOptions)
	println("==> You chose: " + organizationOptions[k])
}
