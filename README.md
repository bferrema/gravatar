# gravatar
A Go wrapper for Gravatar



## Install

```sh
go get -u github.com/brunom4ciel/gravatar
```

## Examples

### Profile URL
```go
url := gravatar.New("mail@example.org").URL()
```
### Avatar URL with parameters
```go
url := gravatar.New("mail@example.org").
		Size(200).
		Default(gravatar.NotFound).
		Rating(gravatar.Pg).
		AvatarURL()
```

### Avatar URL with default image
```go
url := gravatar.New("mail@example.org").
        DefaultURL("http://example.org/image.png").
        AvatarURL()
```

### Force default avatar
```go
url := gravatar.New("mail@example.org").ForceDefault(true).AvatarURL()
```

### JSON profile data URL
```go
url := gravatar.New("mail@example.org").JSONURL()
```

### JSON profile data URL with callback
```go
url := gravatar.New("mail@example.org").JSONURLCallback("alert")
```

### Parsed profile data
```go
profile, err := gravatar.New("mail@example.org").Profiles()
```

## License
MIT licensed. See the LICENSE file for details.
