package server

type Test struct {
	Part1 Part1  `json:"part1,omitempty"`
	Part2 Part2  `json:"part2,omitempty"`
	Part3 Part3  `json:"part3,omitempty"`
}

type Part1 struct {
	Questions []string `json:"questions,omitempty"`
}

type Part2 struct {
	Questions []string `json:"questions,omitempty"`
}

type Part3 struct {
	Questions []string `json:"questions,omitempty"`
}

func NewTest() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "What sort of food do you like eating most?")
	part1.Questions = append(part1.Questions, "Who normally does the cooking in your home?")
	part1.Questions = append(part1.Questions, "Do you watch cookery programmers on TV?")
	part1.Questions = append(part1.Questions, "In general, do you prefer eating out or eating at home?")

	part2 := Part2{}
	part2.Questions = append(part1.Questions, "Describe a house / apartament that someone you know lives in.")
	part2.Questions = append(part1.Questions, "Whose house/apartament this is")
	part2.Questions = append(part1.Questions, "Where the house/ apartament is")
	part2.Questions = append(part1.Questions, "What it looks like inside")
	part2.Questions = append(part1.Questions, "and explain what you like or dislike about this person’s house/ apartament")

	part3 := Part3{}
	part3.Questions = append(part1.Questions, "What kinks of home are most popular in your country ?")
	part3.Questions = append(part1.Questions, "What do you think are the advantages of living in a house rather than an apartament?")
	part3.Questions = append(part1.Questions, "Do you think that everyone would like to live in a larger home? Why is that?")

	return Test{
		part1,
		part2,
		part3,
	}
}