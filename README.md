# Camps Service

This is a quick and small service written in Go to provide the [camps information](https://concentrationcamps.us/) released by the [2600 Magazine](https://www.2600.com/) in computer readable format (JSON).

## TO-DO

* ICE Camps

## Files

#### pull-info.rb

Reads the information from https://concentrationcamps.us/ and convert it into JSON.

#### camps.go

Very simple web service to provide a search functionality over such JSON file.
Right now only provides 2 end points:

* /camps/ >> returns all the camps
* /camps/?state=XX >> returns the camps in the given state, where XX is the abbreviation of the state. Ex: TX for Texas, etc.

#### Dockerfile

Docker file to create a docker image with this service. You can build the image by running the following command:

    docker image build -t camps .
    
And then you can run a docker container:

    docker container run -d -p 8080:8080 camps
    
Your service will be listening on port 8080. Pick and choose the configuration that meets your needs.
