SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';

--SELECT * FROM StockElements LIMIT 10;
--SELECT * FROM StockElements ORDER BY stock_id DESC LIMIT 10;
--SELECT * FROM stocks;
WITH RankedEntries AS (
    SELECT
        stock_id,
        date,
        ROW_NUMBER() OVER (PARTITION BY stock_id ORDER BY date DESC) AS rn
    FROM StockElements
)
DELETE FROM StockElements
WHERE (stock_id, date) IN (
    SELECT stock_id, date
    FROM RankedEntries
    WHERE rn > 100
);
SELECT COUNT(*) FROM StockElements;