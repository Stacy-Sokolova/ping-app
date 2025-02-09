CREATE TABLE containers
(
    id            serial       not null unique,
    port          text not null unique,
    timeofping  varchar(255),
    dateofping varchar(255)
);