/***
    This is the schema defining how TASKS and ROOMMATES are stored
***/


CREATE TABLE `tasks` (
    `tid` INT(10) NOT NULL AUTO_INCREMENT,
    `task_name` VARCHAR(40) NOT NULL,
    `description` VARCHAR(140) NULL DEFAULT NULL,
    `assigned_roommate` JSON,
    `possible_roommates` JSON
)