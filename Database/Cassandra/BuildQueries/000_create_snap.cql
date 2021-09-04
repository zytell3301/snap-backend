CREATE KEYSPACE IF NOT EXISTS snap WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
USE snap;