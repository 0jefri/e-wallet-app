CREATE DATABASE e_wallet;

CREATE TABLE users (
  Id varchar(100) primary key not null unique,
  Username varchar(100) not null unique,
  Email varchar(100) not null unique,
  Password varchar(100) not null,
  FirstName varchar(100) not null,
  LastName varchar(100) not null,
  PhoneNumber varchar(30) not null
);