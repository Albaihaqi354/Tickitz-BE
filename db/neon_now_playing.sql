-- ============================================================================
-- Tickitz — Now Playing Movies (Real Data dari TMDB)
-- ============================================================================
-- 10 film yang sedang tayang dengan poster, backdrop, director, dan cast
-- dari The Movie Database (TMDB).
--
-- Cara pakai: jalankan di Neon SQL Editor.
-- Data ini ditambahkan ke schema yang sudah ada (tidak menghapus data lama).
-- ============================================================================

-- ============================================================================
-- 1. TAMBAH GENRE BARU (Mystery tidak ada di seed awal)
-- ============================================================================
INSERT INTO "genres" ("name") VALUES ('Mystery') ON CONFLICT (name) DO NOTHING;

-- Genre IDs setelah insert:
-- 1=Action, 2=Drama, 3=Comedy, 4=Sci-Fi, 5=Thriller, 6=Horror, 7=Romance,
-- 8=Adventure, 9=Fantasy, 10=Crime, 11=Mystery


-- ============================================================================
-- 2. TAMBAH DIRECTORS BARU (ID 11-20)
-- ============================================================================
INSERT INTO "directors" ("name") VALUES
('Zach Cregger'),
('David Ayer'),
('Na Hong-jin'),
('Rob Savage'),
('Tazbah Chavez'),
('Carles Torrens'),
('Takashi Yamazaki'),
('Mike P. Nelson'),
('Sidharta Tata'),
('Nitesh Tiwari');


-- ============================================================================
-- 3. TAMBAH ACTORS BARU (ID 11-40)
-- ============================================================================
INSERT INTO "actors" ("name") VALUES
('Austin Abrams'),
('Paul Walter Hauser'),
('Kali Reis'),
('Brad Pitt'),
('J.K. Simmons'),
('Anna Lambe'),
('Hwang Jung-min'),
('Zo In-sung'),
('Hoyeon'),
('Jessica Chastain'),
('Arabella Olivia Clark'),
('Jay Duplass'),
('Kiawentiio'),
('Isabel DeRoy-Olson'),
('Michael Podemski'),
('Carlos Alcaide'),
('Julia Martinez'),
('Jose Maria Yazpik'),
('Ryunosuke Kamiki'),
('Minami Hamabe'),
('Min Tanaka'),
('Kiana Madeira'),
('Jessica Rothe'),
('Will Close'),
('Lukman Sardi'),
('Callista Arum'),
('Sulthan Hamonangan'),
('Ranbir Kapoor'),
('Yash'),
('Sai Pallavi');


-- ============================================================================
-- 4. INSERT MOVIES (ID 11-20) — Now Playing dari TMDB
-- ============================================================================
-- Poster URL: https://image.tmdb.org/t/p/w500{path}
-- Backdrop URL: https://image.tmdb.org/t/p/original{path}
-- ============================================================================

INSERT INTO "movies" ("title", "synopsis", "duration", "release_date", "director_id", "poster_url", "backdrop_url", "popularity_score") VALUES
-- 11. Resident Evil
('Resident Evil',
 'Medical courier Bryan unwittingly finds himself fighting for survival as one fateful, horrifying night collapses around him in chaos.',
 94, '2026-09-16', 11,
 'https://image.tmdb.org/t/p/w500/gAgWJhUSQWCgtKR8170nMkYlBda.jpg',
 'https://image.tmdb.org/t/p/original/8R79AZjX1Bt4219Ub3CrRMsPQsW.jpg',
 550.76),

-- 12. Heart of the Beast
('Heart of the Beast',
 'After a harrowing plane crash, Special Forces officer James Belmont and his combat dog, Odin, find themselves stranded deep in the Alaskan wilderness. Together, they are forced into a brutal fight for survival against the elements.',
 101, '2026-09-19', 12,
 'https://image.tmdb.org/t/p/w500/6sYemVUn79UPKz1ZsSG6kPU4RVy.jpg',
 'https://image.tmdb.org/t/p/original/6uVRPFflyJpo8Cx872olmlGTWla.jpg',
 540.24),

-- 13. Hope
('Hope',
 'In the remote South Korean village of Hope Harbor, police chief Bum-seok and officer Sung-ae are called to investigate a mysterious creature wreaking havoc on the community. But all is not as it seems, and perceptions can be misleading. What begins as ignorance plants the seed of disaster, escalating through human conflict into a tragedy of cosmic proportions.',
 157, '2026-07-15', 13,
 'https://image.tmdb.org/t/p/w500/6VY1ZKgpsc8CfGsQcCGiYEgHdH0.jpg',
 'https://image.tmdb.org/t/p/original/j96rurZqmBjyY37Fncl3IB0qJNe.jpg',
 284.38),

-- 14. Other Mommy
('Other Mommy',
 'Young Bela forms a relationship with a sinister entity in her home that looks uncannily like her own mother, drawing her into a terrifying and increasingly dark world.',
 93, '2026-10-07', 14,
 'https://image.tmdb.org/t/p/w500/fGRNCssTRrhYZxNzykGGCoMsAVq.jpg',
 'https://image.tmdb.org/t/p/original/ncm7djsVTHQQMX2HEbEVxMUFchG.jpg',
 282.49),

-- 15. A Story of Derry
('A Story of Derry',
 'Hundreds of years before Pennywise, three youngsters venture into the Western Wood...where they encounter a terrifying ancient spirit.',
 30, '2026-10-31', 15,
 'https://image.tmdb.org/t/p/w500/4PtzGcegF5Po1NmDAW9IklyvAUh.jpg',
 'https://image.tmdb.org/t/p/original/z26LxXDMmAIzKr06FPDudCJyvsM.jpg',
 267.74),

-- 16. Apocalypse Z: New World
('Apocalypse Z: New World',
 'The Canary Islands have become Europe''s last refuge following the outbreak of the TSJ virus. When a special mission sets out for the mainland in search of vital supplies, Roberto joins the team with questionable intentions. The dangerous operation turns into a desperate race for survival, and Roberto and his companions find themselves caught between hordes of infected and divided loyalties.',
 96, '2026-10-09', 16,
 'https://image.tmdb.org/t/p/w500/wXln4DUrwcyjcFPwgQxFPYTlG1Z.jpg',
 'https://image.tmdb.org/t/p/original/o3XwuDUj0n8qnK91p6q8LKfr6rj.jpg',
 250.08),

-- 17. Godzilla Minus Zero
('Godzilla Minus Zero',
 'War reduced Japan to zero, and Godzilla plunged it into minus. Two years have passed since then, and the country has faced agonizing struggles to achieve recovery and finally reclaim its daily life. Just as humanity managed to cling to a hard-won peace, a new threat strikes, crushing hope in an instant. There is no third time. Everything ends here.',
 135, '2026-11-03', 17,
 'https://image.tmdb.org/t/p/w500/hSmIJluOJmTWvJs4d7kHRxjc8s8.jpg',
 'https://image.tmdb.org/t/p/original/oqrYIDH4tvfEzPe1yqT4jzmAWHQ.jpg',
 173.67),

-- 18. Beware Boiuna
('Beware Boiuna',
 'A medical mission deep in the Amazon becomes a desperate fight for survival when a colossal ancient predator awakens, turning the jungle into a deadly battleground where nature decides who lives and dies.',
 103, '2026-10-01', 18,
 'https://image.tmdb.org/t/p/w500/7mI13sTZcuwWEIUK9fckAb8Zmba.jpg',
 'https://image.tmdb.org/t/p/original/9nU0FKNefmhbgi2xXlvZchZsqwm.jpg',
 181.04),

-- 19. Never Surrender
('Never Surrender',
 'When a deadly gang descends on his remote home, a guilt-ridden veteran and the estranged grandchildren he barely knows must unite to survive.',
 140, '2026-11-11', 19,
 'https://image.tmdb.org/t/p/w500/hKROmLddpzpjkxKKVFYFRAOXX1b.jpg',
 'https://image.tmdb.org/t/p/original/umxN2kMcUm6VoxSOdcEcQoXuC4q.jpg',
 150.10),

-- 20. Ramayana: Rise of a Legend
('Ramayana: Rise of a Legend',
 'An ancient epic follows a young prince and princess whose marriage and subsequent exile mark the beginning of a legendary journey. Their story of love, duty and destiny sets the stage for greater challenges ahead.',
 201, '2026-11-05', 20,
 'https://image.tmdb.org/t/p/w500/vBQ7wks68OhJKlb3RnXSl2yXBXK.jpg',
 'https://image.tmdb.org/t/p/original/yHVZncPRZr63mezRX20IyKLrSkv.jpg',
 147.23);


-- ============================================================================
-- 5. INSERT MOVIE_GENRES
-- ============================================================================
-- Movie IDs: 11=Resident Evil, 12=Heart of the Beast, 13=Hope,
-- 14=Other Mommy, 15=A Story of Derry, 16=Apocalypse Z, 17=Godzilla,
-- 18=Beware Boiuna, 19=Never Surrender, 20=Ramayana
-- Genre IDs: 1=Action, 4=Sci-Fi, 5=Thriller, 6=Horror, 8=Adventure,
-- 9=Fantasy, 10=Crime, 11=Mystery, 2=Drama

INSERT INTO "movie_genres" ("movie_id", "genre_id") VALUES
-- Resident Evil: Horror, Sci-Fi, Adventure
(11, 6), (11, 4), (11, 8),
-- Heart of the Beast: Adventure, Thriller
(12, 8), (12, 5),
-- Hope: Sci-Fi, Mystery, Action
(13, 4), (13, 11), (13, 1),
-- Other Mommy: Horror
(14, 6),
-- A Story of Derry: Drama, Horror, Mystery
(15, 2), (15, 6), (15, 11),
-- Apocalypse Z: Action, Thriller, Horror
(16, 1), (16, 5), (16, 6),
-- Godzilla Minus Zero: Sci-Fi, Horror, Action
(17, 4), (17, 6), (17, 1),
-- Beware Boiuna: Horror, Thriller, Action
(18, 6), (18, 5), (18, 1),
-- Never Surrender: Action, Crime, Thriller
(19, 1), (19, 10), (19, 5),
-- Ramayana: Fantasy, Action, Adventure
(20, 9), (20, 1), (20, 8);


-- ============================================================================
-- 6. INSERT MOVIE_CASTS
-- ============================================================================
-- Actor IDs: 11-40 (lihat section 3 untuk mapping nama)

INSERT INTO "movie_casts" ("movie_id", "actor_id") VALUES
-- Resident Evil: Austin Abrams(11), Paul Walter Hauser(12), Kali Reis(13)
(11, 11), (11, 12), (11, 13),
-- Heart of the Beast: Brad Pitt(14), J.K. Simmons(15), Anna Lambe(16)
(12, 14), (12, 15), (12, 16),
-- Hope: Hwang Jung-min(17), Zo In-sung(18), Hoyeon(19)
(13, 17), (13, 18), (13, 19),
-- Other Mommy: Jessica Chastain(20), Arabella Olivia Clark(21), Jay Duplass(22)
(14, 20), (14, 21), (14, 22),
-- A Story of Derry: Kiawentiio(23), Isabel DeRoy-Olson(24), Michael Podemski(25)
(15, 23), (15, 24), (15, 25),
-- Apocalypse Z: Carlos Alcaide(26), Julia Martinez(27), Jose Maria Yazpik(28)
(16, 26), (16, 27), (16, 28),
-- Godzilla Minus Zero: Ryunosuke Kamiki(29), Minami Hamabe(30), Min Tanaka(31)
(17, 29), (17, 30), (17, 31),
-- Beware Boiuna: Kiana Madeira(32), Jessica Rothe(33), Will Close(34)
(18, 32), (18, 33), (18, 34),
-- Never Surrender: Lukman Sardi(35), Callista Arum(36), Sulthan Hamonangan(37)
(19, 35), (19, 36), (19, 37),
-- Ramayana: Ranbir Kapoor(38), Yash(39), Sai Pallavi(40)
(20, 38), (20, 39), (20, 40);


-- ============================================================================
-- 7. INSERT SCHEDULES (gunakan cinema yang sudah ada)
-- ============================================================================
INSERT INTO "schedules" ("movie_id", "cinema_id", "show_date", "show_time", "price") VALUES
-- Resident Evil di CGV Grand Indonesia (cinema 1)
(11, 1, '2026-11-14', '14:00:00', 50000),
(11, 1, '2026-11-14', '19:30:00', 55000),
-- Heart of the Beast di Cinepolis Plaza Senayan (cinema 2)
(12, 2, '2026-11-14', '16:00:00', 50000),
(12, 2, '2026-11-15', '20:00:00', 55000),
-- Hope di XXI Dago (cinema 3)
(13, 3, '2026-11-14', '18:00:00', 45000),
-- Other Mommy di CGV Paris Van Java (cinema 4)
(14, 4, '2026-11-14', '21:00:00', 60000),
-- A Story of Derry di Cinepolis Tunjungan Plaza (cinema 5)
(15, 5, '2026-11-15', '15:00:00', 45000),
-- Apocalypse Z di XXI Hartono Mall (cinema 6)
(16, 6, '2026-11-14', '17:00:00', 50000),
(16, 6, '2026-11-15', '20:30:00', 55000),
-- Godzilla Minus Zero di CGV Sun Plaza (cinema 7)
(17, 7, '2026-11-14', '19:00:00', 75000),
(17, 7, '2026-11-15', '14:00:00', 65000),
-- Beware Boiuna di Cinepolis Ciputra Semarang (cinema 8)
(18, 8, '2026-11-14', '16:30:00', 50000),
-- Never Surrender di XXI Panakkukang (cinema 9)
(19, 9, '2026-11-15', '18:30:00', 55000),
-- Ramayana di CGV Beachwalk (cinema 10)
(20, 10, '2026-11-14', '15:30:00', 70000),
(20, 10, '2026-11-15', '19:00:00', 75000);


-- ============================================================================
-- SELESAI
-- ============================================================================
-- 10 film now playing sudah ditambahkan dengan:
-- - Poster & backdrop dari TMDB
-- - Director & cast asli dari TMDB
-- - Genre mapping ke database
-- - Schedules di 10 cinema
--
-- Total data baru: 1 genre, 10 directors, 30 actors, 10 movies,
-- 27 movie_genres, 30 movie_casts, 16 schedules
-- ============================================================================
