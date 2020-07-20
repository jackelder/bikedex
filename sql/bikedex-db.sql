DROP TABLE IF EXISTS Brand CASCADE;
CREATE TABLE Brand (
	"id" serial PRIMARY KEY,
	"name" VARCHAR(50),
	yearEst INT,
	country VARCHAR(50),
	websiteUrl VARCHAR(1000)
);

DROP TABLE IF EXISTS Color CASCADE;
CREATE TABLE Color (
	"id" serial PRIMARY KEY,
	"name" VARCHAR(50),
	"hex" CHAR(7)
);

DROP TABLE IF EXISTS "Type" CASCADE;
CREATE TABLE "Type" (
	"id" serial PRIMARY KEY,
	"name" VARCHAR(50)
);

DROP TABLE IF EXISTS Subtype CASCADE;
CREATE TABLE Subtype (
	"id" serial PRIMARY KEY,
	"name" VARCHAR(50)
);

DROP TABLE IF EXISTS Model CASCADE;
CREATE TABLE Model (
	"id" serial PRIMARY KEY,
	"name" VARCHAR(50),
	"version" VARCHAR(50),
	msrp DECIMAL(8,2),
	manufacturerUrl VARCHAR(1000),
	brandId int REFERENCES Brand(id),
	typeId int REFERENCES "Type"(id),
	subtypeId int REFERENCES Subtype(id)
);

DROP TABLE IF EXISTS Bike CASCADE;
CREATE TABLE Bike (
	"id" serial PRIMARY KEY,
	imageUrl VARCHAR(1000),
	modelId int REFERENCES Model(id),
	colorId int REFERENCES Color(id)
);
