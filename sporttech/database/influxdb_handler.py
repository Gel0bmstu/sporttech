import pandas as pd
from influxdb import InfluxDBClient, DataFrameClient

HOST = 'localhost'
PORT = 8086
USERNAME = "evv"
PASSWORD = "evvbmstu"

client = InfluxDBClient(host=HOST, port=PORT, username=USERNAME, password=PASSWORD)

testes = [
    {
        "measurement": "brushEvents",
        "tags": {
            "user": "Carol",
            "brushId": "6c89f539-71c6-490d-a28d-6c5d84c0ee2f"
        },
        "time": "2018-03-28T8:01:00Z",
        "fields": {
            "duration": 127
        }
    },
    {
        "measurement": "brushEvents",
        "tags": {
            "user": "Carol",
            "brushId": "6c89f539-71c6-490d-a28d-6c5d84c0ee2f"
        },
        "time": "2018-03-29T8:04:00Z",
        "fields": {
            "duration": 132
        }
    },
    {
        "measurement": "brushEvents",
        "tags": {
            "user": "Carol",
            "brushId": "6c89f539-71c6-490d-a28d-6c5d84c0ee2f"
        },
        "time": "2018-03-30T8:02:00Z",
        "fields": {
            "duration": 129
        }
    }
]


def database_exists(measurement_name):
    existed_databases = [meta['name'] for meta in client.get_list_database()]
    return measurement_name in existed_databases


def write_measurement_json(json_body):
    json_body_sample = json_body[0]
    measurement_name = json_body_sample['measurement']
    if database_exists(measurement_name):
        client.switch_database(measurement_name)
    else:
        client.create_database(measurement_name)
        client.switch_database(measurement_name)
    if scheme_valid(json_body_sample):
        client.write_points(json_body)
        return True
    else:
        return False


def scheme_valid(input_json):
    expected_keys = ['measurement', "tags", "time", "fields"]
    return set(input_json.keys()) == set(expected_keys)


def get_pandas_dataframe(measurement_name):
    pd_client = DataFrameClient(host=HOST, port=PORT, username=USERNAME, password=PASSWORD, database=measurement_name)
    data_dict = pd_client.query(f"SELECT * FROM {measurement_name}")[measurement_name]
    df = pd.DataFrame(data_dict)
    return df


if __name__ == "__main__":
    print(get_pandas_dataframe("brushEvents").columns)
