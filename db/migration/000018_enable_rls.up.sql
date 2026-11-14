-- 000018_enable_rls.up.sql
-- Mengaktifkan Row Level Security (RLS) pada semua tabel publik

ALTER TABLE "users" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "directors" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "actors" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "genres" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "cites" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "payment_methods" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "subscribers" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "movies" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "cinemas" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "movie_genres" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "movie_casts" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "schedules" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "seats" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "orders" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "order_details" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "order_payments" ENABLE ROW LEVEL SECURITY;
ALTER TABLE "point_transactions" ENABLE ROW LEVEL SECURITY;

-- Catatan: Secara default, setelah RLS diaktifkan, semua akses dari API (anon/authenticated) akan ditolak
-- kecuali jika kebijakan (policy) ditambahkan. 
-- Karena backend Go biasanya menggunakan role 'postgres' atau superuser, backend tetap dapat mengakses data.
