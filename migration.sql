-- Database: musicapi

-- DROP DATABASE musicapi;

CREATE DATABASE musicapi
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Kazakhstan.1251'
    LC_CTYPE = 'Russian_Kazakhstan.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


CREATE TABLE IF NOT EXISTS users(
	id int GENERATED ALWAYS AS IDENTITY,
	name varchar(100) UNIQUE,
	password varchar(100),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS musics(
	id int GENERATED ALWAYS AS IDENTITY,
	name varchar(100),
	artist varchar(100),
	url varchar(100),
	album varchar(100),
	length varchar(10),
	genre varchar(100),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS artists(
	id int GENERATED ALWAYS AS IDENTITY,
	nickname varchar(100) UNIQUE,
	fullname varchar(100),
	city varchar(100),
	birth_date varchar(30),
	PRIMARY KEY(id)
);

INSERT INTO users(name, password) VALUES('admin', '$2a$05$7cvNQL8sok3xpnnkpeO7aeAtWfmHiyDnSnkXITKqqql7eVN3hBPki'); --admin/admin
INSERT INTO musics(name, artist, url, album, length, genre) VALUES(
	'Out my head',
	'Fox Stevenson',
	'https://www.youtube.com/watch?v=xj3SZv5lle4',
	'killjoy',
	'03:46',
	'drum and bass'
),(
	'Go like',
	'Fox Stevenson',
	'https://www.youtube.com/watch?v=8GfxW3oS63I',
	'killjoy',
	'03:53',
	'drum and bass'
),(
	'Bruises',
	'Fox Stevenson',
	'https://www.youtube.com/watch?v=sFCpqrMMQGM',
	'Bruises Revisited',
	'03:40',
	'drum and bass'
),(
	'Worlds Collide',
	'Koven',
	'https://www.youtube.com/watch?v=x_j99YPKlnY',
	'Butterfly Effect',
	'03:52',
	'drum and bass'
),(
	'Never gonna give you up',
	'Rick Roll',
	'https://www.youtube.com/watch?v=dQw4w9WgXcQ',
	'Video Hits - Songs of The 80s',
	'03:32',
	'Rock'
);

INSERT INTO artists(nickname, fullname, city, birth_date) VALUES(
	'Koven',
	'Katie Boyle & Max Rowat',
	'London',
	'19-09-1982'
),(
	'Fox Stevenson',
	'Stanley Stevenson Byrne',
	'London',
	'25-01:1991'
);
