package main

func main() {
	var y = []string{"A", "B", "C", "D"}
	var x = y[:3]

	for i, s := range x {
		print(i, s, ",")
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
