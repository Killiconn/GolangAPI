--This file is for the setup of MYSQL Schema for the application

CREATE TABLE credentials (
    Username VARCHAR(255),
    Email VARCHAR(255),
    Encrypt_pasword VARCHAR(255),
    PRIMARY KEY (Username)
);


CREATE TABLE fbreports (
    Username VARCHAR(255),
    Feedback VARCHAR(255),
    FBDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE feedbacks (
    Username VARCHAR(255),
    Feedback VARCHAR(255),
);