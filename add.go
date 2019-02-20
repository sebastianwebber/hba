package hba

import "fmt"

// Rules is a slice of Rule
type Rules []Rule

// Add adds a new rule into a slice of Rule
func (c *Rules) Add(in Rule) error {
	if ok, err := in.IsValid(); !ok {
		return fmt.Errorf("could not add new rule: %v", err)
	}

	*c = append(*c, in)

	return nil
}
