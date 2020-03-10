CREATE TABLE pizzas (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  price decimal(10,2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE option_types (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE options (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  type_id int NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_options_option_types FOREIGN KEY (type_id)
    REFERENCES option_types(id)
);

CREATE TABLE pizzas_options (
  id int NOT NULL AUTO_INCREMENT,
  pizza_id INT NOT NULL,
  option_id INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_pizzas_options_pizzas FOREIGN KEY (pizza_id)
    REFERENCES pizzas(id),
  CONSTRAINT fk_pizzas_options_options FOREIGN KEY (option_id)
    REFERENCES options(id)
);

-- Seeds

INSERT INTO pizzas (name, price) 
  VALUES 
    ("Super Cheesy Margherita", 90.00),
    ("Funghi", 75.00),
    ("Creamy Sucuk", 85.00);

INSERT INTO option_types (name) 
  VALUES 
    ("Topping"),
    ("Crust");

INSERT INTO options (name, type_id)
  VALUES
    ("Tomato sauce", 1),
    ("Mozarella", 1),
    ("Thin Crust", 2),
    ("Thick Crust", 2),
    ("Cheese-filled Crust", 2),
    ("Mushrooms", 1),
    ("Creme Fraiche", 1),
    ("Spinach", 1),
    ("Sucuk", 1),
    ("Red Onion", 1),
    ("Garlic Olive Oil Crust", 2);

INSERT INTO pizzas_options (pizza_id, option_id)
  VALUES
    (1, 1),
    (1, 2),
    (1, 5),
    (2, 1),
    (2, 2),
    (2, 6),
    (2, 3),
    (3, 7),
    (3, 8),
    (3, 9),
    (3, 10),
    (3, 11);