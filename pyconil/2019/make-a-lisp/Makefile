all:
	$(error please pick a target)

test:
	rm -fr __pycache__
	flake8 scm.py test_scm.py
	python -m pytest -v
