-- Create Operations

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, email, password, role_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id,first_name, last_name, phone_number, email, role_id, created_at, updated_at;

-- name: CreateRole :exec
INSERT INTO roles (name)
VALUES ($1);

-- name: CreateCategory :exec
INSERT INTO categories (name, description)
VALUES ($1, $2);

-- name: CreateSubcategory :exec
INSERT INTO subcategories (name, description, category_id)
VALUES ($1, $2, $3);

-- name: CreateWorkday :exec
INSERT INTO workdays (service_id, name, start_time, end_time, max_clients,day_id, open_to_work)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateService :exec
INSERT INTO services (user_id, description, google_map_address, willaya, baladia, subcategory_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetReservationsCount :one
SELECT  COUNT(*) FROM reservations WHERE service_id = $1 AND time = $2 AND weekday_id = $3;
-- name: GetReservationsCountForUpdate :one
SELECT COUNT(*)
FROM reservations
WHERE weekday_id = $1
  AND service_id = $2
  AND time = $3
FOR UPDATE;

-- name: GetReservationsCountByUserIdAndServiceId :one
SELECT  COUNT(*) FROM reservations WHERE user_id = $1 AND service_id = $2 AND time = $3 AND weekday_id = $4;
-- name: CreateReservation :one
INSERT INTO reservations (service_id, user_id, time, weekday_id, ranking, reserve_type)
VALUES ($1, $2, $3, $4, $5, 1)
RETURNING id, service_id, user_id, time, weekday_id, ranking, reserve_type, created_at, updated_at;

-- name: CreateRating :exec
INSERT INTO ratings (service_id, user_id, rating, comment)
VALUES ($1, $2, $3, $4);

-- name: CreateComplaint :exec
INSERT INTO complaints (user_id, type_id, complaint)
VALUES ($1, $2, $3);

-- name: CreateDay :exec
INSERT INTO days (name) VALUES ($1);

-- Read Operations

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT users.*, roles.name AS role_name
FROM users
JOIN roles ON users.role_id = roles.id
WHERE users.email = $1 LIMIT 1;

-- name: GetUserIDByEmail :one
SELECT id FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT users.*, roles.name AS role_name
FROM users
JOIN roles ON users.role_id = roles.id
WHERE users.email = $1 AND users.password = $2
LIMIT 1;


-- name: GetUsers :many
SELECT * FROM users;

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = $1 LIMIT 1;

-- name: GetRoles :many
SELECT * FROM roles;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetSubcategoryByID :one
SELECT * FROM subcategories WHERE id = $1 LIMIT 1;
-- name: GetServiceByUserID :many
SELECT * FROM services WHERE user_id = $1 LIMIT 1;
-- name: GetSubcategories :many
SELECT subcategories.*, categories.name AS category_name
FROM subcategories
JOIN categories ON subcategories.category_id = categories.id;



-- name: GetServiceByID :one
SELECT * FROM services WHERE id = $1 LIMIT 1;

-- name: GetServices :many
SELECT * FROM services ORDER BY id DESC OFFSET $2 LIMIT $1;

-- name: GetReservationsByUserID :many
SELECT * FROM reservations WHERE user_id = $1 OFFSET $3 LIMIT $2;

-- name: GetReservationsByServiceID :many
SELECT * FROM reservations WHERE service_id = $1 And OFFSET $3 LIMIT $2;

-- name: GetRatingsByServiceID :many
SELECT * FROM ratings WHERE service_id = $1;

-- name: GetComplaintsByUserID :many
SELECT * FROM complaints WHERE user_id = $1;

-- Update Operations

-- name: UpdateUserByID :exec
UPDATE users
SET first_name = $1, last_name = $2, phone_number = $3, email = $4, password = $5, role_id = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $7;

-- name: UpdateRoleByID :exec
UPDATE roles
SET name = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateCategoryByID :exec
UPDATE categories
SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3;

-- name: UpdateSubcategoryByID :exec
UPDATE subcategories
SET name = $1, description = $2, category_id = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4;



-- name: UpdateServiceByID :exec
UPDATE services
SET user_id = $1, description = $2, google_map_address = $3, willaya = $4, baladia = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

-- name: UpdateReservationByID :exec
UPDATE reservations
SET service_id = $1, user_id = $2, time = $3, weekday_id = $4, ranking = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

-- name: UpdateRatingByID :exec
UPDATE ratings
SET service_id = $1, user_id = $2, rating = $3, comment = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $5;


-- Delete Operations

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteRoleByID :exec
DELETE FROM roles
WHERE id = $1;
-- name: DelteAllRoles :exec
DELETE FROM roles;

-- name: DeleteCategoryByID :exec
DELETE FROM categories
WHERE id = $1;

-- name: DeleteSubcategoryByID :exec
DELETE FROM subcategories
WHERE id = $1;

-- name: DeleteWorkdayByID :exec
DELETE FROM workdays
WHERE id = $1;

-- name: DeleteServiceByID :exec
DELETE FROM services
WHERE id = $1;

-- name: DeleteReservationByID :exec
DELETE FROM reservations
WHERE id = $1;

-- name: DeleteRatingByID :exec
DELETE FROM ratings
WHERE id = $1;

-- name: DeleteComplaintByID :exec
DELETE FROM complaints
WHERE id = $1;

-- name: CreateReserveType :one

INSERT INTO reserve_types (name) VALUES ($1) RETURNING *;

-- name: GetReserveTypes :many
SELECT * FROM reserve_types;

-- name: DeleteReserveTypeByID :exec
DELETE FROM reserve_types WHERE id = $1;


-- name: GetDaysOfWorkByServiceID :many
SELECT workdays.* FROM workdays JOIN services ON workdays.service_id = services.id WHERE services.id = $1;


-- name: GetAllDays :many
SELECT * FROM days ORDER BY id;

-- name: GetWorkdays :many
SELECT * FROM workdays;

-- name: GetWorkdaysByServiceID :many
SELECT * FROM workdays WHERE service_id = $1 ORDER BY id;

-- name: GetWorksdayByID :many
SELECT * FROM workdays WHERE id = $1;

-- name: GetReservationsByWeekdayID :many
SELECT * FROM reservations
WHERE created_at::date = $1::date;

-- name: GetWorkdaysInRange :many
SELECT * FROM workdays
WHERE service_id = $1;

-- name: UpdateWorkdayByID :exec
UPDATE workdays
SET start_time = $1, end_time = $2, max_clients = $3, updated_at = CURRENT_TIMESTAMP, open_to_work = $4
WHERE id = $5;


-- name: UpdateAvergaeRating :one
UPDATE services
    SET average_rating = (SELECT AVG(rating) FROM ratings WHERE service_id = $1)
    WHERE id = $1 RETURNING average_rating;

-- name: GetAllReserveStatus :many
SELECT * FROM reservations_status;

-- name: CreateReserveStatus :exec
INSERT INTO reservations_status (name) VALUES ($1);

-- name: UpdateReserveStatusName :exec
UPDATE reservations_status SET name = $1 WHERE id = $2;

-- name: SearchServicesByCategory :many
SELECT s.*
FROM services s
JOIN subcategories sc ON s.subcategory_id = sc.id
JOIN categories c ON sc.category_id = $1 LIMIT $2 OFFSET $3;

-- name: SearchServicesBySubCategory :many
SELECT s.* FROM services AS s JOIN subcategories AS sc ON s.subcategory_id = $1 LIMIT $2 OFFSET $3;
-- WHERE sc.name ILIKE $1 OR sc.description ILIKE $1
-- OFFSET $3 LIMIT $2;
-- name: OrderServicesByRating :many
SELECT s.* FROM services AS s JOIN subcategories AS sc ON s.subcategory_id = $1 ORDER BY s.average_rating DESC LIMIT $2 OFFSET $3;

-- name: OrderServicesByDistance :many
SELECT 
    s.*,
    ( 
      6371 * acos( 
          cos(radians($4)) * 
          cos(radians(s.latitude)) * 
          cos(radians(s.longitude) - radians($5)) + 
          sin(radians($4)) * 
          sin(radians(s.latitude)) 
      ) 
    ) AS distance
FROM services AS s
JOIN subcategories AS sc ON s.subcategory_id = sc.id
WHERE s.subcategory_id = $1
ORDER BY distance ASC
LIMIT $2 OFFSET $3;


-- name: UpdateReservationStatusByID :one
UPDATE reservations SET reserv_status = $2 WHERE id = $1 RETURNING *;


-- name: GetReservationInfoByID :one
SELECT service_id, weekday_id 
FROM reservations 
WHERE id = $1;

-- name: GetNextUserReservations :many
SELECT * 
FROM reservations 
WHERE service_id = $1 
  AND weekday_id = $2 
  AND id > $3 
ORDER BY id ASC 
LIMIT $4;

-- name: CreateComplaintType :exec
INSERT INTO complaint_types (name) VALUES ($1);

-- name: GetComplaintTypes :many
SELECT * FROM complaint_types;

-- name: UpdateComplaintType :exec
UPDATE complaint_types SET name = $1 WHERE id = $2;

-- name: GetAllComplaints :many
SELECT * FROM complaints OFFSET $1 LIMIT $2;

-- name: GetComplaintByID :one
SELECT * FROM complaints WHERE id = $1;

-- name: GetAllComplaintTypes :many
SELECT * FROM complaint_types;

-- name: GetAllRatingByServiceID :many
SELECT * FROM ratings WHERE service_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: DeleteRating :exec
DELETE FROM ratings WHERE id = $1 AND user_id = $2;

-- name: CreateDeleteAccountRequest :exec
INSERT INTO delete_requests (user_id) VALUES ($1);

-- name: GetReservationByID :one
SELECT * FROM reservations WHERE id = $1;