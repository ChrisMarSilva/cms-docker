from fastapi import FastAPI, Request, Form
from fastapi.responses import HTMLResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from os import getpid

app = FastAPI()
app.mount("/static", StaticFiles(directory="src/static"), name="static")
templates = Jinja2Templates(directory="src/templates")


@app.get("/")
def index():
    return {"message": f"hello world 2 - PID: {str(getpid())}"}

@app.get("/home", response_class=HTMLResponse)
async def home(request: Request):
    try:
        return templates.TemplateResponse("home.html", {"request": request, "nome": "Chris MarSil"})
    except Exception as e:
        return {"erro": str(e)}


@app.get("/form")
async def form(request: Request):
    try:
        return templates.TemplateResponse('home.html', context={'request': request, 'nome': "Type a number"})
    except Exception as e:
        return {"erro": str(e)}


# if __name__ == "__main__":
#     uvicorn.run("app:app", port=5001, reload=True)

