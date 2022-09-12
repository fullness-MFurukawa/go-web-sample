CREATE DATABASE web_sample default character set utf8;
USE web_sample;

CREATE TABLE products (
    id int primary key auto_increment ,
    name varchar(30) not null ,
    price int not null ,
    register datetime DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO products (name,price,register) VALUES ('ボールペン 黒',150);
INSERT INTO products (name,price,register) VALUES ('ボールペン 赤',150);
INSERT INTO products (name,price,register) VALUES ('ボールペン 青',150);
INSERT INTO products (name,price,register) VALUES ('鉛筆 黒',100);
INSERT INTO products (name,price,register) VALUES ('鉛筆 赤',100);
INSERT INTO products (name,price,register) VALUES ('蛍光ペン 黄',200);
INSERT INTO products (name,price,register) VALUES ('蛍光ペン 赤',200);
INSERT INTO products (name,price,register) VALUES ('蛍光ペン 緑',200);
INSERT INTO products (name,price,register) VALUES ('蛍光ペン 青',200);