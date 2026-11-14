-- ============================================================================
-- Tickitz Neon Database Verification Script
-- ============================================================================
-- Jalankan di Neon SQL Editor untuk cek apakah schema & data sudah sesuai.
-- Script ini TIDAK mengubah data — hanya membaca dan memverifikasi.
-- ============================================================================

-- ============================================================================
-- 1. CEK SEMUA TABEL ADA (17 tabel)
-- ============================================================================
SELECT '=== 1. CEK TABEL ===' AS info;

SELECT table_name, 
       (SELECT count(*) FROM information_schema.columns WHERE table_name = t.table_name AND table_schema = 'public') AS column_count
FROM information_schema.tables t
WHERE table_schema = 'public' 
  AND table_type = 'BASE TABLE'
ORDER BY table_name;

-- Expected: 17 tabel
-- users, directors, actors, genres, cities, payment_methods, subscribers,
-- movies, cinemas, movie_genres, movie_casts, schedules, seats,
-- orders, order_details, order_payments, point_transactions


-- ============================================================================
-- 2. CEK KOLOM PENTING (yang sering bermasalah)
-- ============================================================================
SELECT '=== 2. CEK KOLOM orders ===' AS info;

SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = 'orders' AND table_schema = 'public'
ORDER BY ordinal_position;

-- Expected: ada kolom payment_method (ditambahkan di migration 000019)
-- Kalau TIDAK ada, jalankan: ALTER TABLE public.orders ADD COLUMN IF NOT EXISTS payment_method character varying;


-- ============================================================================
-- 3. CEK JUMLAH DATA PER TABEL
-- ============================================================================
SELECT '=== 3. CEK JUMLAH DATA ===' AS info;

SELECT 'users' AS table_name, count(*) AS row_count FROM users
UNION ALL SELECT 'directors', count(*) FROM directors
UNION ALL SELECT 'actors', count(*) FROM actors
UNION ALL SELECT 'genres', count(*) FROM genres
UNION ALL SELECT 'cities', count(*) FROM cities
UNION ALL SELECT 'payment_methods', count(*) FROM payment_methods
UNION ALL SELECT 'subscribers', count(*) FROM subscribers
UNION ALL SELECT 'movies', count(*) FROM movies
UNION ALL SELECT 'cinemas', count(*) FROM cinemas
UNION ALL SELECT 'movie_genres', count(*) FROM movie_genres
UNION ALL SELECT 'movie_casts', count(*) FROM movie_casts
UNION ALL SELECT 'schedules', count(*) FROM schedules
UNION ALL SELECT 'seats', count(*) FROM seats
UNION ALL SELECT 'orders', count(*) FROM orders
UNION ALL SELECT 'order_details', count(*) FROM order_details
UNION ALL SELECT 'order_payments', count(*) FROM order_payments
UNION ALL SELECT 'point_transactions', count(*) FROM point_transactions
ORDER BY table_name;

-- Expected counts:
-- users: 10, directors: 10, actors: 10, genres: 10, cities: 10,
-- payment_methods: 10, subscribers: 10, movies: 10, cinemas: 10,
-- movie_genres: 27, movie_casts: 29, schedules: 10, seats: 60,
-- orders: 10, order_details: 25, order_payments: 10, point_transactions: 10


-- ============================================================================
-- 4. CEK USER ADMIN ADA
-- ============================================================================
SELECT '=== 4. CEK USER ADMIN ===' AS info;

SELECT id, email, first_name, last_name, role FROM users WHERE role = 'admin';

-- Expected: 1 row — admin@tickitz.com
-- Catatan: password seed user dummy, perlu register ulang via API untuk login


-- ============================================================================
-- 5. CEK FOREIGN KEY CONSTRAINTS
-- ============================================================================
SELECT '=== 5. CEK FOREIGN KEY ===' AS info;

SELECT
    tc.table_name,
    kcu.column_name,
    ccu.table_name AS references_table,
    ccu.column_name AS references_column
FROM information_schema.table_constraints tc
JOIN information_schema.key_column_usage kcu ON tc.constraint_name = kcu.constraint_name
JOIN information_schema.constraint_column_usage ccu ON tc.constraint_name = ccu.constraint_name
WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_schema = 'public'
ORDER BY tc.table_name;

-- Expected: foreign keys untuk movies→directors, cinemas→cities,
-- movie_genres→movies+genres, movie_casts→movies+actors,
-- schedules→movies+cinemas, seats→cinemas, orders→users+schedules,
-- order_details→orders+seats, order_payments→orders+payment_methods,
-- point_transactions→users


-- ============================================================================
-- 6. CEK DATA SAMPLE — MOVIES
-- ============================================================================
SELECT '=== 6. CEK MOVIES ===' AS info;

SELECT id, title, release_date, popularity_score, director_id
FROM movies
ORDER BY id
LIMIT 10;


-- ============================================================================
-- 7. CEK DATA SAMPLE — SCHEDULES + CINEMA + CITY
-- ============================================================================
SELECT '=== 7. CEK SCHEDULES ===' AS info;

SELECT s.id, m.title, c.name AS cinema, ci.name AS city, s.show_date, s.show_time, s.price
FROM schedules s
JOIN movies m ON s.movie_id = m.id
JOIN cinemas c ON s.cinema_id = c.id
JOIN cities ci ON c.city_id = ci.id
ORDER BY s.id;


-- ============================================================================
-- 8. CEK DATA SAMPLE — ORDERS + SEATS
-- ============================================================================
SELECT '=== 8. CEK ORDERS ===' AS info;

SELECT o.id, u.email, o.booking_code, o.total_price, o.payment_status, o.payment_method,
       count(od.id) AS ticket_count
FROM orders o
JOIN users u ON o.user_id = u.id
LEFT JOIN order_details od ON od.order_id = o.id
GROUP BY o.id, u.email
ORDER BY o.id;


-- ============================================================================
-- 9. CEK UNIQUE CONSTRAINTS
-- ============================================================================
SELECT '=== 9. CEK UNIQUE CONSTRAINTS ===' AS info;

SELECT tc.table_name, tc.constraint_name, kcu.column_name
FROM information_schema.table_constraints tc
JOIN information_schema.key_column_usage kcu ON tc.constraint_name = kcu.constraint_name
WHERE tc.constraint_type = 'UNIQUE' AND tc.table_schema = 'public'
ORDER BY tc.table_name;

-- Expected: users_email_key, genres_name_key, cities_name_key,
-- subscribers_email_key, orders_booking_code_key


-- ============================================================================
-- 10. CEK RLS (Row Level Security)
-- ============================================================================
SELECT '=== 10. CEK RLS STATUS ===' AS info;

SELECT relname AS table_name, relrowsecurity AS rls_enabled
FROM pg_class
WHERE relnamespace = 'public'::regnamespace
  AND relkind = 'r'
ORDER BY relname;

-- Catatan: RLS seharusnya DISABLE (false) untuk semua tabel.
-- Jika ada yang ENABLE tanpa policy, backend akan tidak bisa akses data.


-- ============================================================================
-- SELESAI
-- ============================================================================
-- Jika semua tabel ada, kolom payment_method ada di orders, dan
-- row count sesuai expected, maka database sudah siap dipakai.
-- ============================================================================
