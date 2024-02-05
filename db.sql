CREATE DATABASE `ecobricks`;

use `ecobricks`;

CREATE TABLE `categories` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `ecobricks`.`categories` (`id`,`name`) VALUES ("6b4c28f4-6831-495a-9444-19c93452faa3","Ecológicos");
INSERT INTO `ecobricks`.`categories` (`id`,`name`) VALUES ("7c0ca0d4-ff23-4bd7-b131-c563067c4b43","Padrão");

CREATE TABLE `products` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `category_id` varchar(36) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_categories_idx` (`category_id`)
);

INSERT INTO `ecobricks`.`products` (`id`,`name`,`description`,`price`,`category_id`,`image_url`)
VALUES ("7f8c9d8e-9f0a-1b2c-3d4e-5f6g7h8i9j0k","Tijolo Maciço","Tijolo do tipo Maciço", 100, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/1.png"),
("66805003-f9a2-4772-b577-d5ff66207707","Tijolo Laminado","Tijolo do tipo Laminado", 200, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/2.png"),
("121829b9-e9f9-4ca9-bd14-b087afedd587","Tijolo Refratário","Tijolo do tipo Refratário", 300, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/3.png"),
("ef3d9a49-4c73-4192-97dd-55e21c0e2707","Tijolo Baiano","Tijolo do tipo Baiano", 400, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/4.png"),
("6d602b3a-1e72-4b03-a29c-14095e57027b","Lajota","Tijolo do tipo Lajota", 500, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/5.png"),
("dad6f8fb-3149-4d0b-861e-03d29c6accf0","Bloco decorativo","Tijolo do tipo  Bloco Decorativo", 600, "7c0ca0d4-ff23-4bd7-b131-c563067c4b43", "http://localhost:9000/products/6.png"),
("61c176d5-f4f9-4fbd-a892-41058422868e","Tijolo Ecológico (1 furo)","Tijolo do tipo Ecológico, com 1 furo", 700, "6b4c28f4-6831-495a-9444-19c93452faa3", "http://localhost:9000/products/7.png"),
("ed394062-14bc-4ff2-bf43-a77473322171","Tijolo Ecológico (2 furos)","Tijolo do tipo Ecológico, com 2 furos", 800, "6b4c28f4-6831-495a-9444-19c93452faa3", "http://localhost:9000/products/8.png");