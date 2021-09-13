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

### Map Data Structures

                //Key  	//Value
    Languages := map[string]string{"name": "GoLang", "age": "6"}
    fmt.Println(Languages)

Iterating the array

    for key, value := range Languages {
    	fmt.Println(key, value)
    }

### Structs

    type person struct {
        name    string
        age     int
        favFood []string
    }

    func main() {
        favFood := []string{"Peanut Butter", "Kottu", "Ice Cream"}
        Sithum := person{"Sithum", 21, favFood}
        fmt.Println(Sithum)
        fmt.Println(Sithum.favFood)
    }

Another Way

    Sithum := person{name: "Sithum", age: 21, favFood: favFood}
    fmt.Println(Sithum)

##### Go doesn't have classes and objects or constructors (We have to define ourselves)

### Structs Exporting

if Types are lowercase, it means they are private

private->

    //Account struct
    type Account struct {
        owner   string
        balance int
    }

public->

    //Account struct
    type Account struct {
        Owner   string
        Balance int
    }

### Functions that makes constructors

    package accounts

    // Account struct
    type Account struct {
        owner   string
        balance int
    }

    //NewAccount Creates Account
    func NewAccount(owner string) *Account {
        account := Account{owner: owner, balance: 0}
        return &account
    }

then->

    import (
        "github.com/SithumDev07/LearnGO/accounts"
    )

    account := accounts.NewAccount("Sithum")
    fmt.Println(account)

### Methods (in GO Recievers)

When we actually change the value we use have to exactly point to the value rather than making a copy

    // Deposti x Amount in account
    func (a *Account) Deposit(amount int) {
        a.balance += amount
    }

    // Balance of account
    func (a Account) Balance() int {
        return a.balance
    }

### Error Handling (No Try Catch in GO) -> we have to handle manually

    func (a *Account) Withdraw(amount int) error {
        if a.balance < amount {
            return errors.New("Account Balance is insufficient")
        }
        a.balance -= amount
        return nil
    }

Handling

    err := account.Withdraw(90000)
    if err != nil {
    	log.Fatalln(err)
    }

Premake erros

    var errNoMoney = errors.New("Bank Balance is insufficient")

in reciever

    if a.balance < amount {
    	return errNoMoney
    }

### Go Internal Methods

    fmt.Println(account) will look like &{Sithum 79500}

there is reciever call String() and we can use it to change the default behavior

    func (a Account) String() string {
        return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
    }

This will give an output like,

    Sithum's account.
    Has: 79500

### Dictionary

Dictionary is type of map

    package mydict

    type Dictionary map[string]string

Search in dictionary

    var errorNotFound = errors.New("Not Found")

    func (d Dictionary) Search(word string) (string, error) {
        value, exists := d[word]

        if exists {
            return value, nil
        }

        return "", errorNotFound
    }

When making a dictionary and search for a word

    dictionary := mydict.Dictionary{"first" : "GO Lang"}
    definition, error := mydict.Search("Python")

    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(definition)
    }
