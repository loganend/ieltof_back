package server

type Test struct {
	Part1 Part1 `json:"part1,omitempty"`
	Part2 Part2 `json:"part2,omitempty"`
	Part3 Part3 `json:"part3,omitempty"`
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

//func NewTest() Test {
//	part1 := Part1{}
//	part1.Questions = append(part1.Questions, "What sort of food do you like eating most?")
//	part1.Questions = append(part1.Questions, "Who normally does the cooking in your home?")
//	part1.Questions = append(part1.Questions, "Do you watch cookery programmers on TV?")
//	part1.Questions = append(part1.Questions, "In general, do you prefer eating out or eating at home?")
//
//	part2 := Part2{}
//	part2.Questions = append(part1.Questions, "Describe a house / apartament that someone you know lives in.")
//	part2.Questions = append(part1.Questions, "Whose house/apartament this is")
//	part2.Questions = append(part1.Questions, "Where the house/ apartament is")
//	part2.Questions = append(part1.Questions, "What it looks like inside")
//	part2.Questions = append(part1.Questions, "and explain what you like or dislike about this person’s house/ apartament")
//
//	part3 := Part3{}
//	part3.Questions = append(part1.Questions, "What kinks of home are most popular in your country ?")
//	part3.Questions = append(part1.Questions, "What do you think are the advantages of living in a house rather than an apartament?")
//	part3.Questions = append(part1.Questions, "Do you think that everyone would like to live in a larger home? Why is that?")
//
//	return Test{
//		part1,
//		part2,
//		part3,
//	}
//}

func getTest1() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "What kind of place is it?")
	part1.Questions = append(part1.Questions, "What’s the most interesting part of your town/village?")
	part1.Questions = append(part1.Questions, "What kind of jobs do the people in your town/village do?")
	part1.Questions = append(part1.Questions, "Would you say it’s a good place to live? (why?)")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe a house / apartament that someone you know lives in.")
	part2.Questions = append(part2.Questions, "Whose house/apartament this is")
	part2.Questions = append(part2.Questions, "Where the house/ apartament is")
	part2.Questions = append(part2.Questions, "What it looks like inside")
	part2.Questions = append(part2.Questions, "and explain what you like or dislike about this person’s house/ apartament")

	part3 := Part3{}
	part3.Questions = append(part3.Questions, "What kinks of home are most popular in your country ?")
	part3.Questions = append(part3.Questions, "What do you think are the advantages of living in a house rather than an apartament?")
	part3.Questions = append(part3.Questions, "Do you think that everyone would like to live in a larger home? Why is that?")
	part3.Questions = append(part3.Questions, "How can a teacher make lessons for children more interesting?")

	return Test{
		part1,
		part2,
		part3,
	}
}

func getTest2() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "Tell me about the kind of accommodation you live in?")
	part1.Questions = append(part1.Questions, "How long have you lived there?")
	part1.Questions = append(part1.Questions, "What do you like about living there?")
	part1.Questions = append(part1.Questions, "What sort of accommodation would you most like to live in?")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe something you own which is very important to you.")
	part2.Questions = append(part2.Questions, "where you got it from")
	part2.Questions = append(part2.Questions, "how long you have had it ")
	part2.Questions = append(part2.Questions, "what you use it for ")
	part2.Questions = append(part2.Questions, "and explain why it is important to you. ")

	part3 := Part3{}
	part3.Questions = append(part3.Questions, "What kind of things give status to people in your country? ")
	part3.Questions = append(part3.Questions, "Have things changed since your parents’ time? ")
	part3.Questions = append(part3.Questions, "Do you think advertising influences what people buy? ")
	part3.Questions = append(part3.Questions, "Do you think all information on the internet is true? ")

	return Test{
		part1,
		part2,
		part3,
	}
}

func getTest3() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "What sort of food do you like eating most?")
	part1.Questions = append(part1.Questions, "Who normally does the cooking in your home?")
	part1.Questions = append(part1.Questions, "Do you watch cookery programmers on TV?")
	part1.Questions = append(part1.Questions, "In general, do you prefer eating out or eating at home?")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe a piece of art you like.")
	part2.Questions = append(part2.Questions, "what the work of art is")
	part2.Questions = append(part2.Questions, "when you first saw it")
	part2.Questions = append(part2.Questions, "what you know about it")
	part2.Questions = append(part2.Questions, "and explain why you like it.")

	part3 := Part3{}
	part3.Questions = append(part3.Questions, "Can clothing tell you much about a person?")
	part3.Questions = append(part3.Questions, "Why do some companies ask their staff to wear uniforms?")
	part3.Questions = append(part3.Questions, "What are the advantages and disadvantages of having uniforms at work?")
	part3.Questions = append(part3.Questions, "How can people protect the environment?")


	return Test{
		part1,
		part2,
		part3,
	}
}

func getTest4() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "Do you enjoy your birthdays?")
	part1.Questions = append(part1.Questions, "Do you usually celebrate your birthday?")
	part1.Questions = append(part1.Questions, "What did you do on your last birthday?")
	part1.Questions = append(part1.Questions, "Can you remember a birthday you enjoyed as a child?")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe a book you have recently read.")
	part2.Questions = append(part2.Questions, "what kind of book it is")
	part2.Questions = append(part2.Questions, "what it is about")
	part2.Questions = append(part2.Questions, "what sort of people would enjoy it")
	part2.Questions = append(part2.Questions, "and explain why you liked it.")

	part3 := Part3{}
	part3.Questions = append(part3.Questions, "What role should the teacher have in the classroom?")
	part3.Questions = append(part3.Questions, "Do you think computers will one day replace teachers in the classroom?")
	part3.Questions = append(part3.Questions, "How has teaching changed in your country in the last few decades?")
	part3.Questions = append(part3.Questions, "Topics and questions for speaking part 1 and speaking part 2.")


	return Test{
		part1,
		part2,
		part3,
	}
}


func getTest5() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "Can you describe the house where you live to me?")
	part1.Questions = append(part1.Questions, "What is there to do in the area where you live?")
	part1.Questions = append(part1.Questions, "What do you like about the area where you live?")
	part1.Questions = append(part1.Questions, "How do you think it could be improved?")
	part1.Questions = append(part1.Questions, "Do you think it is better to live in the centre of town or outside in the country? Why?")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe someone you respect.")
	part2.Questions = append(part2.Questions, "who the person is")
	part2.Questions = append(part2.Questions, "how you know about this person")
	part2.Questions = append(part2.Questions, "what this person does")
	part2.Questions = append(part2.Questions, "what this person is like")
	part2.Questions = append(part2.Questions, "and explain why you respect this person.")


	part3 := Part3{}
	part3.Questions = append(part3.Questions, "Why do some people prefer to travel abroad rather than in their own country?")
	part3.Questions = append(part3.Questions, "Do you think traveling to another country can change the way people think?")
	part3.Questions = append(part3.Questions, "Do you think it is good for children to experience life in a foreign country?")
	part3.Questions = append(part3.Questions, "How have holidays changed over the past few decades?")


	return Test{
		part1,
		part2,
		part3,
	}
}


func getTest6() Test {
	part1 := Part1{}
	part1.Questions = append(part1.Questions, "Do you enjoy reading? Why?")
	part1.Questions = append(part1.Questions, "What sort of things do you read?")
	part1.Questions = append(part1.Questions, "Tell me something about your favourite book.")
	part1.Questions = append(part1.Questions, "What are the advantages of reading instead of watching television or going to the cinema?")

	part2 := Part2{}
	part2.Questions = append(part2.Questions, "Describe an unexpected event.")
	part2.Questions = append(part2.Questions, "what it was")
	part2.Questions = append(part2.Questions, "when it happened")
	part2.Questions = append(part2.Questions, "who was there")
	part2.Questions = append(part2.Questions, "and explain why you enjoyed it.")

	part3 := Part3{}
	part3.Questions = append(part3.Questions, "What is the difference between white collar and blue collar jobs?")
	part3.Questions = append(part3.Questions, "What skills do you think are needed to get a good job these days?")
	part3.Questions = append(part3.Questions, "Do you think women should be able to do all the same jobs that men do?")
	part3.Questions = append(part3.Questions, "What jobs do you think are most valuable to society?")


	return Test{
		part1,
		part2,
		part3,
	}
}