# base image
FROM python:3.8.2-slim as aliash_tool_base_image
LABEL stage=aliash_tool_base_image
RUN apt-get update \
  && apt-get install gcc -y \
  && apt-get clean

WORKDIR aliash_tool

COPY ./README.md ./README.md
COPY ./main.py ./main.py
COPY ./requirements.txt ./requirements.txt
COPY ./setup.py ./setup.py
COPY ./aliash_tool ./aliash_tool

RUN pip install --upgrade pip
RUN pip install --user --editable .

# production image
FROM python:3.8.0-slim as aliash_tool_pod
COPY --from=aliash_tool_base_image /root/.local /root/.local
COPY --from=aliash_tool_base_image /aliash_tool /aliash_tool
WORKDIR aliash_tool
ENV PATH=/root/.local/bin:$PATH
ENTRYPOINT ["aliash"]
