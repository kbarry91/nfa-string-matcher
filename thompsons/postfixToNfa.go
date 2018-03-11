/*
Package thompsons is used to implement Thompsons construction of a postfix regular expresion to a Non Determistic Finite Automota
*/
package thompsons

/*
Represents a state in a nfa
*/
type state struct {
	symbol rune
	arrow1 *state
	arrow2 *state
}

/*
NfaFragment is a Helper struct to keep track of initial and accept state*/
type NfaFragment struct {
	initial *state
	accept  *state
}

/*
PostToNfa function takes in a postfix regular expression to construct a nfa
*/
func PostToNfa(postfix string) *NfaFragment {
	// create a stack as an array of pointers to empty nfaFragment's
	nfaStack := []*NfaFragment{}

	// step through the postfix regex to pop item off sdtack and push them as required
	for _, r := range postfix {
		switch r {
		case '.':
			//must pop 2 values off stack (This first to come off is fragment2 as it was added last)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1.accept.arrow1 = fragment2.initial
			// set fragment2 initial state as fragment1 accept

			//append  pointer to NfaFragment to the nfaStack
			nfaStack = append(nfaStack, &NfaFragment{initial: fragment1.initial, accept: fragment2.accept})
		case '|':
			//must pop 2 values off stack (This first to come off is fragment2 as it was added last)
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			// both arrows must point to initial state
			initial := state{arrow1: fragment1.initial, arrow2: fragment2.initial}

			// point fragment2 at new accept state
			fragment1.accept.arrow1 = &accept
			fragment2.accept.arrow1 = &accept

			//append  pointer to NfaFragment to the nfaStack assigning new initial and accept state
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial, accept: &accept})
		case '*':
			// only pop 1 fragment as * only uses one state
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			accept := state{}

			// inital points at new initial and ne accept state
			initial := state{arrow1: fragment.initial, arrow2: &accept}
			fragment.accept.arrow1 = fragment.initial
			fragment.accept.arrow2 = &accept
			//append  pointer to NfaFragment to the nfaStack assigning new initial and accept state
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, arrow1: &accept}
			//push the accept state to the stack as the new initial state
			nfaStack = append(nfaStack, &NfaFragment{initial: &initial})
		} // switch
	} // for
	return nfaStack[0]
}
