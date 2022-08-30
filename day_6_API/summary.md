# Advanced Golang - REST API

Representational State Transfer (REST) merupakan style arsitektural yang mendefinisikan satu set constraints yang digunakan untuk membuat layanan web. REST API adalah cara untuk mengakses layanan web dengan cara yang simple dan sederhana tanpa pemrosesan apapun.

## HTTP/HTTPS pada REST API

#### HTTP merupakan protokol untuk mendistribusikan, kolaborasidan hypermedia information system. HTTP sendiri hanya mendukung 1 arah, yang berarti jika kita mengirim request, harus menunggu respon terlebih dahulu sebelum mengirim request lagi. Sedangkan HTTP2 sudah mendukung duplex, sehingga kita bisa mengirim request terus-menerus tanpa menunggu respon terlebih dahulu.

#### Sedangkan HTTPS merupakan sistem protokol yang dienkripsi sehingga komunikasinya lebih aman. Protokol tersebut disebut TLS (Transport Layer Security), yang sebelumnya dikenal sebagai SSL (Secure Sockets Layer).

#### HTTP message/body/payload merupakan bagaimana pertukaran data antara server dan client terjadi. Ada dua tipe pesan: request melalui client untuk mentrigger server, dan response yang merupakan jawaban dari server.

#### HTTP header memungkinkan client dan server membypass informasi tambahan dengan sebuah HTTP request atau responses. HTTP header ditandai dengan colon(:).

##Method

### GET

```
HTTP GET http://www.example.com/users
HTTP GET http://www.example.com/users?size=20&page=5
HTTP GET http://www.example.com/users/123
HTTP GET http://www.example.com/users/123/address
```

### POST

```
HTTP POST http://www.example.com/users
HTTP POST http://www.example.com/users/123/account
```

### PUT (Update)

```
HTTP PUT http://www.example.com/users/123
HTTP PUT http://www.example.com/users/123/account/456
```

### PATCH (Partial update on a resource)

```
HTTP PATCH http://www.example.com/users/123
HTTP PATCH http://www.example.com/users/123/account/456
```

### DELETE

```
HTTP DELETE http://www.example.com/users/123
```

## Status Code

1. 2xx Success
   - 200 OK
   - 201 Created
   - 202 Accepted
2. 4xx Client error
   - 400 Bad Request
   - 401 Unauthorized
   - 403 Forbidden
   - 404 Not Found
   - 405 Method not found
3. 5xx Server error
   - 500 Internal server error
   - 502 Bad gateway
   - 504 Gateway timeout

# JSON (JavaScript Object Notation)

Merupakan format pertukaran data yang sangat ringan. Mudah dibaca dan ditulis oleh manusia. Mudah bagi device untuk menggabungkan dan menggenerate.

#### Code:

```
[
    {
        <key>: <value>
    }
]
```

- Key: `string`
- Value: `string`, `number`, `object`, `boolean`, `null`

Contoh:

```
[
    {
        "name": "Hideo Kojima",
        "email": "zhangpurnama@kojipro.jp",
        "age": 69,
        "is_married": true,
        "spouse_name": null
    },
    {
        "name": "Hideo Kojima",
        "email": "zhangpurnama@kojipro.jp",
        "age": 69,
        "is_married": true,
        "spouse": {
            "name": "bu kojima",
            "age": 50
        }
    },
    {
        "childs": [
            {
                "name": "Kojima Junior",
                "age": 30
            },
            {
                "name": "Kojima Junior MK.2",
                "age": 30
            }
        ]
    }
]
```

## Authorization REST API

### JSON Web Token (JWT)

```
<header>.<payload>.<signature>
```

`<header>` : Algorithm & Type token
`<payload>`: Second part of the token. Contain claims (typically, the user) and additional data.
`<signature>`: encoded header, encoded payload, secret
