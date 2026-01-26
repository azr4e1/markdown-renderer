package main

const (
	BOLDDELIMITER1      = "**"
	BOLDDELIMITER2      = "__"
	ITALICDELIMITER1    = "*"
	ITALICDELIMITER2    = "_"
	UNDERLINEDELIMITER  = "~"
	INLINECODEDELIMITER = "`"
	CROSSEDDELIMITER    = "-"
)

func LineParser(nodes []Text) []Text {
	newNodes := []Text{}
	for _, n := range nodes {
		switch node := n.(type) {
		case Plain:
			return nil
		default:
			newNodes = append(newNodes, n)
			continue
		}
	}
	return nil
}
