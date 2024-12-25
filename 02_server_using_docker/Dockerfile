# the basic use this "Dockerfile" is to pull the image we need from docker hub, we do this
# so that we don't have to do those manually.

FROM golang:1.23.4-alpine3.21

# running ubuntu command
# the command will make a directory inside the image with name "app"
RUN mkdir /app

# adding everything whats inside in "main.go" file
# the below command will take everything to the app directory in the image
ADD . /app

# most imp directory, all the commands will be running inside this directory.
WORKDIR /app

# build the executable file
RUN go build -o main .

# CMD helps us run command, here it will run the executable file we got from the 5th step
CMD [ "/app/main" ]