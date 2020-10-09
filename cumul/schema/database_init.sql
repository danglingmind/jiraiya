-- schema for cumul API
create database cumul;
use cumul;

create table user
(
    userid varcharacter(5) not null primary key unique,
    created timestamp default current_timestamp,
    paid bool default false
);

create table urls(
	urlid int not null primary key auto_increment, 
    userid varcharacter(5) not null,
    created timestamp default current_timestamp,
    url TEXT not null,
    foreign key(userid) references user(userid)
    );