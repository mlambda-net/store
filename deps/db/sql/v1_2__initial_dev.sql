

CREATE PROCEDURE ADD_PRODUCT
(
    _id uuid,
    _code VARCHAR(50),
    _name varchar(200),
    _currency VARCHAR(6),
    _brand VARCHAR(255),
    _category VARCHAR(255),
    _description varchar(500),
    _price decimal(10,2)
) AS $$
    INSERT INTO product(id, code, name, currency, brand, category, description, price)
    VALUES (_id, _code, _name, _currency, _brand, _category, _description, _price)
$$ LANGUAGE SQL;


CREATE  PROCEDURE EDIT_PRODUCT
(
    product_id uuid,
    product_name varchar(200),
    product_description varchar(500)
) AS $$

    UPDATE product SET
    name = product_name,
    description = product_description
    WHERE id = product_id;

$$ LANGUAGE SQL;



CREATE PROCEDURE ADD_PRODUCT_PREVIEW
(
    product_id uuid,
    image_id uuid,
    path varchar(500)
) AS $$

    INSERT INTO image(id, path)
    VALUES (image_id, path);

    UPDATE product
    SET thumbnail = image_id
    WHERE id = product_id;

$$ LANGUAGE SQL;


CREATE PROCEDURE ADD_PRODUCT_IMAGE
(
    product_id uuid,
    image_id uuid,
    path varchar(500)
) AS $$

    INSERT INTO image(id, path)
    VALUES (image_id, path);

    INSERT INTO product_image (product_id, image_id)
    VALUES (product_id, image_id);

$$ LANGUAGE SQL;

CREATE PROCEDURE DELETE_PRODUCT
(
    product_id uuid
) AS $$

    DELETE FROM product where id = product_id;

$$ LANGUAGE SQL;




CREATE PROCEDURE DELETE_IMAGES
(
    pid uuid
) AS  $$

    DELETE FROM image WHERE id IN (SELECT image_id FROM product_image where product_id = pid)

$$ LANGUAGE SQL;


CREATE PROCEDURE LOOKUP_ADD
(
    _name VARCHAR(255),
    _parent VARCHAR(255)
) AS  $$

    INSERT INTO lookup(name, parent)
    VALUES (_name, _parent)

$$ LANGUAGE SQL;

CREATE PROCEDURE LOOKUP_DELETE
(
    _name VARCHAR(255)
) AS  $$

    DELETE FROM lookup WHERE name = _name

$$ LANGUAGE SQL;
