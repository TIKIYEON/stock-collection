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
    value DECIMAL(10, 2),
    FOREIGN KEY (user_id) REFERENCES Users(uid)
);

CREATE TABLE PortfolioElement (
    buying_date DATE,
    buying_price DECIMAL(10, 2),
    price_change DECIMAL(10, 2),
    portfolio_id INT,
    stock_id INT,
    FOREIGN KEY (portfolio_id) REFERENCES Portfolios(pid),
    FOREIGN KEY (stock_id) REFERENCES Stocks(sid)
)


/*
Use to reference the relevant stock elements form purchase 
time to the current date
*/
CREATE TABLE Stocks (
    sid SERIAL PRIMARY KEY,
);

CREATE TABLE StockElements (
    date DATE,
    adj_close DECIMAL(10, 2),
    close DECIMAL(10, 2),
    open DECIMAL(10, 2),
    low DECIMAL(10, 2),
    high DECIMAL(10, 2),
    stock_id INT,
    FOREIGN KEY (stock_id) REFERENCES Stocks(sid) ON DELETE CASCADE
);
