insert into customers (id, name, surname, address, country, phone) values (0, 'Maria', 'Anders', 'Berlin', 'Germany', '030-0074321');
insert into customers (id, name, surname, address, country, phone) values (1, 'Ana', 'Trujillo', 'México D.F.', 'Mexico', '(5) 555-4729');
insert into customers (id, name, surname, address, country, phone) values (2, 'Antonio', 'Moreno', 'México D.F.', 'Mexico', '(5) 555-3932');
insert into customers (id, name, surname, address, country, phone) values (3, 'Thomas', 'Hardy', 'London', 'UK', '(171) 555-7788');
insert into customers (id, name, surname, address, country, phone) values (4, 'Christina', 'Berglund', 'Luleå', 'Sweden', '0921-12 34 65');

insert into products (id, name, description, price) values (0, 'Matice M5', 'DIN 934', 0.25);
insert into products (id, name, description, price) values (1, 'Matice M5 nízká', 'DIN 439', 0.15);
insert into products (id, name, description, price) values (2, 'Matice M5 dlouhá', 'DIN 6334', 1.50);
insert into products (id, name, description, price) values (3, 'Soustruh', 'S2-B', 10000.00);

insert into orders (id, customer_id, sold, total) values (0, 0, CURRENT_TIMESTAMP, 10025);
insert into orders (id, customer_id, sold, total) values (1, 1, CURRENT_TIMESTAMP, 150);

insert into order_item (id, order_id, product_id, quantity, price) values (0, 0, 0, 100, 25);
insert into order_item (id, order_id, product_id, quantity, price) values (1, 0, 3, 1, 10000);
insert into order_item (id, order_id, product_id, quantity, price) values (2, 1, 2, 100, 150);
