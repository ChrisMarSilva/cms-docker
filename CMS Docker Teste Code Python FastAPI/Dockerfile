FROM python:3.9-slim

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade pip
RUN pip install --upgrade 'setuptools >=14.0'
RUN pip3 install fastapi uvicorn jinja2
RUN pip install "uvicorn[standard]" gunicorn
COPY ./requirements.txt /app
RUN pip install --no-cache-dir --upgrade -r requirements.txt

COPY ./src /app/src
COPY ./src/templates /app/src/templates
COPY ./src/static /app/src/s

EXPOSE 80

CMD ["uvicorn", "src.app:app", "--host", "0.0.0.0", "--port", "80", "--reload", "--debug", "--workers", "10"]
# CMD [ "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "15400" ] , "--log_level", "info"

