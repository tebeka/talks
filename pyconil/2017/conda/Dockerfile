FROM ubuntu:16.04

RUN sed -i 's/# deb-src/deb-src/' /etc/apt/sources.list
RUN apt update
RUN apt build-dep -y python3
RUN apt install -y curl zlib1g-dev libssl-dev
RUN curl -LO https://www.python.org/ftp/python/3.6.1/Python-3.6.1.tar.xz
RUN tar xJf Python-3.6.1.tar.xz
RUN cd Python-3.6.1  && ./configure --prefix=/opt && make install
RUN curl -LO https://repo.continuum.io/miniconda/Miniconda3-latest-Linux-x86_64.sh
RUN bash Miniconda3-latest-Linux-x86_64.sh -b -p /miniconda
RUN /miniconda/bin/conda update -y --all
RUN rm -rf Python-3.6.1
RUN rm Miniconda3-latest-Linux-x86_64.sh
RUN echo 'export PATH=/opt/bin:/miniconda/bin:$PATH' >> ~/.bashrc
