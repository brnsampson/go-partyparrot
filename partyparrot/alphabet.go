package partyparrot

import (
	"fmt"
	"math/rand"
	"strings"
)

// PartyParrots are the allowed partyparrot slack emojis we can choose from.
var PartyParrots = [...]string{
	":partyparrot:",
	":rightparrot:",
	":middleparrot:",
	":boredparrot:",
	":shuffleparrot:",
	":aussieparrot:",
	":parrotcop:",
	":gothparrot:",
}

var enSpace = "\u2002"

// PartyConv converts a single rune into a string of party parrot emojis. These are slack style and
// you can add to this list above in the PartyParrots array.
func PartyConv(r rune) string {
	a := partyMap(r)
	parrot := pickParrot()
	space := fmt.Sprintf("%s%s%s", enSpace, enSpace, enSpace)
	var build strings.Builder
	for _, b := range a {
		for _, v := range b {
			if v {
				build.WriteString(fmt.Sprintf("%s%s", parrot, space))
			} else {
				build.WriteString(fmt.Sprintf("%s%s", space, space))
			}
		}
		build.WriteString(fmt.Sprint("\n\n"))
	}
	return build.String()
}

func pickParrot() string {
	i := rand.Intn(len(PartyParrots))
	return PartyParrots[i]
}

// partyMap returns an 8 byte map to represent a rune. Each byte represents a line; on bits
// will become party parrots while the off bits will be whitespace later.
func partyMap(r rune) [4][4]bool {
	switch r {
	case '!':
		return [4][4]bool{
			{false, false, true, false},
			{false, false, true, false},
			{false, false, false, false},
			{false, false, true, false}}
	case '0':
		return [4][4]bool{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false}}
	case '3':
		return [4][4]bool{
			{true, true, true, true},
			{false, false, true, true},
			{false, false, false, true},
			{true, true, true, true}}
	case '6':
		return [4][4]bool{
			{false, true, true, false},
			{true, false, false, false},
			{true, true, true, true},
			{true, true, true, true}}
	case '9':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
			{false, false, false, true}}
	case '7':
		return [4][4]bool{
			{true, true, true, true},
			{false, false, false, true},
			{false, false, true, false},
			{false, true, false, false}}
	case '8':
		return [4][4]bool{
			{true, true, true, false},
			{true, false, true, true},
			{true, true, false, true},
			{false, true, true, true}}
	case '2':
		return [4][4]bool{
			{false, true, true, false},
			{true, false, false, true},
			{false, false, true, false},
			{true, true, true, true}}
	case '5':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, false},
			{false, true, true, true},
			{true, true, true, true}}
	case 'a':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, true}}
	case '1':
		return [4][4]bool{
			{false, true, true, false},
			{true, false, true, false},
			{false, false, true, false},
			{true, true, true, true}}
	case 'c':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, false, false, false},
			{true, true, true, true}}
	case 'b':
		return [4][4]bool{
			{true, false, false, false},
			{true, true, true, false},
			{true, false, true, false},
			{true, true, true, false}}
	case 'e':
		return [4][4]bool{
			{true, true, true, true},
			{true, true, false, false},
			{true, false, false, false},
			{true, true, true, true}}
	case 'd':
		return [4][4]bool{
			{true, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, false}}
	case 'g':
		return [4][4]bool{
			{true, true, true, false},
			{true, false, false, false},
			{true, false, false, true},
			{true, true, true, true}}
	case 'f':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, false},
			{true, false, false, false}}
	case 'i':
		return [4][4]bool{
			{true, true, true, true},
			{false, true, true, false},
			{false, true, true, false},
			{true, true, true, true}}
	case 'h':
		return [4][4]bool{
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, true},
			{true, false, false, true}}
	case 'k':
		return [4][4]bool{
			{true, false, false, true},
			{true, true, true, false},
			{true, true, true, false},
			{true, false, true, true}}
	case 'j':
		return [4][4]bool{
			{true, true, true, true},
			{false, false, false, true},
			{true, false, false, true},
			{false, true, true, false}}
	case 'm':
		return [4][4]bool{
			{true, false, true, true},
			{true, true, true, true},
			{true, true, false, true},
			{true, false, false, true}}
	case 'l':
		return [4][4]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, true, true, true}}
	case 'o':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, true}}
	case 'n':
		return [4][4]bool{
			{true, false, false, true},
			{true, true, false, true},
			{true, false, true, true},
			{true, false, false, true}}
	case 'q':
		return [4][4]bool{
			{false, true, true, true},
			{false, true, false, true},
			{false, true, true, true},
			{false, false, false, true}}
	case 'p':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, false}}
	case 's':
		return [4][4]bool{
			{true, true, true, true},
			{true, true, false, false},
			{false, false, true, true},
			{true, true, true, true}}
	case 'r':
		return [4][4]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, false},
			{true, false, true, true}}
	case 'u':
		return [4][4]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, true}}
	case 't':
		return [4][4]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, true, true, false}}
	case 'w':
		return [4][4]bool{
			{true, false, false, true},
			{true, false, true, true},
			{true, true, true, true},
			{false, true, false, false}}
	case 'v':
		return [4][4]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false}}
	case 'y':
		return [4][4]bool{
			{true, false, false, true},
			{true, true, true, true},
			{false, true, true, false},
			{false, true, true, false}}
	case 'x':
		return [4][4]bool{
			{true, false, false, true},
			{false, true, true, false},
			{false, true, true, false},
			{true, false, false, true}}
	case 'z':
		return [4][4]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, false, false},
			{true, true, true, true}}
	case '4':
		return [4][4]bool{
			{true, false, false, true},
			{true, true, true, true},
			{false, false, false, true},
			{false, false, false, true}}
	}
	return [4][4]bool{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false}}
}
