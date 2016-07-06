MercadoLibre's GO SDK
========

This is unofficial GO SDK for MercadoLibre's Platform.

## How to install ##
```go
go get github.com/tiaguinho/mercadolibre-go-sdk
```

## How do I use it? ##

The first thing to do is to instance a ```Meli``` class. You'll need to give a ```clientId``` and a ```clientSecret```. You can obtain both after creating your own application. For more information on this please read: [creating an application](http://developers.mercadolibre.com/application-manager/)

## Using the package ##

After import the package, you have to get new ```meli.Client``` struct

```go
client := meli.New(1234, "secret")
```

If you already have the access_token and refresh_token you can call another method for receive the same client

```go
client := meli.NewWithAccessToken(1234, "secret", "access_token", "refresh_token")
```

### Redirect users to authorize the application ###

First get the link to redirect the user.

```go
redirectUrl := meli.GetAuthUrl("redirect_url", meli.AuthUrls["site_code"])
```

You have to change de ```redirect_url``` for the url of your application and ```site_code``` for the code of the country you are implementing, for example: MLB, MLA, MCO (see the file auth.go to get the list).

Once the user is redirected to your ```redirect_url```, you'll receive in the query string, a parameter named ```code```. You'll need this value for authorize the app.

```go
client.Authorize(url.Query().Get("code"), "redirect_url")

```

This will get a ```access_token``` and ```refresh_token``` for your application and your user.

After that your are ready to make call to the API.

### Making GET calls ###

```go
params := map[string]string{"access_token": client.MLToken.AccessToken}

body, err := client.Get("users/me", params)
```

### Making POST calls ###

```go
params := map[string]string{"access_token": client.MLToken.AccessToken}

product := MLProduct{Foo: Bar}

body, err := client.Post("items", product, params)
```

### Making PUT calls ###

```go
params := map[string]string{"access_token": client.MLToken.AccessToken}

product := MLProduct{Foo: Bar2}

body, err := client.Put("items", product, params)
```

### Making DELETE calls ###

```go
params := map[string]string{"access_token": client.MLToken.AccessToken}

body, err := client.Delete("questions/123", params)
```

## Examples ##

For more examples, check out ```methods_test.go``` file.
