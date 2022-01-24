import psycopg2
from const import *

class SQLighter:

    def __init__(self, database):
        self.connection = psycopg2.connect(
            host=host,
            port=port,
            user=user,
            password=password,
            dbname=database
        )

        self.cursor = self.connection.cursor()

    def add_json(self, json_data):
        with self.connection:
            return self.cursor.execute(f"INSERT INTO json_info (json_data) VALUES(%s)", (json_data,))

    def add_cloud(self, cloud, json_txt):
        with self.connection:
            return self.cursor.execute(f"INSERT INTO {cloud} (json_data) VALUES(%s)", (json_txt,))

    def send_json (self, cloud):
        with self.connection:
            self.cursor.execute(f"SELECT json_data from {cloud}")
            rows = self.cursor.fetchall()
            js = ''
            for row in rows:
                for i in row:
                    js += str(i)
            return js

    def delete_table (self, cloud):
        with self.connection:
            return self.cursor.execute(f"DELETE FROM {cloud};")



    def close(self):
        """Закрываем соединение с БД"""
        #print ('ЗАКРЫЛОСЬ')
        self.connection.close()
