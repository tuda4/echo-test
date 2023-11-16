CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "uuid" varchar(255) NOT NULL ,
                         "first_name" varchar(255),
                         "last_name" varchar(255),
                         "birthday" timestamp,
                         "address" varchar(255),
                         "phone" varchar(20),
                         "created_at" timestamptz NOT NULL DEFAULT 'now()',
                         "updated_at" timestamptz NOT NULL DEFAULT 'now()',
                         "deleted_at" timestamptz
);

CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "uuid" varchar(255) UNIQUE NOT NULL,
                            "email" varchar(255) NOT NULL,
                            "hash_password" varchar(255) NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT 'now()',
                            "updated_at" timestamptz NOT NULL DEFAULT 'now()',
                            "deleted_at" timestamptz
);

CREATE TABLE "sessions" (
                           "id" bigserial PRIMARY KEY,
                           "uuid" varchar(255) NOT NULL,
                           "refresh_token" varchar NOT NULL,
                           "user_agent" varchar NOT NULL,
                           "client_id" varchar NOT NULL,
                           "is_blocked" boolean NOT NULL DEFAULT false,
                           "expired_at" timestamptz NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "users" ADD FOREIGN KEY ("uuid") REFERENCES "accounts" ("uuid");

ALTER TABLE "sessions" ADD FOREIGN KEY ("uuid") REFERENCES "accounts" ("uuid");
