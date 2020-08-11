create table customers (
	customerNumber int not null,
    customerName varchar(50) not null,
    phone varchar(50) not null,
    creditLimit decimal(10, 2) default null,
    primary key(customerNumber)
    )


create table orders (
    orderNumber int not null,
    orderDate date not null,
    shippedDate date default null,
    status varchar(15) not null,
    customerNumber int not null,
    primary key(orderNumber),
    constraint orders_cust_fk foreign key(customerNumber) references customers(customerNumber)
)

create table orderdetails (
    orderNumber int not null,
    productCode varchar(15) not null,
    quantityOrdered int not null,
    priceEach decimal(10,2) not null,
    orderLineNumber smallint not null,
    primary key(orderNumber, productCode),
    constraint odetails_order foreign key (orderNumber) references orders(orderNumber)
)

insert into zentools.customers (customerNumber, customerName, phone, creditLimit)
select customerNumber, customerName, phone, creditLimit from classicmodels.customers

insert into zentools.orders (orderNumber, orderDate, shippedDate, status, customerNumber)
select orderNumber, orderDate, shippedDate, status, customerNumber from classicmodels.orders

insert into zentools.orderdetails(orderNumber, productCode, 
                quantityOrdered, priceEach, orderLineNumber)
select orderNumber, productCode, quantityOrdered, 
                priceEach, orderLineNumber from classicmodels.orderdetails;


create table zentools.customers as
select customerNumber, customerName, phone, creditLimit from classicmodels.customers;

create table zentools.orders as
select orderNumber, orderDate, shippedDate, 
status, customerNumber from classicmodels.orders;

create table zentools.orderdetails as 
select * from classicmodels.orderdetails;