-- create users table


CREATE TABLE "users"(
    "id" BIGSERIAL NOT NULL UNIQUE,
    "name" VARCHAR(30) NOT NULL,
    "username" VARCHAR(50) DEFAULT NULL,
    "email" VARCHAR(50) NOT NULL,
    "email_verified_at" TIMESTAMPTZ DEFAULT NULL,
    "phone" VARCHAR(15) DEFAULT NULL,
    "phone_verified_at" TIMESTAMPTZ DEFAULT NULL,
    "password" VARCHAR(255) NOT NULL,
    "status" INTEGER DEFAULT 0,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

-- create index users
CREATE INDEX "users_id_index" ON
    "users"("id");
CREATE INDEX "users_username_index" ON
    "users"("username");
CREATE INDEX "users_email_index" ON
    "users"("email");
CREATE INDEX "users_phone_index" ON
    "users"("phone");
CREATE INDEX "users_status_index" ON
    "users"("status");
CREATE INDEX "users_deleted_at_index" ON
    "users"("deleted_at");
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_username_unique" UNIQUE("username");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE
    "users" ADD CONSTRAINT "users_phone_unique" UNIQUE("phone");

-- create authorizations
CREATE TABLE "authorizations"(
    "id" BIGSERIAL NOT NULL,
    "user_id" INTEGER NOT NULL,
    "provider" VARCHAR(20) NOT NULL,
    "uid" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ NULL
);
-- index
CREATE INDEX "authorizations_id_index" ON
    "authorizations"("id");
CREATE INDEX "authorizations_uid_index" ON
    "authorizations"("uid");
CREATE INDEX "authorizations_deleted_at_index" ON
    "authorizations"("deleted_at");
ALTER TABLE
    "authorizations" ADD PRIMARY KEY("id");
CREATE INDEX "authorizations_user_id_index" ON
    "authorizations"("user_id");

-- create organizations table
CREATE TABLE "organizations"(
    "id" BIGSERIAL NOT NULL,
    "parent_id" INTEGER NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "type" VARCHAR(20) CHECK
        ("type" IN('ORGANIZATION', 'BUSINESS')) NOT NULL,
        "name" VARCHAR(50) NOT NULL,
        "category_id" INTEGER NOT NULL,
        "legal_id" INTEGER NOT NULL,
        "branch" INTEGER DEFAULT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "organizations_id_index" ON
    "organizations"("id");
CREATE INDEX "organizations_parent_id_index" ON
    "organizations"("parent_id");
CREATE INDEX "organizations_owner_id_index" ON
    "organizations"("owner_id");
CREATE INDEX "organizations_type_index" ON
    "organizations"("type");
CREATE INDEX "organizations_deleted_at_index" ON
    "organizations"("deleted_at");
ALTER TABLE
    "organizations" ADD PRIMARY KEY("id");
CREATE INDEX "organizations_category_id_index" ON
    "organizations"("category_id");
CREATE INDEX "organizations_legal_id_index" ON
    "organizations"("legal_id");

-- create organization legals table
CREATE TABLE "organization_legals"(
    "id" BIGSERIAL NOT NULL,
    "slug" VARCHAR(50) NOT NULL,
    "value" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "organization_legals_id_index" ON
    "organization_legals"("id");
CREATE INDEX "organization_legals_deleted_at_index" ON
    "organization_legals"("deleted_at");
ALTER TABLE
    "organization_legals" ADD PRIMARY KEY("id");
CREATE INDEX "organization_legals_slug_index" ON
    "organization_legals"("slug");
CREATE INDEX "organization_legals_created_at_index" ON
    "organization_legals"("created_at");

-- creata categories table
CREATE TABLE "categories"(
    "id" BIGSERIAL NOT NULL,
    "slug" VARCHAR(75) NOT NULL,
    "value" VARCHAR(50) NOT NULL,
    "type" VARCHAR(20) CHECK
        ("type" IN('ORGANIZATION', 'PRODUCT')) NOT NULL,
        "parent_id" INTEGER DEFAULT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "categories_id_index" ON
    "categories"("id");
CREATE INDEX "categories_slug_index" ON
    "categories"("slug");
CREATE INDEX "categories_type_index" ON
    "categories"("type");
CREATE INDEX "categories_parent_id_index" ON
    "categories"("parent_id");
CREATE INDEX "categories_deleted_at_index" ON
    "categories"("deleted_at");
ALTER TABLE
    "categories" ADD PRIMARY KEY("id");


-- create products table
CREATE TABLE "products"(
    "id" BIGSERIAL NOT NULL,
    "business_id" INTEGER NOT NULL,
    "category_id" INTEGER NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "brand" VARCHAR(50) NOT NULL,
    "description" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "products_id_index" ON
    "products"("id");
CREATE INDEX "products_business_id_index" ON
    "products"("business_id");
CREATE INDEX "products_slug_index" ON
    "products"("slug");
CREATE INDEX "products_category_id_index" ON
    "products"("category_id");
CREATE INDEX "products_deleted_at_index" ON
    "products"("deleted_at");
ALTER TABLE
    "products" ADD PRIMARY KEY("id");
ALTER TABLE
    "products" ADD CONSTRAINT "products_slug_unique" UNIQUE("slug");

-- create contacts table
CREATE TABLE "contacts"(
    "id" BIGSERIAL NOT NULL,
    "parent_id" INTEGER NOT NULL,
    "type" VARCHAR(20) CHECK
        ("type" IN('USER', 'ORGANIZATION')) NOT NULL,
        "label" VARCHAR(255) NOT NULL,
        "value" VARCHAR(255) NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "contacs_id_index" ON
    "contacts"("id");
CREATE INDEX "contacs_parent_id_index" ON
    "contacts"("parent_id");
CREATE INDEX "contacs_type_index" ON
    "contacts"("type");
CREATE INDEX "contacs_deleted_at_index" ON
    "contacts"("deleted_at");
ALTER TABLE
    "contacts" ADD PRIMARY KEY("id");

-- create members table
CREATE TABLE "members"(
    "id" INTEGER NOT NULL,
    "parent_id" INTEGER NOT NULL,
    "organization_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "permission" VARCHAR(255) NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ DEFAULT NULL,
        "validated_at" TIMESTAMPTZ DEFAULT NULL,
        "validated_by" INTEGER DEFAULT NULL
);
-- index
CREATE INDEX "members_id_index" ON
    "members"("id");
CREATE INDEX "members_parent_id_index" ON
    "members"("parent_id");
CREATE INDEX "members_organization_id_index" ON
    "members"("organization_id");
CREATE INDEX "members_user_id_index" ON
    "members"("user_id");
CREATE INDEX "members_deleted_at_index" ON
    "members"("deleted_at");
ALTER TABLE
    "members" ADD PRIMARY KEY("id");

-- create assets table
CREATE TABLE "assets"(
    "id" INTEGER NOT NULL,
    "parent_id" INTEGER NOT NULL,
    "type" VARCHAR(255) CHECK
        ("type" IN('USER', 'ORGANIZATION', 'PRODUCT')) NOT NULL,
        "is_primary" BOOLEAN NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "location" VARCHAR(255) NOT NULL,
        "slug" VARCHAR(255) NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

-- index
CREATE INDEX "assets_id_index" ON
    "assets"("id");
CREATE INDEX "assets_parent_id_index" ON
    "assets"("parent_id");
CREATE INDEX "assets_type_index" ON
    "assets"("type");
CREATE INDEX "assets_slug_index" ON
    "assets"("slug");
CREATE INDEX "assets_deleted_at_index" ON
    "assets"("deleted_at");
ALTER TABLE
    "assets" ADD PRIMARY KEY("id");


-- create addresses table
CREATE TABLE "addresses"(
    "id" INTEGER NOT NULL,
    "parent_id" INTEGER NOT NULL,
    "is_primary" BOOLEAN NULL,
    "name" VARCHAR(255) NULL,
    "type" VARCHAR(255) CHECK ("type" IN('USER', 'ORGANIZATION')) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "district" VARCHAR(50) NOT NULL,
    "city" VARCHAR(50) NOT NULL,
    "province" VARCHAR(50) NOT NULL,
    "country" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL
);
-- index
CREATE INDEX "addresses_id_index" ON
    "addresses"("id");
CREATE INDEX "addresses_parent_id_index" ON
    "addresses"("parent_id");
CREATE INDEX "addresses_type_index" ON
    "addresses"("type");
CREATE INDEX "addresses_districtindex" ON
    "addresses"("district");
CREATE INDEX "addresses_cityindex" ON
    "addresses"("city");
CREATE INDEX "addresses_provinceindex" ON
    "addresses"("province");
CREATE INDEX "addresses_deleted_at_index" ON
    "addresses"("deleted_at");
ALTER TABLE
    "addresses" ADD PRIMARY KEY("id");

-- RELATION

ALTER TABLE
    "authorizations" ADD CONSTRAINT "authorizations_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "organizations" ADD CONSTRAINT "organizations_owner_id_foreign" FOREIGN KEY("owner_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "contacts" ADD CONSTRAINT "contacts_user_foreign" FOREIGN KEY("parent_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "members" ADD CONSTRAINT "members_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "organizations" ADD CONSTRAINT "organizations_legal_id_foreign" FOREIGN KEY("legal_id") REFERENCES "organization_legals"("id") ON DELETE CASCADE;
ALTER TABLE
    "organizations" ADD CONSTRAINT "organizations_organization_foreign" FOREIGN KEY("parent_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "contacts" ADD CONSTRAINT "contacts_organization_foreign" FOREIGN KEY("parent_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "assets" ADD CONSTRAINT "assets_organization_foreign" FOREIGN KEY("parent_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "members" ADD CONSTRAINT "members_organization_id_foreign" FOREIGN KEY("organization_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "organizations" ADD CONSTRAINT "organizations_category_id_foreign" FOREIGN KEY("category_id") REFERENCES "categories"("id") ON DELETE CASCADE;
ALTER TABLE
    "categories" ADD CONSTRAINT "categories_category_foreign" FOREIGN KEY("parent_id") REFERENCES "categories"("id") ON DELETE CASCADE;
ALTER TABLE
    "products" ADD CONSTRAINT "products_business_id_foreign" FOREIGN KEY("business_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "products" ADD CONSTRAINT "products_category_id_foreign" FOREIGN KEY("category_id") REFERENCES "categories"("id") ON DELETE CASCADE;
ALTER TABLE
    "members" ADD CONSTRAINT "members_member_foreign" FOREIGN KEY("parent_id") REFERENCES "members"("id") ON DELETE CASCADE;
ALTER TABLE
    "members" ADD CONSTRAINT "members_validated_by_foreign" FOREIGN KEY("validated_by") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "assets" ADD CONSTRAINT "assets_users_foreign" FOREIGN KEY("parent_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "assets" ADD CONSTRAINT "assets_product_foreign" FOREIGN KEY("parent_id") REFERENCES "products"("id") ON DELETE CASCADE;
ALTER TABLE
    "addresses" ADD CONSTRAINT "addresses_organization_id_foreign" FOREIGN KEY("parent_id") REFERENCES "organizations"("id") ON DELETE CASCADE;
ALTER TABLE
    "addresses" ADD CONSTRAINT "addresses_users_id_foreign" FOREIGN KEY("parent_id") REFERENCES "users"("id") ON DELETE CASCADE;