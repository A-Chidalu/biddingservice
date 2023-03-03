CREATE DATABASE BidDB;

USE BidDB;

CREATE TABLE BID (
     id INT NOT NULL AUTO_INCREMENT,
     user_id INT NOT NULL,
     item_id INT NOT NULL,
     amount DECIMAL(12, 4) NOT NULL,
     timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     is_terminating BOOLEAN NOT NULL DEFAULT FALSE,
     PRIMARY KEY (id)
);

INSERT INTO BID (user_id, item_id, amount, is_terminating) VALUES (1, 100, 10.00, false);
INSERT INTO BID (user_id, item_id, amount, is_terminating) VALUES (2, 101, 20.50, true);
INSERT INTO BID (user_id, item_id, amount, is_terminating) VALUES (3, 102, 5.25, false);
INSERT INTO BID (user_id, item_id, amount, is_terminating) VALUES (4, 103, 15.75, false);
INSERT INTO BID (user_id, item_id, amount, is_terminating) VALUES (5, 104, 30.00, true);


