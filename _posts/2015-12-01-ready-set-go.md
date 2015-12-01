---
layout: post
title:  "Ready, Set, Go"
date:   2015-12-01 05:00:00
categories: Golang Coding
---

As part of day one of Devcember I decided to instead start delving into learning Go. There is a site called [Starfighters](http://www.starfighters.io/) that is close to being released that will host a series of development related coding challenges. Since I am relatively sure that I will need an easy way to do various HTTP requests I have decided to spend some time this month working on a library that will abstract the process of sending and receiving JSON requests in Go.

The first part of the module that I have written is the ability to send/receive GET requests which returns an interface which contains the Unmarshaled JSON data. Here is the code for said module.

~~~
func Get(headers map[string]string, url string) (interface{}) {
    var responseBody interface{}

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    panicError(err)

    for key, val := range headers {
        req.Header.Add(key, val)
    }

    resp, err := client.Do(req)
    panicError(err)

    defer resp.Body.Close()
    rawBody, err := ioutil.ReadAll(resp.Body)
    panicError(err)

    err = json.Unmarshal(rawBody, &responseBody)
    panicError(err)

    return responseBody
}
~~~

As you can see, there are a few calls to panicError(err) spread throughout the code. When I was looking into writing Go code, I noticed that there was a common pattern of writing

~~~
if err != nil {
    panic(err)
}
~~~
in order to check the error status of a previous function call. Since dealing with HTTP requests can be a source of many errors, I figured that it would be a bit tedious to constantly have to copy/paste when I could just call a function instead. Then again, I am copy/pasting that function call anyways so...

There is also a main.go file which simply has some tests for the module. After only reading a few chapters of the new Go book I quickly realized that there is probably a standard way to write tests in Go so I will work my way through the Go book and once I reach the chapter that explains testing I will refactor the way that testing is done. If you would like to view more of the code, the repositiory is hosted [here](https://github.com/sylvesterwillis/gorest).  

I am also still reading through The Go Programming Language book, and will be writing more about it in the future. 