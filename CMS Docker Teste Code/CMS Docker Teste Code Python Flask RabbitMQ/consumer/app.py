from flask import Flask, request  # pip install flask
import pika  # pip install pika
from threading import Lock, Thread
import time

app = Flask(__name__)

# host_redis = os.environ.get('HOST_REDIS', 'redis')
# port_redis = os.environ.get('PORT_REDIS', 6379)
# redis = Redis(host=host_redis, port=port_redis)

# @app.route('/')
# def hello():
#    redis.incr('hits')
#    return 'Hello World! %s times' % redis.get('hits')


# @app.route('/')
# @app.route("/", methods=['GET', 'POST'])
# def home():

# 	html = f'''
# 		<!doctype html>
# 		<html lang="en">
# 			<head>
# 				<meta charset="utf-8">
# 				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
# 				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
# 				<title>CMS RabbitQM/Redis</title>
# 			</head>
# 			<body>
# 				<div class="container-fluid text-center">
# 					<br>
# 					<br>
# 					<h1>CMS Teste RabbitQM</h1>
# 					<br>
# 					<div class="row">
# 						<div class="col-3"> </div>
# 						<div class="col-6">
# 							<form action="/enviar" method="post"> 
# 							  <button type="submit" class="btn btn-primary">Enviar Message RabbitMQ</button>
# 							  <a class="btn btn-warning" href="/receber" role="button">Receber Message RabbitMQ</a>
# 							  <br><br>
# 							  <input type="text" class="form-control" id="TxtMsgEnviar" name="TxtMsgEnviar">
# 							</form>
# 						</div>
# 						<div class="col-3"> </div>
# 					</div>
# 				</div>
# 				<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
# 				<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
# 			</body>
# 		</html>  
# 		'''
# 	return html


# @app.route("/enviar", methods=['GET', 'POST'])
# def enviar():

# 	data = None
# 	if request.method == 'POST':
# 		data = request.form
# 	elif request.method == 'GET':
# 		data = request.args

# 	msg = None
# 	msg_ok =  ''
# 	msg_erro =  ''

# 	try:

# 		msg = data.get('TxtMsgEnviar')

# 		if msg:
# 			connection = pika.BlockingConnection(pika.ConnectionParameters(host="localhost"))
# 			channel = connection.channel()
# 			channel.queue_declare(queue="hello", durable=True,)
# 			channel.basic_publish(exchange='', routing_key="hello", body=msg)
# 			# channel.close()
# 			connection.close()
# 			msg_ok = 'Mensagem Enviada'
# 		else:
# 			msg_erro = f'Sem Mensagem para Enviar'

# 	except Exception as e:
# 		msg_erro = f'Erro ao Enviar Mensagem: {str(e)}'

# 	html = f'''
# 		<!doctype html>
# 		<html lang="en">
# 			<head>
# 				<meta charset="utf-8">
# 				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
# 				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
# 				<title>CMS RabbitQM/Redis</title>
# 			</head>
# 			<body>
# 				<div class="container-fluid text-center">
# 					<br>
# 					<br>
# 					<h1>CMS Teste RabbitQM</h1>
# 					<br>
# 					<div class="row">
# 						<div class="col-3"> </div>
# 						<div class="col-6">
# 							<form action="/enviar" method="post"> 
# 							  <button type="submit" class="btn btn-primary">Enviar Message RabbitMQ</button>
# 							  <a class="btn btn-warning" href="/receber" role="button">Receber Message RabbitMQ</a>
# 							  <br><br>
# 							  <input type="text" class="form-control" id="TxtMsgEnviar" name="TxtMsgEnviar">
# 							  <br>
# 							  <small class="form-text text-success font-weight-bold">{msg_ok}</small>
# 							  <small class="form-text text-danger font-weight-bold">{msg_erro}</small>
# 							  <br>
# 							 <!--   <a class="btn btn-secondary" href="/" role="button">Voltar</a> -->
# 							</form>
# 						</div>
# 						<div class="col-3"> </div>
# 					</div>
# 				</div>
# 				<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
# 				<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
# 			</body>
# 		</html>  
# 		'''
# 	return html


lista = []

@app.route('/')
@app.route("/receber", methods=['GET', 'POST'])
def receber():

	msg_ok =  ''
	msg_erro = 'Nenhuma Mensagem Recebida'
	lista = []

	try:

		try:
			connection = pika.BlockingConnection(pika.ConnectionParameters("localhost"))
		except Exception as e:
			connection = pika.BlockingConnection(pika.ConnectionParameters("cms_python_flask_rabbitmq"))
			
		channel = connection.channel()
		channel.queue_declare(queue="hello", durable=True)

		def _callback(ch, method, properties, body):
			lista.append(body.decode())
			#channel.basic_ack(delivery_tag = method.delivery_tag)
			#channel.cancel()
			#channel.stop_consuming()

		def _closeChannel(channel_):
			time.sleep(.001)
			#if channel_.is_open:
			channel_.stop_consuming()
			# channel_.close()

		# channel.basic_qos(prefetch_count=1)
		channel.basic_consume(queue="hello", on_message_callback=_callback, auto_ack=True)

		t = Thread(target=_closeChannel, args=[channel])
		t.start()
		# t.join()

		try:
			channel.start_consuming()
		except KeyboardInterrupt:
			pass # channel.stop_consuming()

		if lista:
			i = 0
			for row in lista:
				i += 1
				msg_ok += f'#{i} - "{row}"" <br>'
				msg_erro = ''

		# for method_frame, properties, body in channel.consume('hello'):
		# 	msg_ok =+ f'{body.decode()} <br>'
		# 	channel.basic_ack(method_frame.delivery_tag)
		# 	#channel.cancel()
		# 	#channel.stop_consuming()
		# 	#if method_frame.delivery_tag == 10:
		# 	break

		# channel.cancel()
		# channel.close()
		connection.close()

		# with open(r"C:\Users\chris\Desktop\Docker\CMS Teste Python Flask Docker\myfile.txt", "w") as file1: 
		# 	file1.write("Hello \n") 

		# with open(r"myfile.txt", "r") as file1: 
		# 	msg_ok += file1.read()

	except Exception as e:
		msg_erro = f'Erro ao Receber Mensagem: {str(e)}'

	html = f'''
		<!doctype html>
		<html lang="en">
			<head>
				<meta charset="utf-8">
				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
				<title>CMS RabbitQM</title>
			</head>
			<body>
				<div class="container-fluid text-center">
					<br>
					<br>
					<h1>CMS - RabbitQM - Consumer</h1>
					<br>

					<div class="row">
						<div class="col-3"> </div>
						<div class="col-6">
							<!-- <a class="btn btn-primary" href="/" role="button">Enviar Message RabbitMQ</a> -->
							<a class="btn btn-warning" href="/receber" role="button">Receber Message RabbitMQ</a>
							<br>
							<br>
						    <small class="form-text text-danger font-weight-bold">{msg_erro}</small>
						    <small class="form-text text-success font-weight-bold">{msg_ok}</small>
						</div>
						<div class="col-3"> </div>
					</div>

				</div>
				<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
				<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
			</body>
		</html>  
		'''
	return html


if __name__ == "__main__":
    app.run(host='0.0.0.0' , port=5001, debug=True) 

