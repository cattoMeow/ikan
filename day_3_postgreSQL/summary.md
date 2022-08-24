# Day 3: Tugas PostgreSQL

Untuk tugas hari ini, saya mendapatkan tugas untuk membuat 3 tabel, melakukan operasi, serta melakukan join.

## Membuat table

disini kita akan membuat 3 buah tabel: customers, products, dan orders.

### customers

```
create table customers (id int not null primary key, customer_name char(50) not null);
```

### products

```
create table products (id int primary key not null, name varchar(50) not null);
```

### orders

```
create table orders (id int PRIMARY KEY NOT NULL, customer_id int not null references customers(id), product_id INT NOT NULL references products(id), order_date date NOT NULL, total float NOT NULL);
```

## Operasi

Setelah tabel dibuat, kita bisa melakukan operasi berikut untuk menambah, mengupdate, dan menghapus data yang ada dalam tabel.

### insert

```
insert into public.customers (customer_name) values ('Zhang Purnama');
```

### update

```
update customers set customer_name = 'Hideo Kojima' where id = 1;
```

### delete

```
delete from public.customers where id = 3;
```

## Joins

Join merupakan suatu langkah untuk menggabungkan tabel yang ada dengan parameter-parameter yang sesuai kebutuhan

### Join

```
select * from public.customers join public.orders on orders.id = customers.id;S
```

output yang dihasilkan akan menampilkan hasil gabungan tabel customer serta barang yang dibeli. Sehingga memudahkan pelacakan pembeli serta harga yang harus dibayarkan.
