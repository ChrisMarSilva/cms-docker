FROM python:3.9.0-alpine
ADD . /app
WORKDIR /app
RUN pip install --no-cache-dir -r requirements.txt
ADD . /app
EXPOSE 5000
ENTRYPOINT ["python", "app/app.py"]