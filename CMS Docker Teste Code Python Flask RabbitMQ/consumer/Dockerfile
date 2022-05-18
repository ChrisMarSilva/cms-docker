FROM python:3.9.0-alpine
ADD . /consumer
WORKDIR /consumer
RUN pip install --no-cache-dir -r requirements.txt
EXPOSE 5001
CMD python app.py