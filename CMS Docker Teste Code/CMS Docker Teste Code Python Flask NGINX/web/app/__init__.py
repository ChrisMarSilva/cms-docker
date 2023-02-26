
def create_app():

    from flask import Flask, render_template, make_response, request
    from flask_cors import CORS
    import json
    import random

    app = Flask(__name__)
    app.config['DEBUG'] = True

    CORS(app)
    # cors = CORS(app, resources={r"/api/*": {"origins": "*"}})

    @app.route('/')
    def home():
        from os import getpid
        return f'TESTE #3 - PID: {str(getpid())}'

    @app.route('/health', methods=['GET'])
    def get(self):
        return json.dumps({'success': True, 'message': "healthy"})

    @app.route('/wait')
    def wait():
        import time
        time.sleep(3)
        return 'WAIT #3'

    @app.route('/template')
    def template():
        return render_template('home.html')

    @app.route("/login", methods=['GET', 'POST'])
    def login():
        try:

            data = None
            if request.method == 'POST':
                data = request.form
            elif request.method == 'GET':
                data = request.args
            if not data:
                data = request.get_json(silent=True)

            try:
                email = data.get('email')
                senha = data.get('senha')
            except:
                response = {"message": "Login/Senha não informado"}
                return make_response(json.dumps(response), 400)

            response = {
                "codigo": '123',
                "nome": "Garçom 123",
            }

            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Login"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/logout", methods=['GET', 'POST'])
    def logout():
        try:

            response = {"message": ""}
            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Login"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/cartoes", methods=['GET', 'POST'])
    def lista_cartao():
        try:

            response = [
                {
                    "codigo": str(numero).zfill(4),
                    "nome": f"Pessoa {str(numero).zfill(4)}",
                    "situacao": "Ativo" if numero % 5 else "Bloqueado",
                }
                for numero in range(1, 101)  # 101 # 11
            ]

            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Lista Cartao"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/garotas", methods=['GET', 'POST'])
    def lista_garota():
        try:

            response = [
                {
                    "codigo": str(numero).zfill(4),
                    "nome": f"Garota {str(numero).zfill(4)}",
                }
                for numero in range(1, 11)  # 101 # 21 # 11
            ]

            response.append({"codigo": "9996", "nome": "Camila"})
            response.append({"codigo": "9997", "nome": "Andressa"})
            response.append({"codigo": "9998", "nome": "Juju"})
            response.append({"codigo": "9999", "nome": "Andressa"})

            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Lista Garota"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/produtos", methods=['GET', 'POST'])
    def lista_produto():
        try:

            response = [
                {
                    "codigo": str(numero).zfill(4),
                    "descricao": f"Produto {str(numero).zfill(4)}",
                    "valor": 10.00,  # round(random.uniform(1, 2000), 2),
                }
                for numero in range(1, 11)
            ]

            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Lista Produto"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/fazerpedido", methods=['GET', 'POST'])
    def fazer_pedido():
        try:

            data = None
            if request.method == 'POST':
                data = request.form
            elif request.method == 'GET':
                data = request.args
            if not data:
                data = request.get_json(silent=True)

            # response = {"message": data.get('garcom')}
            # response = {"message": data.get('listaCartao')}
            # response = {"message": json.dumps(data.get('listaCartao'))}
            # response = {"message": data.get('garota')}
            # response = {"message": data.get('listaProduto')}
            # response = {"message": json.dumps(data.get('listaProduto'))}

            try:
                cod_garcom = data.get('garcom')
                lst_cartoes = data.get('listaCartao')
                cod_garota = data.get('garota')
                lst_produtos = data.get('listaProduto')
                # cod_autorizador = data.get('cod_autorizador')  # codigo do autorizador
                # senha_autorizador = data.get('senha_autorizador')  # senha do autorizador
            except:
                response = {"message": "Parâmetros não informado"}
                return make_response(json.dumps(response), 400)

            response = {}
            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Fazer Pedido"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    @app.route("/consumocartao", methods=['GET', 'POST'])
    def consumo_cartao():
        try:

            data = None
            if request.method == 'POST':
                data = request.form
            elif request.method == 'GET':
                data = request.args
            if not data:
                data = request.get_json(silent=True)

            try:
                cod_cartao = data.get('codigo')  # codigo do cartao
            except:
                response = {"message": "Cartão não informado"}
                return make_response(json.dumps(response), 400)

            response = [
                {
                    "descricao": f"{str(numero).zfill(4)} Produto Produto {str(numero).zfill(4)}",
                    "quantidade": random.randint(1, 10),
                    "valor": round(random.uniform(1, 2000), 2)
                }
                for numero in range(1, 21)
            ]

            return make_response(json.dumps(response), 200)  # 200 (OK)

        except:
            response = {"message": "Falha Geral Consumo Cartao"}
            # 400 (Bad Request)
            return make_response(json.dumps(response), 400)

    return app


# if __name__ == '__main__':
#     #app = create_app()
#     #app.run(host='0.0.0.0', port=5000)
#     pass
