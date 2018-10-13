# URL Minifier

## Service description

This service is analog of [сlck.ru](https://clck.ru/). He generates short url from the specified, and redirect to full url when someone requested short url.
This service is a test project.

### Language and libares

Language: Go.
Librares: no.

### TODO

- hash algorithm collisions
- scaling
- remote url storage (resolve scaling problem)
- url swelling
- redirect chains (redict abuse problem)
- logs, metrics

## HTTP API

### Minificate

Generates short url from the specified one.

```http
GET http://localhost:8080/--?url=https://google.ru
```

Response:

```http
200 OK

localhost:8080/39ac065e
```

### Redirect

Redirect to full url.

```http
GET http://localhost:8080/39ac065e
```

Response redirects to the full url.

## Local launch

1. Clone the repository.
2. Launch the project.

```bash
git clone https://github.com/lillilli/url_minificator.git && cd url_minificator
make run
```

### Docker

1. Clone the repository.
2. Make image.
3. Launch image.

```bash
git clone https://github.com/lillilli/url_minificator.git && cd url_minificator
make image
make run:image
```
