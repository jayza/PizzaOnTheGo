CREATE TABLE Pizza (
  ID int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  Name varchar(255) NOT NULL,
  Price int NOT NULL
);

INSERT INTO Pizza (Name, Price) VALUES ("Margherita", 9000);