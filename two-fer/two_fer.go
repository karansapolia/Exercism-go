// Package twofer provides the ShareWith function which generates
// the classic "Two Fer" phrase: "One for <name>, one for me."
// If the provided name is empty, it defaults to "you".
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
