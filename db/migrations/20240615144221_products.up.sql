CREATE TABLE `Products` (
    `ID` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `Name` VARCHAR(255),
    `Image` VARCHAR(255),
    `Format` VARCHAR(255),
    `Price` DECIMAL(15, 2),
    `CreatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `UpdatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);
INSERT INTO Products (ID, Name, Image, Format, Price)
VALUES (1001, "Apple", "", "Kilo", 30000),
    (1002, "Beef", "", "Kilo", 40000),
    (1003, "Kol", "", "Piece", 32000);