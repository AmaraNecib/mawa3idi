-- إنشاء جدول الأدوار
CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);
-- إنشاء جدول المستخدمين
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(10) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    role_id INT NOT NULL,
    -- date of creation like this 2024-07-17 13:18:36z
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);


-- إنشاء جدول الأصناف
CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);

-- إنشاء جدول الأصناف الفرعية
CREATE TABLE IF NOT EXISTS subcategories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);

-- إنشاء جدول الخدمات
CREATE TABLE IF NOT EXISTS services (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    description TEXT NOT NULL,
    subcategory_id INT NOT NULL,
    google_map_address VARCHAR(255) NOT NULL,
    willaya VARCHAR(50) NOT NULL,
    baladia VARCHAR(50) NOT NULL,
    average_rating FLOAT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (subcategory_id) REFERENCES subcategories(id)
);
CREATE TABLE IF NOT EXISTS days (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);
-- إنشاء جدول أيام الأسبوع
CREATE TABLE IF NOT EXISTS workdays (
    id BIGSERIAL PRIMARY KEY,
    service_id INT NOT NULL,
    name VARCHAR(20) NOT NULL,
    open_to_work BOOLEAN DEFAULT FALSE NOT NULL,
    day_id INT NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    max_clients INT NOT NULL,
    FOREIGN KEY (service_id) REFERENCES services(id),
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (day_id) REFERENCES days(id)
);
-- reserve type
CREATE TABLE IF NOT EXISTS reserve_types (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);

-- reservations status
CREATE TABLE IF NOT EXISTS reservations_status (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);

-- إنشاء جدول الحجوزات
CREATE TABLE IF NOT EXISTS reservations (
    id BIGSERIAL PRIMARY KEY,
    service_id INT NOT NULL,
    user_id INT NOT NULL,
    time DATE NOT NULL,
    weekday_id INT NOT NULL, 
    ranking INT DEFAULT 0 NOT NULL,
    reserve_type INT DEFAULT 1 NOT NULL,
    reserv_status INT DEFAULT 1 NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (service_id) REFERENCES services(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (weekday_id) REFERENCES workdays(id),
    FOREIGN KEY (reserve_type) REFERENCES reserve_types(id),
    FOREIGN KEY (reserv_status) REFERENCES reservations_status(id)
);


-- table for rating services
CREATE TABLE IF NOT EXISTS ratings (
    id BIGSERIAL PRIMARY KEY,
    service_id INT NOT NULL,
    user_id INT NOT NULL,
    rating INT NOT NULL,
    comment TEXT,
    FOREIGN KEY (service_id) REFERENCES services(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);
-- table for complaints types
CREATE TABLE IF NOT EXISTS complaint_types (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);
-- table for complaints
CREATE TABLE IF NOT EXISTS complaints (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    type_id INT NOT NULL,
    complaint TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (type_id) REFERENCES complaint_types(id),
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW()
);

-- devicess token
CREATE TABLE IF NOT EXISTS devices (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table for delete account requests
CREATE TABLE IF NOT EXISTS delete_requests (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    updated_at TIMESTAMP(0) DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);