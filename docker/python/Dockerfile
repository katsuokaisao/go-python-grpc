FROM python:3.12

# 今回はpoetryなどは面倒なので使わない
RUN pip install --upgrade pip \
    && pip install grpcio-tools \
    grpcio \
    pipenv

WORKDIR /home/src/python

COPY src/python/Pipfile .

RUN pipenv install

WORKDIR /home/

CMD ["bash", "-c", "cd src/python && pipenv run python main.py"]
