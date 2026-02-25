package test_cfg

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var arpaRegexp = regexp.MustCompile(`^(?:[0-9a-f]\.){1,32}ip6\.arpa$`)

type uniqueRand struct {
	bound     uint
	generated map[int]struct{}
}

func (u *uniqueRand) Int() int {
	for {
		var i int
		if u.bound > 0 {
			i = rand.Intn(int(u.bound))
		} else {
			i = rand.Int()
		}

		if _, ok := u.generated[i]; !ok {
			u.generated[i] = struct{}{}
			return i
		}
	}
}

func (u *uniqueRand) Hex() string {
	return fmt.Sprintf("%x", u.Int())
}

func newUniqueRand(bound uint) *uniqueRand {
	return &uniqueRand{bound: bound, generated: make(map[int]struct{})}
}

func generateSubDomains(template string, bound int, count int) []string {
	if bound < count && bound > 0 { // bound == 0 means no bound
		panic("bound must be greater than len")
	}

	generator := newUniqueRand(uint(bound))

	domains := make([]string, count)
	for j := 0; j < count; j++ {
		domains[j] = fmt.Sprintf(template, generator.Int())
	}
	return domains
}

func ValidateArpaDomain(arpaDomain string) error {
	if !arpaRegexp.MatchString(arpaDomain) {
		return errors.New("value must be a valid ARPA domain")
	}
	return nil
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func extractArpaSegments(arpaDomain string) ([]string, error) {
	if err := ValidateArpaDomain(arpaDomain); err != nil {
		return nil, err
	}

	// Remove the trailing ".ip6.arpa"
	trimmed := arpaDomain[:len(arpaDomain)-len(".ip6.arpa")]
	// Split the remaining string into segments
	segments := strings.Split(trimmed, ".")
	return segments, nil
}

func generateArpaSubDomains(arpaDomain string, randomSegments int, count int) []string {
	arpaSegments, err := extractArpaSegments(arpaDomain)
	arpaPosition := 32 - len(arpaSegments)
	if err != nil {
		panic(err)
	}

	// Ensure that the number of bytes to generate does not exceed the available segments for randomization
	if randomSegments < 0 || randomSegments > arpaPosition {
		panic("invalid randomSegments length")
	}

	// Each hex digit has 16 possible values, so we need 16^randomSegments possibilities
	generator := newUniqueRand(uint(1 << (randomSegments * 4)))

	domains := make([]string, count)
	for j := 0; j < count; j++ {
		// Create a slice to hold the segments of the new domain
		domainSegments := make([]string, 32)
		// Overwrite the last segments with the arpa segments
		copy(domainSegments[arpaPosition:], arpaSegments)
		// Fill the first segments with random hex values
		hex := generator.Hex()
		for pos, char := range ReverseString(hex) {
			domainSegments[pos] = string(char)
		}
		// Fill remaining segments with "0"
		length := []rune(hex)
		for i := len(length); i < arpaPosition; i++ {
			domainSegments[i] = "0"
		}
		// Join the segments to form the full domain
		domains[j] = strings.Join(domainSegments, ".") + ".ip6.arpa"
	}
	return domains
}
