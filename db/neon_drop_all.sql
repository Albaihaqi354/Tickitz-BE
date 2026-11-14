-- ============================================================================
-- Tickitz — Drop All Tables
-- ============================================================================
-- Jalankan ini HANYA jika ingin reset database Neon dari awal.
-- Setelah ini, jalankan neon_schema.sql untuk buat ulang semua tabel + data.
-- ============================================================================

DROP TABLE IF EXISTS public.point_transactions CASCADE;
DROP TABLE IF EXISTS public.order_payments CASCADE;
DROP TABLE IF EXISTS public.order_details CASCADE;
DROP TABLE IF EXISTS public.orders CASCADE;
DROP TABLE IF EXISTS public.seats CASCADE;
DROP TABLE IF EXISTS public.schedules CASCADE;
DROP TABLE IF EXISTS public.movie_casts CASCADE;
DROP TABLE IF EXISTS public.movie_genres CASCADE;
DROP TABLE IF EXISTS public.cinemas CASCADE;
DROP TABLE IF EXISTS public.movies CASCADE;
DROP TABLE IF EXISTS public.subscribers CASCADE;
DROP TABLE IF EXISTS public.payment_methods CASCADE;
DROP TABLE IF EXISTS public.cities CASCADE;
DROP TABLE IF EXISTS public.genres CASCADE;
DROP TABLE IF EXISTS public.actors CASCADE;
DROP TABLE IF EXISTS public.directors CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;

-- Drop sequences jika ada yang tersisa
DROP SEQUENCE IF EXISTS public.users_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.directors_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.actors_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.genres_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.cities_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.payment_methods_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.subscribers_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.movies_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.cinemas_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.schedules_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.seats_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.orders_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.order_details_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.order_payments_id_seq CASCADE;
DROP SEQUENCE IF EXISTS public.point_transactions_id_seq CASCADE;
