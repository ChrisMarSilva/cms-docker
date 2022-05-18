# -*- coding: utf-8 -*-
import time
import datetime as dt
import schedule  # pip install schedule
import mysql.connector # pip install mysql-connector-python
#  pip install --upgrade pip


# def work():
#     print(f"{dt.datetime.now()} - I'm working")
#     #print(f"{dt.datetime.now()} - I'm working - INI...")
#     #time.sleep(10)
#     #print(f"{dt.datetime.now()} - I'm working - FIM...")
#schedule.every(2).seconds.do(work)
# schedule.every(1).minutes.do(work)
# schedule.every(10).minutes.do(work)
# schedule.every().hour.do(work)
# schedule.every().day.at("10:30").do(work)
# schedule.every().monday.do(work)
# schedule.every().wednesday.at("13:15").do(work)
# schedule.every().minute.at(":17").do(work)


print(f"{dt.datetime.now()} - TESTE CMS - INICIO")
print(f"")
# time.sleep(10) # 10seg  # time.sleep(.300) # Wait for 300 milliseconds

#while True:
    #schedule.run_pending()
    #time.sleep(1)

try:
    print(f"#1 Try conectiong localhost:3306")
    conn = mysql.connector.connect(user='root', password='', host='localhost', port=3306, database='tamonabo_BDCMSTamoNaBolsa')
except Exception as e:
    print("#1", str(e))

    try:
        print(f"#2 Try conectiong 127.0.0.1:3306")
        conn = mysql.connector.connect(user='root', password='', host='127.0.0.1', port=3306, database='tamonabo_BDCMSTamoNaBolsa')
    except Exception as e:
        print("#2", str(e))
    
        try:
            print(f"#3 Try conectiong 189.51.158.212:3306")
            conn = mysql.connector.connect(user='root', password='', host='189.51.158.212', port=3306, database='tamonabo_BDCMSTamoNaBolsa')
        except Exception as e:
            print("#3", str(e))
    
            try:
                print(f"#4 Try conectiong ns742.hostgator.com.br:3306")
                conn = mysql.connector.connect(user='tamonabo_rootcms', password='Chrs8723', host='ns742.hostgator.com.br', port=3306, database='tamonabo_BDCMSTamoNaBolsa')
            except Exception as e:
                print("#4", str(e))

cursor  = conn.cursor() 
cursor.execute("SELECT A.ID, A.NOME FROM TBUSUARIO A")

print(f"")
print(f"USUARIOS:")
print(f"")

rows = cursor.fetchall()
for row in rows:
    print(f'{dt.datetime.now()} - #{str(row[0])} - {str(row[1])}')
    time.sleep(.010)

cursor.close()
conn.close()

print(f"")
print(f"{dt.datetime.now()} - TESTE CMS - FIM")


while True:
    print(f"{dt.datetime.now()} - I'm working...")
    time.sleep(1)
    




'''



import Queue
import time
import threading
import schedule


def job():
    print("I'm working")


def worker_main():
    while 1:
        job_func = jobqueue.get()
        job_func()
        jobqueue.task_done()

jobqueue = Queue.Queue()

schedule.every(10).seconds.do(jobqueue.put, job)
schedule.every(10).seconds.do(jobqueue.put, job)
schedule.every(10).seconds.do(jobqueue.put, job)
schedule.every(10).seconds.do(jobqueue.put, job)
schedule.every(10).seconds.do(jobqueue.put, job)

worker_thread = threading.Thread(target=worker_main)
worker_thread.start()

while 1:
    schedule.run_pending()
    time.sleep(1)


    mport threading
import time
import schedule

def job():
    print("I'm running on thread %s" % threading.current_thread())

def run_threaded(job_func):
    job_thread = threading.Thread(target=job_func)
    job_thread.start()

schedule.every(10).seconds.do(run_threaded, job)
schedule.every(10).seconds.do(run_threaded, job)
schedule.every(10).seconds.do(run_threaded, job)
schedule.every(10).seconds.do(run_threaded, job)
schedule.every(10).seconds.do(run_threaded, job)

while 1:
    schedule.run_pending() # How to execute jobs in parallel
    time.sleep(1)


     https://schedule.readthedocs.io/en/stable/

import functools
import time
import schedule # 

def with_logging(func): # This decorator can be applied to
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print('LOG: Running job "%s"' % func.__name__)
        result = func(*args, **kwargs)
        print('LOG: Job "%s" completed' % func.__name__)
        return result
    return wrapper

@with_logging
def job(param: str='xxxxx'):
    print(f"{str(param)}: I'm working...")


def catch_exceptions(cancel_on_failure=False):
    def catch_exceptions_decorator(job_func):
        @functools.wraps(job_func)
        def wrapper(*args, **kwargs):
            try:
                return job_func(*args, **kwargs)
            except:
                import traceback
                print(traceback.format_exc())
                if cancel_on_failure:
                    return schedule.CancelJob
        return wrapper
    return catch_exceptions_decorator

@catch_exceptions(cancel_on_failure=True)
def bad_task():
    return 1 / 0

def job_that_executes_once():
    return schedule.CancelJob


schedule.clear() 

schedule.every(10).seconds.do(job, param='seconds')
schedule.every(5).to(10).seconds.do(job, param='seconds-5-10') # This job will execute every 5 to 10 seconds.

schedule.every(5).minutes.do(bad_task)
schedule.every(10).minutes.do(job, param='minutes')# After every 10mins geeks() is called.  
schedule.every(5).to(10).minutes.do(job, param='minute-5-10')  # After every 5 to 10mins in between run work() 
schedule.every().minute.at(":17").do(job, param='minute')

schedule.every().hour.do(job, param='hour')# After every hour geeks() is called. 

schedule.every().day.at("10:30").do(job, param='day')# Every day at 12am or 00:00 time bedtime() is called. 
schedule.every().day.at('22:30').do(job_that_executes_once)

schedule.every().monday.do(job, param='monday')# Every monday good_luck() is called 
schedule.every().wednesday.at("13:15").do(job, param='wednesday')
schedule.every().tuesday.at("18:00").do(job, param='tuesday')  # Every tuesday at 18:00 sudo_placement() is called 

# schedule.clear('daily-tasks') 
# schedule.every().day.do(greet, 'Andrea').tag('daily-tasks', 'friend')
# schedule.every().hour.do(greet, 'John').tag('hourly-tasks', 'friend')
# schedule.every().hour.do(greet, 'Monica').tag('hourly-tasks', 'customer')
# schedule.every().day.do(greet, 'Derek').tag('daily-tasks', 'guest')
# schedule.clear('daily-tasks')

# 0-Monday-segunda-feira 
# 1-Tuesday-terça-feira 
# 2-Wednesday-quarta-feira 
# 3-Thursday-quinta-feira 
# 4-Friday-sexta-feira 
# 5-Saturday-Sábado 
# 6-Sunday-domingo

contador = 0
while True:
    contador += 1
    #print(f'{str(contador)} ', end='')
    schedule.run_pending()
    if contador >= 100:
        schedule.clear()
        break
    time.sleep(1)

'''