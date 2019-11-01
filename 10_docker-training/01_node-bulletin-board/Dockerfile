# 
FROM node:13.0.1

# Create app directory in image filesystem 
WORKDIR /usr/src/app

# COPY the file package.json from your host to the present location (.)
# in your image (so in this case, to /usr/src/app/package.json)
COPY ./bulletin-board-app/package.json .

# RUN the command npm install inside your image filesystem 
# which will read package.json to determine your app’s node dependencies, and install them
RUN npm install

# COPY in the rest of your app’s source code from your host to your image filesystem
COPY ./bulletin-board-app .

# The CMD directive is our first example of specifying some metadata in our image
# that describes how to run a container based off of this image.
CMD [ "npm", "start"]