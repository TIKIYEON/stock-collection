CREATE TABLE IF NOT EXISTS Users (
    uid SERIAL PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS Portfolio (
    pid SERIAL PRIMARY KEY,
    value DECIMAL(10, 2),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES Users(uid)
);

CREATE TABLE IF NOT EXISTS Stocks (
    sid SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS PortfolioElement (
    buying_date DATE,
    buying_price DECIMAL(10, 2),
    price_change DECIMAL(10, 2),
    portfolio_id INT,
    stock_id INT,
    FOREIGN KEY (portfolio_id) REFERENCES Portfolio(pid),
    FOREIGN KEY (stock_id) REFERENCES Stocks(sid)
);

CREATE TABLE IF NOT EXISTS StockElements (
    date DATE,
    open DECIMAL(10, 2),
    high DECIMAL(10, 2),
    low DECIMAL(10, 2),
    close DECIMAL(10, 2),
    adj_close DECIMAL(10, 2),
    volume INT,
    stock_id INT,
    FOREIGN KEY (stock_id) REFERENCES Stocks(sid) ON DELETE CASCADE
);

