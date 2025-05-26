package classes

//// Background color utilities

func Bg_white() Utility {
	css := "background-color: #ffffff;"
	class := "bg-white"
	return NewUtility(class, css)
}

func Bg_black() Utility {
	css := "background-color: #000000;"
	class := "bg-black"
	return NewUtility(class, css)
}
