# Develop in Docker: a Node backend and a React front-end talking to each other

The Node backend is a simple application server which will later serve the React front-end’s build, as well as proxy all external API services to the React front-end.

On bootstrapping my dev environment I decided to give Docker a try and develop both parts in Docker containers.

# Motivation

- First of all, the main goal why docker is used to handle this task is to manage dependencies in the local dev environment. I had the experience of working on multiple projects requiring different Node versions. Despite convenient Node version managing tools like nvm, having to switch version every now and then was really a pain.

- Besides, a dev Dockerfile could also serve as a foundation for later deployment to the staging and production environments.

- And what’s more, every time when you start working, running a one-liner docker-compose up --build is definitely easier than jumping into different folders and then run npm start.

# Goals

- Run both the Node and the React app in its own Docker container

- Communicate between the two apps running in containers.

- Every edit in the local IDE will automatically be reflected in the apps running in containers.

# References:

- https://medium.com/@xiaolishen/develop-in-docker-a-node-backend-and-a-react-front-end-talking-to-each-other-5c522156f634