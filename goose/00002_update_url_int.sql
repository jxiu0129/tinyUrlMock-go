-- +goose Up
ALTER TABLE `TinyUrlMock_go`.`Url` RENAME COLUMN `UID` TO `ID`;
-- +goose Down
