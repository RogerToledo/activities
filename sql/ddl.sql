-- DROP TABLE public.todos;

CREATE TABLE public.todos (
	id serial4 NOT NULL,
	"name" varchar NULL,
	description text NULL,
	done bool NULL,
	CONSTRAINT todos_pkey PRIMARY KEY (id)
);