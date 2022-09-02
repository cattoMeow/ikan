## Echo Framework

### Mengapa menggunakan Echo Framework?

Performa tinggi, extensible, framework web pada GO yang minimalis.

#### Kelebihan:

1. Optimized Router
2. Scalable
3. Automatic TLS
4. HTTP/2
5. Middleware
6. Data Binding
7. Data rendering
8. Templates
9. Extensible

## Unit Testing

Merupakan sebuah tahapan dalam software development yaitu pengujian salah satu bagian terkecil disebuah aplikasi, disebut units, yaitu operasi yang dilakukan secara individual dan independent.

### Unit testing pada Goalng

- File unit testing harus diakhiri dengan nama `_testing.go` pada package yang sama
- Unit testing pada Golang ditulis dengan bentuk fungsi, yang memiliki parameter tipe `*testing.T`, dengan nama dung si yang dimulai dengan kata `Test`

## Logging

Merupakan rekaman event yag terjadi pada sistem operasi atau sebuah software atau pesan antar user pada aplikasi komunikasi. Logging adalah tindakan untuk menyimpan log. Pada kasus sederhana, pesan-pesan tersebut ditulis pada satu file log.

## Clean Architecture

Clean arch have three main component named delivery, service, and repository.
| Nama | Kegunaan | Contoh |
| ----- | -------|------|
| Delivery | Delivery is where we can put all tech stack that we use and consume by client in one bucket. | The example of tech stack is REST API, GraphQL, and SOA |
| Service | Service is where we can store our business logic. | For example, code that can perform get user.|
| Repository | It’s a place where we can store 3rd party integration code | 3rd party API of rajaongkir, MySQL, NoSQL, cache.|
