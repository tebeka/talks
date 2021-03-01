FROM python:3.9-slim
RUN apt-get update && apt-get install -y curl gcc make pkg-config

RUN curl -LO https://golang.org/dl/go1.15.7.linux-amd64.tar.gz
RUN tar xzf go*.tar.gz -C /opt
RUN rm go*.tar.gz
ENV PATH=/opt/go/bin:$PATH

RUN python -m pip install --upgrade pip
RUN python -m pip install numpy

ARG USER=fosdem
RUN useradd -ms /bin/bash $USER

WORKDIR /home/$USER/outliers

COPY go.mod .
RUN chown -R $USER .

USER $USER
RUN go mod download

RUN echo 'export PS1="\$ "' >> /home/$USER/.bashrc

CMD "/bin/bash"
