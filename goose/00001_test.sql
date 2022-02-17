-- +goose Up

-- -----------------------------------------------------
-- Table `tinyUrlMock_go`.`UsedKeys`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tinyUrlMock_go`.`UsedKeys` (
  `UniqueKey` VARCHAR(6) NOT NULL UNIQUE)
ENGINE = InnoDB;


-- -- -----------------------------------------------------
-- -- Table `tinyUrlMock_go`.`UnusedKeys`
-- -- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tinyUrlMock_go`.`UnusedKeys` (
  `UniqueKey` VARCHAR(6) NOT NULL UNIQUE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tinyUrlMock_go`.`UsedKeys`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tinyUrlMock_go`.`Url` (
  `UID` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `ShortenUrl` VARCHAR(6) NOT NULL,
  `OriginalUrl` VARCHAR(100) NOT NULL,
  `CreatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   `UpdatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   `DeletedAt` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`UID`))
ENGINE = InnoDB;


-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
