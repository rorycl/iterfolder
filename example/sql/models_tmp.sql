
DROP TABLE IF EXISTS manu_model;

CREATE temp TABLE manu_model (
    manufacturer TEXT
    ,model TEXT
);

INSERT INTO manu_model VALUES
    ('Acura','MDX')
    ,('Chevrolet','Bolt EV')
    ,('Chevrolet','Equinox')
    ,('Chevrolet','Traverse')
    ,('GMC','Hummer EV')
    ,('Honda','Civic')
    ,('Hyundai','Tucson')
    ,('Jeep','Grand Wagoneer')
    ,('Kia','Sorento')
    ,('Mercedes-Benz','C-Class')
    ,('Mitsubishi','Eclipse Cross')
    ,('Acura','ILX')
    ,('Acura','RDX')
    ,('Acura','TLX')
    ,('Alfa Romeo','Giulia')
    ,('Alfa Romeo','Stelvio')
    ,('Audi','A6 allroad')
    ,('Audi','A7')
    ,('Audi','Q3')
    ,('Audi','Q5')
    ,('Audi','Q8')
    ,('Audi','S5')
    ,('Audi','S6')
    ,('BMW','3 Series')
    ,('BMW','X1')
    ,('BMW','X3')
    ,('BMW','X5')
    ,('BMW','X7')
    ,('Buick','Enclave')
    ,('Buick','Encore')
    ,('Buick','Encore GX')
    ,('Buick','Envision')
    ,('Cadillac','CT4')
    ,('Cadillac','CT5')
    ,('Cadillac','Escalade')
    ,('Cadillac','Escalade ESV')
    ,('Cadillac','XT4')
    ,('Cadillac','XT5')
    ,('Cadillac','XT6')
    ,('Chevrolet','Blazer')
    ,('Chevrolet','Bolt EV')
    ,('Chevrolet','Camaro')
    ,('Chevrolet','Colorado Crew Cab')
    ,('Chevrolet','Colorado Extended Cab')
    ,('Chevrolet','Corvette')
    ,('Chevrolet','Equinox')
    ,('Chevrolet','Express 3500 Cargo')
    ,('Chevrolet','Malibu')
    ,('Chevrolet','Silverado 1500 Crew Cab')
    ,('Chevrolet','Silverado 2500 HD Crew Cab')
    ,('Chevrolet','Silverado 2500 HD Double Cab')
    ,('Chevrolet','Silverado 2500 HD Regular Cab')
    ,('Chevrolet','Silverado 3500 HD Crew Cab')
    ,('Chevrolet','Spark')
    ,('Chevrolet','Suburban')
    ,('Chevrolet','Tahoe')
    ,('Chevrolet','Trailblazer')
    ,('Chevrolet','Traverse')
    ,('Chevrolet','Trax')
    ,('Chrysler','300')
    ,('Chrysler','Pacifica')
    ,('Chrysler','Voyager')
    ,('Dodge','Challenger')
    ,('Dodge','Charger')
    ,('Dodge','Durango')
    ,('FIAT','500X')
    ,('Ford','Bronco Sport')
    ,('Ford','EcoSport')
    ,('Ford','Edge')
    ,('Ford','Escape')
    ,('Ford','Expedition')
    ,('Ford','Expedition MAX')
    ,('Ford','Explorer')
    ,('Ford','F150 SuperCrew Cab')
    ,('Ford','F250 Super Duty Crew Cab')
    ,('Ford','Mustang')
    ,('Ford','Ranger SuperCrew')
    ,('Ford','Transit Connect Cargo Van')
    ,('GMC','Acadia')
    ,('GMC','Canyon Crew Cab')
    ,('GMC','Canyon Extended Cab')
    ,('GMC','Sierra 1500 Crew Cab')
    ,('GMC','Sierra 2500 HD Crew Cab')
    ,('GMC','Sierra 3500 HD Crew Cab')
    ,('GMC','Terrain')
    ,('GMC','Yukon')
    ,('GMC','Yukon XL')
    ,('Genesis','G70')
    ,('Genesis','G80')
    ,('Genesis','G90')
    ,('Genesis','GV80')
    ,('Honda','Accord')
    ,('Honda','Accord Hybrid')
    ,('Honda','CR-V')
    ,('Honda','CR-V Hybrid')
    ,('Honda','Civic')
    ,('Honda','Civic Type R')
    ,('Honda','HR-V')
    ,('Honda','Insight')
    ,('Honda','Odyssey')
    ,('Honda','Passport')
    ,('Honda','Pilot')
    ,('Honda','Ridgeline')
    ,('Hyundai','Accent')
    ,('Hyundai','Elantra')
    ,('Hyundai','Ioniq Electric')
    ,('Hyundai','Ioniq Hybrid')
    ,('Hyundai','Ioniq Plug-in Hybrid')
    ,('Hyundai','Kona')
    ,('Hyundai','Kona Electric')
    ,('Hyundai','Palisade')
    ,('Hyundai','Santa Fe')
    ,('Hyundai','Sonata')
    ,('Hyundai','Sonata Hybrid')
    ,('Hyundai','Tucson')
    ,('Hyundai','Veloster')
    ,('Hyundai','Venue')
    ,('INFINITI','Q50')
    ,('INFINITI','QX50')
    ,('INFINITI','QX80')
    ,('Jaguar','E-PACE')
    ,('Jaguar','I-PACE')
    ,('Jeep','Cherokee')
    ,('Jeep','Compass')
    ,('Jeep','Gladiator')
    ,('Jeep','Grand Cherokee')
    ,('Jeep','Grand Cherokee L')
    ,('Jeep','Renegade')
    ,('Jeep','Wrangler')
    ,('Jeep','Wrangler Unlimited')
    ,('Kia','Forte')
    ,('Kia','K5')
    ,('Kia','Rio')
    ,('Kia','Sedona')
    ,('Kia','Seltos')
    ,('Kia','Sorento')
    ,('Kia','Soul')
    ,('Kia','Sportage')
    ,('Kia','Stinger')
    ,('Kia','Telluride')
    ,('Land Rover','Discovery')
    ,('Land Rover','Range Rover')
    ,('Land Rover','Range Rover Sport')
    ,('Land Rover','Range Rover Velar')
    ,('Lexus','ES')
    ,('Lexus','GX')
    ,('Lexus','IS')
    ,('Lexus','LC')
    ,('Lexus','LS')
    ,('Lexus','LX')
    ,('Lexus','NX')
    ,('Lexus','RC')
    ,('Lexus','RX')
    ,('Lexus','UX')
    ,('Lincoln','Aviator')
    ,('Lincoln','Corsair')
    ,('Lincoln','Nautilus')
    ,('Lincoln','Navigator')
    ,('Lincoln','Navigator L')
    ,('MAZDA','CX-3')
    ,('MAZDA','CX-30')
    ,('MAZDA','CX-5')
    ,('MAZDA','CX-9')
    ,('MAZDA','MAZDA3')
    ,('MAZDA','MAZDA6')
    ,('MAZDA','MX-5 Miata')
    ,('MAZDA','MX-5 Miata RF')
    ,('Mercedes-Benz','A-Class')
    ,('Mercedes-Benz','C-Class')
    ,('Mercedes-Benz','G-Class')
    ,('Mercedes-Benz','GLA')
    ,('Mercedes-Benz','GLB')
    ,('Mercedes-Benz','GLC')
    ,('Mercedes-Benz','GLE')
    ,('Mercedes-Benz','Mercedes-AMG G-Class')
    ,('Mitsubishi','Outlander Sport')
    ,('Nissan','Altima')
    ,('Nissan','Frontier Crew Cab')
    ,('Nissan','Frontier King Cab')
    ,('Nissan','Kicks')
    ,('Nissan','LEAF')
    ,('Nissan','Maxima')
    ,('Nissan','Murano')
    ,('Nissan','NV3500 HD Cargo')
    ,('Nissan','Rogue')
    ,('Nissan','Sentra')
    ,('Nissan','TITAN XD Crew Cab')
    ,('Nissan','Versa')
    ,('Polestar','2')
    ,('Porsche','718 Boxster')
    ,('Porsche','718 Cayman')
    ,('Porsche','718 Spyder')
    ,('Porsche','911')
    ,('Porsche','Cayenne')
    ,('Porsche','Cayenne Coupe')
    ,('Porsche','Macan')
    ,('Porsche','Panamera')
    ,('Porsche','Taycan')
    ,('Ram','1500 Crew Cab')
    ,('Ram','2500 Crew Cab')
    ,('Ram','2500 Mega Cab')
    ,('Ram','ProMaster Cargo Van')
    ,('Ram','ProMaster City')
    ,('Ram','ProMaster Window Van')
    ,('Rivian','R1S')
    ,('Rivian','R1T')
    ,('Subaru','Ascent')
    ,('Subaru','Crosstrek')
    ,('Subaru','Forester')
    ,('Subaru','Impreza')
    ,('Subaru','Legacy')
    ,('Subaru','Outback')
    ,('Tesla','Model S')
    ,('Tesla','Model X')
    ,('Toyota','4Runner')
    ,('Toyota','Avalon')
    ,('Toyota','Avalon Hybrid')
    ,('Toyota','C-HR')
    ,('Toyota','Camry')
    ,('Toyota','Camry Hybrid')
    ,('Toyota','Corolla')
    ,('Toyota','Corolla Hatchback')
    ,('Toyota','Corolla Hybrid')
    ,('Toyota','GR Supra')
    ,('Toyota','Highlander')
    ,('Toyota','Highlander Hybrid')
    ,('Toyota','Land Cruiser')
    ,('Toyota','Prius')
    ,('Toyota','Prius Prime')
    ,('Toyota','RAV4')
    ,('Toyota','RAV4 Hybrid')
    ,('Toyota','RAV4 Prime')
    ,('Toyota','Sequoia')
    ,('Toyota','Sienna')
    ,('Toyota','Tacoma Access Cab')
    ,('Toyota','Tacoma Double Cab')
    ,('Toyota','Tundra CrewMax')
    ,('Toyota','Tundra Double Cab')
    ,('Toyota','Venza')
    ,('Volkswagen','Atlas')
    ,('Volkswagen','Golf')
    ,('Volkswagen','Jetta')
    ,('Volkswagen','Passat')
    ,('Volkswagen','Tiguan')
    ,('Volvo','S90')
    ,('Volvo','V60')
    ,('Volvo','V90')
    ,('Volvo','XC40')
    ,('Volvo','XC60')
    ,('Volvo','XC90')
;
