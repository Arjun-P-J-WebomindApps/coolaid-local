-------------------------------------------------------------------------------------------------------------------------------
--Query 
-------------------------------------------------------------------------------------------------------------------------------

-- name: GetAllCustomers :many
SELECT *
FROM customers;

-- name: GetCustomersFromName :many
SELECT *
FROM customers
WHERE customer_company_name=$1;

-- name: GetCustomerFromID :one
SELECT *
FROM customers
WHERE id=$1;

-- name: GetCustomerByMobileNumber :one
SELECT *
FROM customers
WHERE mobile=$1;

-- name: GetCustomerByMobileCompanyContactPersonAndCustomerType :one
SELECT *
FROM customers
WHERE mobile = @Mobile::text
  AND customer_company_name = @CustomerCompanyName::text
  AND contact_person = @ContactPerson::text
  AND type = @CustomerType::text;

-- name: SearchCustomersByAttributes :many
SELECT *
FROM customers
WHERE 
    (
        NULLIF(@Customer_Company_Name::text,'') IS NULL OR 
        LOWER(customer_company_name) ILIKE '%' || LOWER(@Customer_Company_Name::text) || '%'
    )
  OR
    (   
        NULLIF(@Contact_Person::text,'') IS NULL OR 
        LOWER(contact_person) ILIKE '%' || LOWER(@Contact_Person::text) || '%'
    )
  OR
    (
        NULLIF(@Mobile::text ,'')IS NULL OR
        mobile ILIKE '%' || @Mobile::text || '%'
    )
  ORDER BY customer_company_name ASC;

-----------------------------------------------------------------------------------------------------------------------------
--MUTATION
-----------------------------------------------------------------------------------------------------------------------------


-- name: CreateCustomer :one
INSERT INTO customers
    (id,customer_company_name,contact_person,mobile,type,customer_designation,address,flat,street,city,state,pincode,payment_mode,created_at,updated_at,deleted_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING *;


-- name: UpdateCustomerByID :one
UPDATE customers AS c SET
  customer_company_name = COALESCE(sqlc.narg('customer_company_name'), c.customer_company_name),
  contact_person        = COALESCE(sqlc.narg('contact_person'), c.contact_person),
  mobile                = COALESCE(sqlc.narg('mobile'), c.mobile),
  type                  = COALESCE(sqlc.narg('type'), c.type),
  customer_designation  = COALESCE(sqlc.narg('customer_designation'), c.customer_designation),
  address               = COALESCE(sqlc.narg('address'), c.address),
  flat                  = COALESCE(sqlc.narg('flat'), c.flat),
  street                = COALESCE(sqlc.narg('street'), c.street),
  city                  = COALESCE(sqlc.narg('city'), c.city),
  state                 = COALESCE(sqlc.narg('state'), c.state),
  pincode               = COALESCE(sqlc.narg('pincode'), c.pincode),
  payment_mode          = COALESCE(sqlc.narg('payment_mode'), c.payment_mode),
  updated_at            = NOW()
WHERE c.id = $1
RETURNING *;


-- name: DeleteCustomerByID :one
DELETE FROM customers
WHERE id = @ID
RETURNING *;

