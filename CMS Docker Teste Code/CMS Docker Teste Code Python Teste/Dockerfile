FROM python:3.9.0-alpine

ADD . /code
WORKDIR /code

RUN pip install --upgrade pip
RUN pip install --upgrade 'setuptools >=14.0'
RUN pip install --no-cache-dir -r requirements.txt

CMD python main.py
#CMD [ "python", "./code/main.py" ]
