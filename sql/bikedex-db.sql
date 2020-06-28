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

INSERT INTO Brand("name", yearEst, country, websiteUrl) VALUES
('Cannondale', 1971, 'USA', 'https://www.cannondale.com/en-us'),
('Trek', 1975, 'USA', 'https://www.trekbikes.com/us/en_US/'),
('Fuji', 1899, 'Japan', 'https://www.fujibikes.com/usa/'),
('Bianchi', 1885, 'Italy', 'https://www.bianchi.com/'),
('Pinarello', 1952, 'Italy', 'https://pinarello.com/'),
('Specialized', 1974, 'USA', 'https://www.specialized.com/us/en'),
('Kona', 1988, 'USA', 'https://www.konaworld.com/');

INSERT INTO Color("name", "hex") VALUES
('Red', '#FF0000'),
('Blue', '#0000FF'),
('Yellow', '#FFFF00'),
('Black', '#000000'),
('White', '#FFFFFF'),
('Gray', '#808080'),
('Green', '#008000'),
('Purple', '#800080'),
('Orange', '#FFA500'),
('Silver', '#C0C0C0'),
('Turqouise', '#009998');

INSERT INTO "Type" ("name") VALUES
('Road'),
('Mountain'),
('Electric'),
('Active'),
('Kids');

INSERT INTO Subtype ("name") VALUES
('Competition'),
('Endurance'),
('Touring'),
('Cyclocross'),
('Trail'),
('All Mountain'),
('Cross Country'),
('Sport'),
('Electric'),
('Kids'),
('Cross Terrain'),
('Mountain'),
('Urban'),
('Fitness'),
('Cruiser');

INSERT INTO Model ("name", "version", msrp, brandId, typeId, subtypeId, manufacturerUrl) VALUES
('Cross', '1.3', 1299.99, 3, 1, 4, 'https://www.fujibikes.com/usa/bikes/road/cyclocross/cross/cross-1-3'),
('Chisel', 'Standard', 1375.00, 2, 2, 7, 'https://www.specialized.com/us/en/chisel/p/171237?color=264078-171237&searchText=91720-7001'),
('Treadwell Neo', 'Standard', 2300.00, 1, 3, 14, 'https://www.cannondale.com/en-us/bikes/electric/e-fitness/treadwell-neo/treadwell-neo?sku=c63150m10sm');

INSERT INTO Bike (modelId, colorId, imageUrl) VALUES
(1, 4, 'https://cdn.shopify.com/s/files/1/1109/6048/products/media_32d5c456-d7ac-43a8-b612-72eb88605b46_1024x1024.jpg?v=1589020416'),
(2, 4, 'https://s7d5.scene7.com/is/image/Specialized/?layer=0&wid=1920&hei=640&fmt=jpg&src=is{Specialized/pdp-product-bg-dark?wid=1920&hei=640}&layer=1&src=is{Specialized/91720-70_CHISEL-29-BLK-SUMBLU-HYP_HERO?wid=920&hei=600&$hybris-pdp-hero$}'),
(2, 11, 'https://s7d5.scene7.com/is/image/Specialized/?layer=0&wid=1920&hei=640&fmt=jpg&src=is{Specialized/pdp-product-bg-dark?wid=1920&hei=640}&layer=1&src=is{Specialized/91720-72_CHISEL-29-DSTTUR-DOVGRY_HERO?wid=920&hei=600&$hybris-pdp-hero$}'),
(3, 7, 'https://embed.widencdn.net/img/dorelrl/flycfmnisj/2000px@1x/C20_C63150M_Treadwell_Neo_MAT_PD.png');

