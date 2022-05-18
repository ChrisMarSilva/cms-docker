# -*- coding: utf-8 -*-
import os
import time
import datetime as dt


if __name__ == "__main__":
    while True:
        print(f"{dt.datetime.now()} - I'm working...")
        print("TZ:", os.getenv('TZ', 'N/D'))
        print("TNB_BD_HOST:", os.getenv('TNB_BD_HOST', 'N/D'))
        time.sleep(1)
