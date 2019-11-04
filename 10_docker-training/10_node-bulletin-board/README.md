1. Make sure you’re in the directory node-bulletin-board in a terminal or powershell, and build your bulletin board image:

```
docker image build -t bulletinboard:1.0 .
```

2. Start a container based on your new image:

```
docker container run --publish 8000:8080 --detach --name bb bulletinboard:1.0
```

    We used a couple of common flags here:

        --publish asks Docker to forward traffic incoming on the host’s port 8000, to the container’s port 8080 (containers have their own private set of ports, so if we want to reach one from the network, we have to forward traffic to it in this way; otherwise, firewall rules will prevent all network traffic from reaching your container, as a default security posture).
        --detach asks Docker to run this container in the background.
        --name lets us specify a name with which we can refer to our container in subsequent commands, in this case bb.

    Also notice, we didn’t specify what process we wanted our container to run. We didn’t have to, since we used the CMD directive when building our Dockerfile; thanks to this, Docker knows to automatically run the process npm start inside our container when it starts up.

3. Visit your application in a browser at localhost:8000. You should see your bulletin board application up and running. At this step, we would normally do everything we could to ensure our container works the way we expected; now would be the time to run unit tests, for example.

4. Once you’re satisfied that your bulletin board container works correctly, delete it:

```
docker container rm --force bb
```
