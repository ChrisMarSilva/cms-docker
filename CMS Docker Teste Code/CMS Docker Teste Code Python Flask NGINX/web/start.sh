#!/bin/sh
exec gunicorn -w 10 --t 2 -b 0.0.0.0:5000 wsgi:app
# exec gunicorn --config gunicorn.py wsgi:app
# exec gunicorn --bind 0.0.0.0:5000 app.wsgi:app
# exec gunicorn --bind 0.0.0.0:5000 --workers 4 "app.create_app:create_app()"
# exec gunicorn --workers=10 --threads=5 --timeout=120 --bind=0.0.0.0:5000 wsgi:app
# exec gunicorn --chdir app wsgi:app -w 2 --threads 2 -b 0.0.0.0:5000