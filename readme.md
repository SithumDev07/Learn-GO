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

Add in Dictionary

    var errorWordExists = errors.New("Word is already exists")

    func (d Dictionary) Add(word, def string) error {

        _, error := d.Search(word)

        switch error {
            case errorNotFound:
                d[word] = def
            case nil:
                return errorWordExists
        }
        return nil
    }

When adding word into dictionary

    dictionary := mydict.Dictionary{}

    baseword := "Hello"

    err := dictionary.Add(baseword, "First")

    if err != nil {
        fmt.Println(err)
    }

Update in dictionary

    var errorCantUpdate = errors.New("Cant Update non existing word)

    func (d Dictionary) Update (word, definition string) error {

        _, err := d.Search(word)

        switch err {
            case nil:
                d[word] = definition
            case errorNotFound:
                return errorCantUpdate
        }

        return nil
    }

When Updating a word dictionary

    dictionary := mydict.Dictionary{}

    baseword := "Hello"

    dictionary.Add(baseword, "First")

    err := dictionary.Update(baseword, "Second")

    if err != nil {
        fmt.Println(err)
    }

Deleting a word in dictionary

    var errorCantDelete = errors.New("Cant delete non exisiting word")

    func (d Dictionary) Delete(word string) error {

        _, err := d.Search(word)

        switch err {
            case nil:
                delete(d, word)
            case errorNotFound:
                return errorCantDelete
        }

        return nil

    }

When deleting a word from dictionary

    dictionary := mydict.Dictionary{}

    baseword := "Hello"

    dictionary.Add(baseword, "First")

    err := dictionary.Delete(baseword)

    if err != nil {
        fmt.Println(err)
    }

### creating variables as group

    var (
        errorNotFound = errors.New("Word not found")
        errorCantUpdate = errors.New("Cant update non existing word")
        errorCantDelete = errors.New("Cant delete non existing word")
    )

### hitURL()

using http standard library from go

    var requestFail = erros.New("Request Failed")

    func hitURL(url string) error {

        res, err := http.Get(url)

        if err == nil || res,StatusCode >= 400 {
            return requestFail
        }

        return nil
    }

### Maps

    var results map[string]string

    results["First"] = "Hello"

when we initialize a map and try to add the complier won't giving errors but at runtime there will be errors, because map will nil, can't add values into nil

So, we need to initialize map either,

    var results = map[string]string{}

or, using make function

    var results = make(map[string]string)

### Go Routines

    func Count(person string) {
        for i:=0; i<10; i++ {
            fmt.Println(person, "is looking good, i)
            time.Sleep(time.Second)
        }
    }

when calling, do concurrently, next to other

    go Count("Sithum")
    Count("Go")

note: Go routine is alive only main function is alive

### Channels

when usng go routines they only awake if the main function awakes, so we have to use channels in order to communicate with main function from go routuines

    func isGood(person string, c chan bool) {
        time.Sleep(time.Second * 3)
        fmt.Println(person)
        c <- true
    }

when calling from main

    c := make(chan bool)
    people := [2]string{"Sithum", "GoLang"}

    for _, person := range people {
        go IsGood(person, c)
    }

    fmt.Println(<-c)
    fmt.Println(<-c)

recieving is a blocking operation means that main funtion waits until recieves

### Fast URL Checker

    type result struct {
        url string
        status string
    }

    func hitURL(url string, c chan<- result) {
        res, err := http.Get(url)
        status := "OK"
        if err != nil || res.StatusCode >= 400 {
            status = "FAILED"
        }

        c <- result {
            url: url,
            status: status
        }
    }

When we action

    results := make(map[string]string)
    c := make(chan result)

    urls := []string {
        "https://www.google.com/",
    	"https://www.airbnb.com/",
    	"https://www.amazon.com/",
    	"https://www.reddit.com/",
    }

    for _, url := range urls {
        go hitURL(url, c)
    }

    for i:=0; i < len(urls); i++ {
        answer := <-c
        results[answer.url] = answer.status
    }

    for url, status := range results {
        fmt.Println(url, status)
    }
