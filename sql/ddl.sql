CREATE TABLE category
(
    id          SERIAL
        CONSTRAINT category_pk
            PRIMARY KEY,
    name        VARCHAR(100) NOT NULL
        CONSTRAINT category_unique
            UNIQUE,
    image       VARCHAR(100) NOT NULL,
    description VARCHAR(100)
);

CREATE TABLE customer
(
    id         SERIAL
        CONSTRAINT customer_pk
            PRIMARY KEY,
    email      VARCHAR(100) NOT NULL
        UNIQUE,
    password   VARCHAR(100) NOT NULL,
    phone      VARCHAR(14)  NOT NULL
        UNIQUE,
    first_name VARCHAR(50)  NOT NULL,
    last_name  VARCHAR(50)  NOT NULL,
    created_at TIMESTAMP    NOT NULL
);

CREATE TABLE supplier
(
    id           SERIAL
        CONSTRAINT supplier_pk
            PRIMARY KEY,
    name         VARCHAR(100) NOT NULL
        CONSTRAINT supplier_unique
            UNIQUE,
    image        VARCHAR(100) NOT NULL,
    description  VARCHAR(500),
    time_opening VARCHAR(5)   NOT NULL,
    time_closing VARCHAR(5)   NOT NULL
);

CREATE TABLE product
(
    id          SERIAL
        CONSTRAINT product_pk
            PRIMARY KEY,
    supplier_id INTEGER      NOT NULL
        CONSTRAINT product_fk_supplier
            REFERENCES supplier
            ON DELETE CASCADE,
    category_id INTEGER      NOT NULL
        CONSTRAINT product_fk_category
            REFERENCES category
            ON DELETE CASCADE,
    name        VARCHAR(100) NOT NULL,
    image       VARCHAR(100) NOT NULL,
    description VARCHAR(200),
    price       REAL         NOT NULL
);

CREATE TABLE supplier_category
(
    supplier_id INTEGER NOT NULL
        REFERENCES supplier,
    category_id INTEGER NOT NULL
        REFERENCES category,
    PRIMARY KEY (supplier_id, category_id)
);

CREATE TABLE order_supplier
(
    order_id    INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL
        REFERENCES supplier,
    PRIMARY KEY (order_id, supplier_id)
);

CREATE TABLE "order"
(
    id                  SERIAL
        CONSTRAINT order_pk
            PRIMARY KEY,
    customer_id         INTEGER      NOT NULL
        CONSTRAINT order_fk_customer
            REFERENCES customer,
    recipient_full_name VARCHAR(100) NOT NULL,
    address             VARCHAR(100) NOT NULL,
    price               REAL         NOT NULL,
    created_at          TIMESTAMP    NOT NULL
);

CREATE TABLE order_product
(
    order_id         INTEGER NOT NULL
        REFERENCES "order",
    product_id       INTEGER NOT NULL
        REFERENCES product,
    product_quantity INTEGER NOT NULL,
    PRIMARY KEY (order_id, product_id),
    UNIQUE (order_id, product_id, product_quantity)
);

