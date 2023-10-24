package conditionals

import "fmt"

var actualVariable = "someGlobalVariable"

func Conditions(age int) (bool, string) {
	fmt.Printf("***IF, ELSE IF, ELSE CONDITIONS***\n")
	/*
		SHADOWING: when we create a local variable with thesame name as another one in an outer or global scope
	*/
	actualVariable := "somethingElse"
	_ = actualVariable //this removes the warning thrown by the compiler that canGoToClub variable is not used

	// conditionals in golang works almost as it does in javascript except for the fact that CONDITIONS dont go in a paranthesis ()
	var canGoToClub bool
	_ = canGoToClub

	if age < 18 {
		canGoToClub = false
		return canGoToClub, "Underaged"
	} else if age > 18 && age < 21 {
		canGoToClub = true
		return canGoToClub, "Of age but can't drink"
	} else {
		canGoToClub = true
		return canGoToClub, "You can do it all"
	}

	// An alternate way of writing IF statement which is just a syntactic sugar for the way above is to use the below apprach
	/*
		SYNTAX: if INITIAL_STATEMENT; CONDITION {}

		NB: This can only be used if the variable we are comparing is only use in the if statement;
		To use the below, comment the above format
	*/

	// if enteredAge := age; age < 18 {
	// 	_ = enteredAge
	// 	canGoToClub = false
	// 	return canGoToClub, "Underaged"
	// } else if age > 18 && age < 21 {
	// 	canGoToClub = true
	// 	return canGoToClub, "Of age but can't drink"
	// } else {
	// 	canGoToClub = true
	// 	return canGoToClub, "You can do it all"
	// }

}
