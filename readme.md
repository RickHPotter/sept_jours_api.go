# Concepts

## Decoding Data Using JSON

> Go Programming Language for Dummies

JavaScript Object Notation is a standard, text-based format for representing data using the JavaScript object syntax. It's a lightweight and easy-to-parse data representation language, commonly used for communicating between we servers and clients.

JSON supports the following data types:
String, Boolean, Number, Array, and lastly, Null.

``` golang
{
    "name": "CR7",      // string
    "auth": {           // object
        "secret-key": "supersecret",
        "token": "very-big-token",
    }
    "is Member": true,   // boolean
    "dependent": 3,     // number
    "contact": [        // array
        {
            "type": "Github",
            "contact": "http..."
        },
        {
            "type": "LinkedIn",
            "contact": "http..."
        }
    ],
    "retirement": null   // null
}
```

In Go, a `struct` is used in the process of encoding and decoding a JSON.

First, create the struct that will serve as a vessel for JSON.

``` golang
type Object struct {
    name string
    strct auth {
        secret-key string
        token string
    }
    is Member bool  // error
    dependent int
    contact []Contact
    retirement string
}

type Contact struct {
    type string
    contact string
}
```

As you can see, variables in Go, as in any other languages can't have space in their names. This is easily countered if you get used to mapping custom attribute names. Therefore, `is Member` can be any other name (literally) and the decode/encode will still work as long as you map its name after declaring the type.

```golang
    //
    PremiumMember bool `json:"is Member"`
    //
```

In fact, this should be done in all variables of the struct given that very often we'll need to have them avaialable for use in other packages and for that to happen, they have to be capitalised which is not always the case in JSONs.

After defining a struct so that the JSON string can map onto it, you can use `json.Unmarshal()` to decode the JSON string.

``` golang
    type Object struct {
        AttributeOne string `json:"attribute-one"`
    }

    jsonStringOrJsonFile := 
    `{
        "attribute-one": "value"    
    }`

    var object Object

    // json.Unmarshal(from []byte, to *struct)
    json.Unmarshal([]byte(jsonStringOrJsonFile), &object) 

    fmt.Println(object.AttributeOne) // value
```

You could also retrieve a value off JSON without the need of a struct. In this case, you would use:

```golang
    var jsonName map[string]interface{}
    json.Unmarshal([]byte(jsonStringOrJsonFile), &jsonName)

    fmt.Println(jsonName["attribute-one"]) // value
```

## Encoding Struct into JSON

From the same package as `json.Unmarshal()`, there is `json.Marshal()`. And even `json.MarshalIdent()` which idents the JSON output for better readability.

```golang
    object := Object {AttributeOne: "value" }
    // json.MarsahlIdent(object, prefixEachLine, prefixEachNewLine)
    objectJSON, err := json.MarsahlIdent(object, "", "    ")
    if err != nil {
        panic(err.Error())
    } else {
        // objectJSON is a []byte, therefore it needs to converted
        // caution: fn not implemented
        fn.WriteString(string(objectJSON)) 
    }
```

## GORM

?? TODO: ??

## REST Api

REST stands for Representational State Transfer, it's a software architetural style that defines how a web service should work and behave. REST is designed to take advantage of existing protocols, one that is usually a must is HTTP.

### HTTP

An HTTP request is made of a `Header` (contains metadata, such as encoding information, HTTP methods, such as GET, POST, and so on), and a `Body` (data to transmit over the network, can contain data in any format, specified in the Content-Type field).

An HTTP response is what is sent back after a request is retrieved by the API. It includes the `StatusCode` and an optional result expcted by the client, such a list of users in a GET Request.

## Considerations

I'd like to state the difference between GIN, Gorilla Mux (I haven't done this one yet), and net/http.

> [Source] <https://golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http>

### GIN

Gin is the most popular framework used by developers in Golang. It is extremely fast web framework and it suits the requirements of developers when they create microservices and web applications.

With Gin, developers can wrap up the logic of a code within a very few statements. This is because most of the work has been done in the framework itself.

Gin is open-sopuce, lightning fast, supports middleware, has routing and has rendering facilities.

### GORILLA

Gorilla Mux is one of the most powerful routers that helps in simplifying the route definitions. The multiplexer is capable of multiplexing HTTP routes to various handlers.

It is a package that customises the built-in HTTP router. It has a ton of capabilities that help developers of web apps work more efficiently. The package can be use din conjunction with other HTTP libraries.

In comparison with Gin, developers tend to agree that routing with Gorilla Mux is easier and more flexible. The developers also think that it is wise to stick to lightweight packages like Gorilla Mux when there's no need to use high-end features found in Gin framework.

### NET/HTTP

It is one of the most important and simplest packages that any developer has to be familiar with. This package allows you to create powerful HTTP Servers in Golang with potential compositional constructs.

In comparison with Gorila Mux, developers mostly agree that the net/http package is sufficient for creating Web Servers, although it's true that Gorila Mux eases out the development process.

In comparison with Gin, nett/http is more boilerplate and less handy facilities, including deserialisation wrappers, so you won't have to bother about the boilerplate surrounding the `json.Unmarshal` operation.
