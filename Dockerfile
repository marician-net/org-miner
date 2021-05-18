FROM avcosystems/golang-node


RUN apt -y update
RUN apt -y install ocl-icd-opencl-dev

COPY . .

RUN chmod +x serve.sh

CMD ["./serve.sh"]