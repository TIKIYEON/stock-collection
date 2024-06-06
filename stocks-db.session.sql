SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';

SELECT * FROM StockElements LIMIT 10;
SELECT * FROM StockElements ORDER BY stock_id DESC LIMIT 10;
SELECT * FROM stocks;