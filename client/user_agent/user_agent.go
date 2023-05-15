package user_agent

import (
	"fmt"
	"strings"
)

type UserAgentProduct struct {
	Name    string
	Version string
	Comment string
}

func (p UserAgentProduct) String() string {
	builder := []string{}

	if p.Name != "" {
		if p.Version != "" {
			builder = append(builder, fmt.Sprintf("%s/%s", p.Name, p.Version))
		} else {
			builder = append(builder, p.Name)
		}
	}

	if p.Comment != "" {
		builder = append(builder, fmt.Sprintf("(%s)", p.Comment))
	}

	return strings.Join(builder, " ")
}

type UserAgentProducts []UserAgentProduct

func (ua UserAgentProducts) String() string {
	builder := make([]string, len(ua))

	for i, p := range ua {
		builder[i] = p.String()
	}

	return strings.Join(builder, " ")
}
