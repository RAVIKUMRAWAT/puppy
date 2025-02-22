package puppy

func Bark() string {
	return "Hello, Puppy!"
}

func BarkTwice() string {
	return Bark() + " " + Bark()
}