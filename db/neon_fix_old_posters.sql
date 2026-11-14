-- ============================================================================
-- Fix poster & backdrop URL untuk 10 film seed lama (ID 1-10)
-- Film lama pakai placeholder URL /poster1.jpg yang tidak exist.
-- Diganti dengan URL TMDB yang valid.
-- Jalankan di Neon SQL Editor.
-- ============================================================================

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/xDMIl84Qo5Tsu62c9DGWhmPI67A.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/qdIMHd4sEfJSckfVJfKQvisL02a.jpg'
WHERE id = 1; -- The Last Frontier → pakai poster Interstellar style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/d5iIlFn5s0ImszYzBPb8JPIfbXD.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/fRGxZuo7jJUWQsVg9PREb98Aclp.jpg'
WHERE id = 2; -- Echoes of Tomorrow → Dune style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/t9XkeE8UhNECXnVH9frbuSSzxEF.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/cyecClT0aTMBVibJ3rRjDQWRfaB.jpg'
WHERE id = 3; -- Midnight in Paris Revisited → Barbie style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/vZloFAK7NmvMGKE7VkF5UHaz0I.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/9n2tJBplPbgR2ca05hS5CKXwP2c.jpg'
WHERE id = 4; -- The Silent Witness → thriller style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/8b8R8l88Qje9dn9OE8PY05Nxl1X.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/odJ4hx6g6vBt4lBWKFD1tI8WS4x.jpg'
WHERE id = 5; -- Quantum Paradox → sci-fi style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/t2yyOv40HZeVlLjYsCsPHnWLk4W.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/4HodYYKEIsGOdinkGi2Ucz6X9i0.jpg'
WHERE id = 6; -- Ocean's Legacy → heist style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/loRmRzQXZeqG78TqZuyvSlEQfZb.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/wNAhuOZ3Zf84jCIlrcI6JhgmY5q.jpg'
WHERE id = 7; -- The Forgotten Kingdom → fantasy style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/gKkl37BQuKTanygYQG1pyYgLVgf.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/1t9T1DsLcABT3TqAP0Mx1teFRqH.jpg'
WHERE id = 8; -- Neon Dreams → cyberpunk style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/vB8o2p4ETnrfiWEgVxHmHWP9yRl.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/mSDsSDwaP3E7dEfUPWy4J0djt4O.jpg'
WHERE id = 9; -- Love in the Time of AI → romance style

UPDATE movies SET
  poster_url   = 'https://image.tmdb.org/t/p/w500/ghd5zOQnDaDW1mxVLHJfSbkibibz.jpg',
  backdrop_url = 'https://image.tmdb.org/t/p/original/TU9NIjwzjoKPwQHoHshkFcQUCG.jpg'
WHERE id = 10; -- The Final Symphony → drama style
