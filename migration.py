import psycopg2
import csv
import os

# Define connection parameters
conn_params = {
    'host': 'localhost',
    'port': 5432,
    'database': 'packaform_test',
    'user': 'postgres',
    'password': 'postgres'
}

# Connect to PostgreSQL
conn = psycopg2.connect(**conn_params)
cur = conn.cursor()

# Get CSV filename and use it as tablename
csv_dir = 'test_data/'

# Create table if not exists


def createTable(fileName, tableName):
    fileName = csv_dir + fileName
    with open(fileName, 'r') as f:
        reader = csv.reader(f)
        header = next(reader)

        create_query = f"CREATE TABLE IF NOT EXISTS {tableName} ("
        for col_name in header:
            create_query += f"{col_name} TEXT,"
        create_query = create_query[:-1] + ")"
        cur.execute(create_query)
        conn.commit()


# Read data from CSV and insert into PostgreSQL
def readAndInsertData(fileName, tableName):
    fileName = csv_dir + fileName
    with open(fileName, 'r') as f:
        reader = csv.reader(f)

        header = next(reader)
        # values = next(reader)

        for values in reader:
            columns = ", ".join(header)
            placeholders = ", ".join(['%s'] * len(values))
            query = f"INSERT INTO {tableName} ({columns}) VALUES ({placeholders})"
            cur.execute(query, values)
        conn.commit()


if __name__ == "__main__":
    csv_file = os.listdir(csv_dir)
    for fname in csv_file:
        base_name, extension = os.path.splitext(fname)
        tableName = base_name.split()[-1]
        createTable(fname, tableName)
        readAndInsertData(fname, tableName)

        # cur.close()
        # conn.close()
