-- +goose Up
-- +goose StatementBegin
CREATE TABLE characters (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name STRING NOT NULL,
	role STRING NOT NULL,
	bounty INT
);

INSERT INTO characters (name, role, bounty) VALUES ('Monkey D. Luffy', 'Captain', 1500000000);
INSERT INTO characters (name, role, bounty) VALUES ('Roronoa Zoro', 'Swordsman', 320000000);
INSERT INTO characters (name, role, bounty) VALUES ('Nami', 'Navigator', 66000000);
INSERT INTO characters (name, role, bounty) VALUES ('Usopp', 'Sniper', 30000000);
INSERT INTO characters (name, role, bounty) VALUES ('Sanji', 'Cook', 77000000);
INSERT INTO characters (name, role, bounty) VALUES ('Tony Tony Chopper', 'Doctor', 1000000);
INSERT INTO characters (name, role, bounty) VALUES ('Nico Robin', 'Archaeologist', 130000000);
INSERT INTO characters (name, role, bounty) VALUES ('Franky', 'Shipwright', 44000000);
INSERT INTO characters (name, role, bounty) VALUES ('Brook', 'Musician', 33000000);
INSERT INTO characters (name, role, bounty) VALUES ('Jinbe', 'Helmsman', 40000000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE characters;
-- +goose StatementEnd
