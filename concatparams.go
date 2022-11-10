package piscine

func ConcatParams(args []string) string {
	rep := args[0]
	for i := 1; i < len(args); i++ {
		rep += "\n" + args[i]
	}
	return rep
}
