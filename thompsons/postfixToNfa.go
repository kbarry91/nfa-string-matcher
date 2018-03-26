/*
Package thompsons is used to implement Thompsons construction of a pofix regular expresion to a Non Determistic Finite Automota
*/
package thompsons

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

			// set fragment2 initial State as fragment1 accept
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

		default:
			accept := state{}
			initial := state{symbol: r, arrow1: &accept}
			//push the accept State to the stack as the new initial State
			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})

		} // switch
	} // for
	return nfaStack[0]
}
