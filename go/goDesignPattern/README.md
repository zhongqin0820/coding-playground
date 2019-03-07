# Created at: 2019-03-06

# Contents
## Creational Patterns
### Singleton
#### Key Factors
- define a struct
- global private member `instance` in the package scope
- check `nil` when call `Init()` method
- `GetInstance()` return a corresponding pointer object

#### How to use
- use `GetInstance()` to obtain the `instance` to call the method

### Builder 
#### Key Factors
- `Manufactor`: has a member of `BuildProcess`
- `BuildProcess`: interfaces of the definition of the whole process and the product object
- `Product`: struct of the corresponding attributes of the whole process
- `XXXBuilder`: implements the `BuildProcess` interface and has a member of `Product`

#### How to use
1. initilize an object of `Manufactor` m
2. initilize an object of `XXXBuilder` `b`(uilder)
3. set up the buider of `m`: `m.SetBuider(b)`
4. run the produce pipeline : `m.Build()`
5. the corresponding `Product` of `XXXBuilder` build is :`var x = b.GetProduct()`
