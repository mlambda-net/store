
INSERT INTO lookup(name) VALUES ('root');
INSERT INTO lookup(name, parent) VALUES ('Category', 'root');
INSERT INTO lookup(name, parent) VALUES ('Brand', 'root');
INSERT INTO lookup(name, parent) VALUES ('Technology', 'Category');
INSERT INTO lookup(name, parent) VALUES ('Tea', 'Category');
INSERT INTO lookup(name, parent) VALUES ('Kitchen', 'Category');
INSERT INTO lookup(name, parent) VALUES ('Home', 'Category');
INSERT INTO lookup(name, parent) VALUES ('None', 'Brand');

