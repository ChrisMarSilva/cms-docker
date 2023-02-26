# FROM pypy:3-slim-buster
FROM python:3.9.0-slim-buster

WORKDIR /code

ENV TZ="America/Sao_Paulo" PYTHONUNBUFFERED=1 PYTHONDONTWRITEBYTECODE=1 PYTHONHASHSEED=random PIP_NO_CACHE_DIR=off PIP_DISABLE_PIP_VERSION_CHECK=on

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# install system dependencies
RUN apt-get update && apt-get install -y build-essential python-dev python-pip
RUN pip install --upgrade pip
RUN pip install --upgrade 'setuptools >=14.0'
RUN pip install --no-cache-dir pandas
RUN pip install uWSGI==2.0.17

COPY requirements.txt /code/
RUN pip install --no-cache-dir -r requirements.txt

COPY . /code/

EXPOSE 5000

#CMD python wsgi.py
# ENTRYPOINT ["python", "wsgi.py"]
# RUN chmod +x ./start.sh
# ENTRYPOINT ["sh", "start.sh"]
# ENTRYPOINT ["gunicorn", "--config", "gunicorn.py", "wsgi:app"]
# ENTRYPOINT ["gunicorn", "--config", "wsgi.py", "app:create_app()"]
#CMD [ "gunicorn", "-b", ":5000", "run" ]
CMD ["gunicorn", "wsgi:app", "-b", "0.0.0.0:5000"]
# CMD ["uwsgi", "--ini", "uwsgi.ini"]

# HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD ["executable"]
# HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl -f http://localhost:5000/health || exit 1 
# HEALTHCHECK --interval=30s --timeout=30s --start-period=5s CMD curl -f --retry 10 --max-time 15 --retry-delay 10 --retry-max-time 60 "http://localhost:5000/health" || exit 1   
