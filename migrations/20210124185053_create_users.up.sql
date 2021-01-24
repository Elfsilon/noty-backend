CREATE TABLE users (
    ID bigserial NOT NULL PRIMARY KEY,
    name varchar(40) NOT NULL DEFAULT 'User',
    photo text DEFAULT NULL,
    email varchar(320) DEFAULT NULL,
    login varchar(60) NOT NULL UNIQUE,
    password varchar NOT NULL
);

CREATE TABLE categories (
    ID bigserial NOT NULL PRIMARY KEY,
    title varchar(100) DEFAULT 'Category name' NOT NULL,
    description text DEFAULT NULL,
    last_modification_datetime timestamp DEFAULT current_timestamp NOT NULL
);

CREATE TABLE documents (
    ID bigserial NOT NULL PRIMARY KEY,
    title varchar(100) DEFAULT 'Document name',
    description text DEFAULT NULL,
    last_modification_datetime timestamp DEFAULT current_timestamp NOT NULL,
    content text DEFAULT NULL
);

CREATE TABLE usr_cat (
    usr_ID bigserial NOT NULL,
    cat_ID bigserial NOT NULL,
    PRIMARY KEY (usr_ID, cat_ID),
    FOREIGN KEY (usr_ID)
        REFERENCES users (ID)
        ON DELETE CASCADE,
    FOREIGN KEY (cat_ID)
        REFERENCES categories (ID)
        ON DELETE CASCADE
);

CREATE TABLE cat_doc (
    cat_ID bigserial NOT NULL,
    doc_ID bigserial NOT NULL,
    PRIMARY KEY (cat_ID, doc_ID),
    FOREIGN KEY (cat_ID)
        REFERENCES categories (ID)
        ON DELETE CASCADE,
    FOREIGN KEY (doc_ID)
        REFERENCES documents (ID)
        ON DELETE CASCADE
);