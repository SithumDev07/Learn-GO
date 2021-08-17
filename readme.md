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

#### naked return
    we don't have to return value explicitly, 
    func lenAndUpper(name string) (length int, uppercase string) {
        length = len(name)
        uppercase = strings.ToUpper(name)
        return
    }

#### we can execute something after the function executed

    func lenAndUpper(name string) (length int, uppercase string) {
        defer fmt.Println("I' Done") -> this will execute after the function executed
        length = len(name)
        uppercase = strings.ToUpper(name)
        return
    }

## Loops
    only have for loop in GO
    func superAdd(numbers ...int) int {
        for index, number := range numbers { -> index will give the index and number will give the value
            fmt.Println(index, number)
        }
        return 1
    }

##### Other way doing
    func superAdd(numbers ...int) int {
        for i := 0; i < len(numbers); i++ {
            fmt.Println(numbers[i])
        }
        return 1
    }

### Conditions

    func canIDrink(age int) bool {
        if age < 18 {
            return false
        }
        return true
    }

### Variable expressions

Creating Varibale Exclusively for if condition

    if sriLankanAge := age + 2; sriLankanAge < 18 {
		return false
	}
	return true

### Switch Statements

    switch sriLankanAge := age + 2; sriLankanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return false
	}
	return false

### Arrays

    names := [5]string{"Sithum", "Dashantha"}

### Slices

Arrays in go without the length, append will return a new value rather than modifying the array

    names := []string{"GO", "Python"}
	names = append(names, "JavaScript")
	fmt.Println(names)