from app import create_app

app = create_app()

if __name__ == "__main__":
    app.run()
    #app.run(host='0.0.0.0', port='5001')

# from gevent.pywsgi import WSGIServer
# if __name__=='__main__':
#     WSGIServer(('127.0.0.1', 5001), app, numthreads=50).serve_forever()
#     WSGIServer(('-.0.0.0', 5001), app, numthreads=50).serve_forever()
