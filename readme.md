# GoLang

### const and var
    var name string = "golang"
    shorthand -> name := "golang"

    const name string = "golang"

    also, shorthand version is nor working outside a funtion
    then have to use as -> var name string = "golang"

### Functions
    func multiply(a int, b int) int {
        return a * b
    }

    Should say type of arguments and return type

    If aruguments are same typed -> 
        func multiply(a, b int) int {
            return a * b
        }
    
    in go, we can return many things ->
    func lenAndUpper(name string) (int, string) {
        return len(name), strings.ToUpper(name)
    }

    when we calling the function -> 
    totalLength, upperName := lenAndUpper("golang") -> this will return and assign in order

    Also we can ignore variables (Compiler look into them and ignore them)
    totalLength, _ := lenAndUpper("golang") -> uppercase (ignore)

#### we can pass many values in to functions
func repeatMe(words ...string) {
	fmt.Println(words)
} 

repeatMe("Powered", "Golang", "Learning", "Hustling") -> will return as a literal array