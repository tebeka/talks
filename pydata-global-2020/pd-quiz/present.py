"""Present"""
from subprocess import run
from sys import executable
from tempfile import NamedTemporaryFile
import re

from markdown import markdown
from fastapi import FastAPI, Request, Response
from starlette.background import run_in_threadpool

app = FastAPI()
slides = []


def split_to_slides(text):
    lines = []

    for line in text.splitlines():
        if re.match(r'^#{1,2}\s', line):
            if lines:
                yield '\n'.join(lines)
            lines.clear()
        lines.append(line)

    if lines:
        yield '\n'.join(lines)


@app.get('/')
async def index(page: int = 0):
    page = min(page, len(slides) - 1)
    html = markdown(slides[page])
    if page < len(slides) - 1:
        html += f'<hr/><p><a href="/?page={page+1}">Next</a></p>'
    return Response(content=html, media_type='text/html')


def run_code(code):
    tmp = NamedTemporaryFile('wb', suffix='.py')
    tmp.write(code)
    tmp.flush()

    out = run([executable, tmp.name], capture_output=True)
    return {
        'ok': out.returncode == 0,
        'stdout': out.stdout.decode('utf-8'),
        'stderr': out.stderr.decode('utf-8'),
    }


@app.post('/execute')
async def execute(req: Request):
    code = await req.body()
    out = await run_in_threadpool(run_code, code)
    return out


@app.on_event('startup')
def on_startup():
    with open('quiz.slide') as fp:
        data = fp.read()
    slides.extend(split_to_slides(data))


if __name__ == '__main__':
    from argparse import ArgumentParser  # , FileType
    import uvicorn

    parser = ArgumentParser(description=__doc__)
    # parser.add_argument('file', help='slides file', type=FileType('r'))
    parser.add_argument(
        '--port', help='port to listen on', type=int, default=9999,
    )
    args = parser.parse_args()

    # data = args.file.read()
    # args.file.close()
    # slides.extend(split_to_slides(data))

    uvicorn.run(
        app='present:app',  # must pass app as str for reload
        port=args.port,
        host='0.0.0.0',
        reload=True,
    )
