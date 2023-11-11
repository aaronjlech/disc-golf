DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   pg_tables
        WHERE  schemaname = 'public'
        AND    tablename = 'manufacturer'
    ) THEN
        CREATE TABLE manufacturer (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL
        );
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM   pg_tables
        WHERE  schemaname = 'public'
        AND    tablename = 'discs'
    ) THEN
        CREATE TABLE discs (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            speed DECIMAL,
            turn DECIMAL,
            glide DECIMAL,
            fade DECIMAL,
            manufacturer_id INTEGER REFERENCES manufacturer(id)
            -- Add more columns as needed
        );
    END IF;
END $$;