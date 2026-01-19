CREATE TABLE public.movie_casts (
    movie_id integer NOT NULL,
    actor_id integer NOT NULL
);

ALTER TABLE ONLY public.movie_casts
    ADD CONSTRAINT movie_casts_pkey PRIMARY KEY (movie_id, actor_id);

ALTER TABLE ONLY public.movie_casts
    ADD CONSTRAINT movie_casts_actor_id_fkey FOREIGN KEY (actor_id) REFERENCES public.actors(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.movie_casts
    ADD CONSTRAINT movie_casts_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;