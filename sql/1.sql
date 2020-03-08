drop database if exists tokens;

create database tokens;

use tokens;

drop table if exists token;

create table token (
    tokenId int not null auto_increment,
    token varchar(512) not null,
    createdDate datetime not null,
    expirationDate datetime,
    PRIMARY KEY (tokenId),
    unique key (token)
);

-- for unit tests
insert into token (token, createdDate) values ("1234567", NOW());

drop table if exists used;

create table used (
    usedId int not null auto_increment,
    path varchar(512) not null,
    usedDate datetime,
    PRIMARY KEY (usedId)
);

CREATE USER 'golanguser'@'localhost' IDENTIFIED BY 'golangpassword';

grant ALL on tokens.* to 'golanguser'@'localhost' IDENTIFIED BY 'golangpassword';