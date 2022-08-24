# Day 3: Tugas PostgreSQL

## membuat table

### customers

```
create table customers (id int not null primary key, customer_name char(50) not null);
```

### product

```
create table products (id int primary key not null, name varchar(50) not null);
```

### orders

```
create table orders (id int PRIMARY KEY NOT NULL, customer_id int not null references customers(id), product_id INT NOT NULL references products(id), order_date date NOT NULL, total float NOT NULL);
```

## operasi

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

## Join

```
select * from public.customers join public.orders on orders.id = customers.id;S
```
