FROM python:3.9.0-alpine
ADD . /code
WORKDIR /code
RUN pip install --no-cache-dir -r requirements.txt
CMD python main.py