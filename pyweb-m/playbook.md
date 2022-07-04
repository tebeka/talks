setup:
    - terminator
    - slides

- slides
- show calc.py

    __name__
    __doc__
    Mention __file__

- run

    python calc.py -h
    python calc.py 2*pi
    python -m calc 2*pi

- Why: no need for script
- tests/test_calc.py
    
    pytest (fails)
    python -m pytest

- Show slide on sys.path
    python -m site

    python -m httpd (fails)
    mv httpd/__init__.py httpd/__main__.py
    python -m httpd

- Prefer "python -m pip install" (slides)

- Standard library support
    
    cat tebeka.json
    cat tebeka.json | python -m json.tool
    (mention jq)

    python -m zipfile -l code.zip
    cat quote.txt.gz| python -m gzip -d

    python -m http.server
    python -m telnetlib towel.blinkenlights.nl

    python -m dis calc.py
    python -m pdb calc.py "gcd(20, 8)"
    python -m timeit -s "from calc import calc" "calc('gcd(20, 8)')"
    python -m cProfile calc.py "gcd(20, 8)"

    python -m this
