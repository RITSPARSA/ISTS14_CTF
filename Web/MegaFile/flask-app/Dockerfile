# Set the base image to Ubuntu
FROM ubuntu

# File Author / Maintainer
MAINTAINER Maintaner app

# Add the application resources URL
RUN echo "deb http://archive.ubuntu.com/ubuntu/ $(lsb_release -sc) main universe" >> /etc/apt/sources.list

# Update the sources list
RUN apt-get update

# Install basic applications
RUN apt-get install -y tar git curl nano wget dialog net-tools build-essential

# Install Python and Basic Python Tools
RUN apt-get install -y python python-dev python-distribute python-pip

# Install Flask
RUN pip install Flask

# Copy the application folder inside the container
RUN mkdir /app
ADD src /app

# Get pip to download and install requirements:
#RUN pip install -r /my_application/requirements.txt

# Expose ports
EXPOSE 8080

# Set the default directory where CMD will execute
WORKDIR /app

# Set the default command to execute    
# when creating a new container
# i.e. using CherryPy to serve the application
CMD python app.py
