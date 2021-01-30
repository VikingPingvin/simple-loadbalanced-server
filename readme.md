# Load Balanced Web server with simple web app

# Instructions
```bash
docker-compose up --scale app=2
```

```bash
           Name                          Command               State         Ports
-----------------------------------------------------------------------------------------        
intland_task_app_1            ./simpleapp                      Up      80/tcp
intland_task_app_2            ./simpleapp                      Up      80/tcp
intland_task_loadbalancer_1   /docker-entrypoint.sh ngin ...   Up      0.0.0.0:80->80/tcp 
```

# Services
### Application  
Simple Go web application that reads the `/etc/hostname` file and returns it as a simple bytestream (alternative long is `cat /proc/self/cgroup | grep ...`).  
Default HTTP response is used, HTTP Headers are not touched.

Possible improvements: 
- Conditional build variable for webserver port (*main.go line 17*)
- HTTP HEADER with proper informations

### LoadBalancer
Simple NGINX Webserver that proxies the URL to the default Docker DNS resolver using default round-robin setting.


#### Daniel Schreiber | 2021.01.30