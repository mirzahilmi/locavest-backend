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
VALUES (1001, "Apple", "https://hips.hearstapps.com/hmg-prod/images/apples-royalty-free-image-164084111-1537885595.jpg", "Kilo", 30000),
    (1002, "Beef", "https://cdn.britannica.com/68/143268-050-917048EA/Beef-loin.jpg", "Kilo", 40000),
    (1003, "Kol", "https://cdn0-production-images-kly.akamaized.net/G1k3qqNZ_oAdQCphNYUKTi71ArE=/800x450/smart/filters:quality(75):strip_icc():format(webp)/kly-media-production/medias/2544528/original/078203300_1545144854-cabbage-847079_1920.jpg", "Piece", 32000);