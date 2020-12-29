# foaas-api

A simple Go wrapper for [FOAAS](https://foaas.com).

### Usage

By default, the client will return English plain text.

```go
fuck := foaas.Client{}

awesome, err := fuck.Awesome("Baphy")
if err != nil {
	log.Fatalln(err)
}

fmt.Println(awesome)
```

The defaults can be changed like this.

```go
fuck := foaas.Client{}
//  Accepts ISO 639-1 language codes. E.G. es, en, de.
fuck.Lang = "es"    
//  Uses shoutcloud to return all-caps fucks.
fuck.Shout = true
//  Available formats are: html, json, xml and text (default).
fuck.Format = "json" 

awesome, err := fuck.Awesome("Baphy")
if err != nil {
	log.Fatalln(err)
}

fmt.Println(awesome)
```
