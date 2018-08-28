-- Adminer 4.6.3 PostgreSQL dump

\connect "cereals";

DROP TABLE IF EXISTS "geekers";
DROP SEQUENCE IF EXISTS geekers_id_seq;
CREATE SEQUENCE geekers_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."geekers" (
    "id" integer DEFAULT nextval('geekers_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "alias" text,
    "slack_id" text,
    "slack_name" text,
    CONSTRAINT "geekers_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_geekers_deleted_at" ON "public"."geekers" USING btree ("deleted_at");


DROP TABLE IF EXISTS "menu_items";
DROP SEQUENCE IF EXISTS menu_items_id_seq;
CREATE SEQUENCE menu_items_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."menu_items" (
    "id" integer DEFAULT nextval('menu_items_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "food_name" text,
    "price" integer,
    "for_date" integer,
    CONSTRAINT "menu_items_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_menu_items_deleted_at" ON "public"."menu_items" USING btree ("deleted_at");


DROP TABLE IF EXISTS "receipts";
DROP SEQUENCE IF EXISTS receipts_id_seq;
CREATE SEQUENCE receipts_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."receipts" (
    "id" integer DEFAULT nextval('receipts_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "geeker_id" integer,
    "menu_item_id" integer,
    "for_date" integer,
    CONSTRAINT "receipts_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "receipts_geeker_id_fkey" FOREIGN KEY (geeker_id) REFERENCES geekers(id) NOT DEFERRABLE,
    CONSTRAINT "receipts_menu_item_id_fkey" FOREIGN KEY (menu_item_id) REFERENCES menu_items(id) NOT DEFERRABLE
) WITH (oids = false);

CREATE INDEX "idx_receipts_deleted_at" ON "public"."receipts" USING btree ("deleted_at");


-- 2018-08-28 06:10:15.427467+00
