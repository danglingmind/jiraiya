-- schema for cumul API
create database cumul;
use cumul;

create table user(
	userid varcharacter(100) not null primary key unique,
    password varchar(100) not null, 
	created timestamp default current_timestamp,
    paid bool default false
);


create table urls(
	urlid int not null primary key auto_increment, 
    userid varcharacter(50) not null,
    created timestamp default current_timestamp,
    url TEXT not null,
    urlname varchar(100) not null,
    foreign key (userid) references user(userid)
    );
    
    
drop table urls;
drop table user;


insert into user (userid, paid) values('abcdd',false);
select * from user;

select * from user;
select urlname, url from urls where userid = 'hari' group by urlname, url;



desc user;
desc urls;

drop table urls;
drop table user;
