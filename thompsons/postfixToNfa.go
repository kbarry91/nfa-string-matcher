/*
Package thompsons is used to implement Thompsons construction of a pofix regular expresion to a Non Determistic Finite Automota
*/
package thompsons

import "fmt"

/*
Represents a State in a nfa
*/
type state struct {
	symbol rune
	arrow1 *state
	arrow2 *state
}

/*
NfaFragment is a Helper struct to keep track of initial and accept State*/
type nfaFragment struct {
	initial *state
	accept  *state
}

/*
PostToNfa function takes in a pofix regular expression to construct a nfa
*/
func PostToNfa(pofix string) *nfaFragment {
	// create a stack as an array of pointers to empty nfaFragment's
	nfaStack := []*nfaFragment{}

	// step through the pofix regex to pop item off stack and push them as required
	for _, r := range pofix {
		switch r {

		case '.':
			//must pop 2 values off stack (This first to come off is fragment2 as it was added last)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// set  fragment1 accept as fragment2 initial State
			fragment1.accept.arrow1 = fragment2.initial
			//append  pointer to NfaFragment to the nfaStack
			nfaStack = append(nfaStack, &nfaFragment{initial: fragment1.initial, accept: fragment2.accept})

		case '|':
			//must pop 2 values off stack (This first to come off is fragment2 as it was added last)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//accept := state{}
			// both arrows must point to initial State
			initial := state{arrow1: fragment1.initial, arrow2: fragment2.initial}
			accept := state{}

			// point fragment2 at new accept State
			fragment1.accept.arrow1 = &accept
			fragment2.accept.arrow1 = &accept

			//append  pointer to NfaFragment to the nfaStack assigning new initial and accept State
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		case '*':
			// only pop 1 fragment as * only uses one State
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			accept := state{}

			// inital points at new initial and ne accept State
			initial := state{arrow1: fragment.initial, arrow2: &accept}

			fragment.accept.arrow1 = fragment.initial
			fragment.accept.arrow2 = &accept

			//append  pointer to NfaFragment to the nfaStack assigning new initial and accept State
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		case '+':
			// pop offtsack
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}

			//set initial state
			initial := state{arrow1: fragment.initial, arrow2: &accept}

			// loop the arrow to make initial  state accept state
			fragment.accept.arrow1 = &initial

			// add the fragment to the stack
			nfaStack = append(nfaStack, &nfaFragment{initial: fragment.initial, accept: &accept})

		case '?':
			// pop offtsack
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// set initial state
			initial := state{arrow1: fragment.initial, arrow2: fragment.accept}

			// apend fragment to the stack
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: fragment.accept})

		default:
			accept := state{}
			initial := state{symbol: r, arrow1: &accept}
			//push the accept State to the stack as the new initial State
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		} // switch
	} // for

	// must handle and debug any issues
	if len(nfaStack) != 1 {
		fmt.Println("nfa Invalid", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

/*
stringMatcher compares a pofix expression to a string and returns either true or false
*/
func StringMatcher(pofix string, userInput string) bool {

	// default isMatch to false
	isMatch := false

	// create nfa from regular expression
	ponfa := PostToNfa(pofix)

	// need to be able to access current state and the following next states

	curState := []*state{}
	nextState := []*state{}

	//  add initial state of pofix nfa and all available states
	curState = addNextState(curState[:], ponfa.initial, ponfa.accept)

	// loop through the user entered string a character at a time
	for _, char := range userInput {

		// check all current states
		for _, cur := range curState {

			// if current state is set to the rune  add that state and other other state that can be accesed
			if cur.symbol == char {
				nextState = addNextState(nextState[:], cur.arrow1, ponfa.accept)
			}
		}
		// swap curState with nextState to make the nextState the new curState for next iteration
		curState, nextState = nextState, []*state{}
	}

	// iterate through the current states to check if they are accepted by nfa
	for _, cur := range curState {

		//if current state equals accept state of nfa
		if cur == ponfa.accept {
			isMatch = true
			break
		}
	}
	return isMatch
}

/*
	addState gets the current statewhile checking if the state has a e arrow and follows the arrow  to get next states
*/
func addNextState(nextState []*state, s *state, a *state) []*state {
	nextState = append(nextState, s)

	// a s.symbol with 0 rune  means state has e arrows coming from it
	if s != a && s.symbol == 0 {
		nextState = addNextState(nextState, s.arrow1, a)

		// if there is another edge/arrow it must be added
		if s.arrow2 != nil {
			nextState = addNextState(nextState, s.arrow2, a)
		}
	}
	return nextState
}
