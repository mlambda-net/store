
CREATE TABLE image
(
    id uuid NOT NULL,
    path VARCHAR(500)
);

ALTER TABLE image
ADD CONSTRAINT PK_Image
PRIMARY KEY (id);

CREATE TABLE product
(
    id uuid NOT NULL,
    code varchar(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    currency VARCHAR(6) NOT NULL,
    category VARCHAR(255) NOT NULL,
    brand VARCHAR(255),
    description VARCHAR(500),
    price DECIMAL(10,2),
    thumbnail uuid
);

ALTER TABLE product
ADD CONSTRAINT PK_Product
PRIMARY KEY (id);

ALTER TABLE product
ADD CONSTRAINT FK_Product_Thumbnail
FOREIGN KEY (thumbnail)
REFERENCES image(id)
ON DELETE CASCADE;

CREATE UNIQUE INDEX product_unique_code
ON product(code);

CREATE TABLE product_image
(
    product_id uuid,
    image_id uuid
);

ALTER TABLE product_image
ADD CONSTRAINT PK_Product_Images
PRIMARY KEY (product_id, image_id);

ALTER TABLE product_image
ADD CONSTRAINT FK_ProductImages_Product
FOREIGN KEY (product_id)
REFERENCES product(id);

ALTER TABLE product_image
ADD CONSTRAINT FK_ProductImages_Images
FOREIGN KEY (image_id)
REFERENCES image(id)
ON DELETE CASCADE;

CREATE TABLE lookup
(
    name varchar(255) not null,
    parent varchar(255)
);

ALTER TABLE lookup
ADD CONSTRAINT PK_lookup
PRIMARY KEY (name);

ALTER TABLE lookup
ADD CONSTRAINT FK_lookup_parent
FOREIGN KEY (parent)
REFERENCES lookup(name);


