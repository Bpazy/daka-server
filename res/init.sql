CREATE TABLE `info` (
    `INFO_ID`     INTEGER PRIMARY KEY AUTOINCREMENT,
    `USERNAME`    VARCHAR(64) NOT NULL,
    `DISTANCE`    VARCHAR(64) NOT NULL,
    `DATE`        DATE        NOT NULL,
    `CREATE_TIME` DATE        NULL    DEFAULT CURRENT_TIMESTAMP,
    `UDPATE_TIME` DATE        NULL    DEFAULT CURRENT_TIMESTAMP
);