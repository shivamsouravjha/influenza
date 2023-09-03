CREATE TABLE companyEmail (    
    company_id INT PRIMARY KEY AUTO_INCREMENT,
    company_name VARCHAR(255),    
    company_mail TEXT,   
);


Sample company information

INSERT INTO companyEmail (company_name,  company_mail) VALUES
( 'trell',"trell.in"),
( 'zomato', 'zomato.com');
