package main
type MyError struct {
	Msg string
	File string
	Line int
}

func (e *MyError) Error() string { 
	return fmt.Sprintf("%s:%d: %s”, e.File, e.Line, e.Msg)
}

func main (){
}
