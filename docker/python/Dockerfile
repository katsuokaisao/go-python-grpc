FROM python:3.12

# 今回はpoetryなどは面倒なので使わない
RUN pip install --upgrade pip \
    && pip install pipenv

WORKDIR /home/src/python

COPY src/python/Pipfile .

RUN pipenv install

CMD ["pipenv", "run", "python", "main.py"]
