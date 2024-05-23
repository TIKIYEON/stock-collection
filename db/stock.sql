CREATE DATABASE stocks;

CREATE TABLE Users (
    uid SERIAL PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20)
);

/* 
 
 */
CREATE TABLE Portfolios (
    pid SERIAL PRIMARY KEY,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES Users(uid)
);


/*
Use to reference the relevant stock elements form purchase 
time to the current date
*/
CREATE TABLE Stocks (
    sid SERIAL PRIMARY KEY,
);

CREATE TABLE StockElements (
    sid SERIAL PRIMARY KEY,
    date DATE,
    adj_close DECIMAL(10, 2),
    close DECIMAL(10, 2),
    open DECIMAL(10, 2),
    low DECIMAL(10, 2),
    high DECIMAL(10, 2),
    stock_id INT,
    FOREIGN KEY (stock_id) REFERENCES Stocks(sid)
);
