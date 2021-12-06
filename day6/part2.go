package day6

/*
Suppose the lanternfish live forever and have unlimited food and space. Would
they take over the entire ocean?

After 256 days in the example above, there would be a total of 26984457539
lanternfish!
*/
func Part2(path string) (int, error) {
	return runFishSim(path, 256)
}
