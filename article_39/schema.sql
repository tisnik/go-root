create table customers (
    ID            integer primary key asc,
    name          text not null,
    surname       text not null,
    address       text not null,
    country       text not null,
    phone         text not null
);

create table products (
    ID            integer primary key asc,
    name          text not null,
    description   text not null,
    price         number(6, 2) not null
);

create table orders (
    ID            integer primary key asc,
    customer_id   integer not null,
    sold          datetime not null,
    total         number(6, 2) not null,
    foreign key(customer_id) references customers(ID)
);

create table order_item (
    ID            integer primary key asc,
    order_id      integer not null,
    product_id    integer not null,
    quantity      integer not null,
    price         number(6,2) not null,
    foreign key(order_id) references orders(ID),
    foreign key(product_id) references products(ID)
);

