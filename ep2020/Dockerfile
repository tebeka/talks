FROM python:3.8-slim
RUN apt-get update
RUN apt-get install -y cowsay vim xz-utils
RUN mkdir -p ~/.ipython/profile_default    
COPY ipython_config.py ~/.ipython/profile_default
RUN python -m pip install --upgrade pip
RUN python -m pip install \
    ipython~=7.15.0 \
    ipython-sql~=0.4.0 \
    matplotlib~=3.2 \
    pandas~=1.0.4
RUN bash -c "echo set -o vi >> ~/.bashrc"
RUN bash -c "echo export PS1=\\'$ \\' >> ~/.bashrc"
RUN bash -c "echo clear >> ~/.bashrc"
WORKDIR /ep2020
