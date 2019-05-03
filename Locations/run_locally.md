### Instructions to run the restaurant location application locally ### <br>
<br>
1. Go to the root directory of the project and execute a pwd command <br>
2. Set go path for proper execution <br>

```
export GOPATH=output/of/pwd/command 
export PATH=$PATH:$GOPATH/bin
```
<br>
3. Run the go files

```
go run *.go
````
<br>
4. Running the above command should give something as follows

```
Riak Server1: http://riakDNSaddress.us-west-2.elb.amazonaws.com:8000
Riak Server2: http://riakDNSaddress.us-west-2.elb.amazonaws.com:8000
Riak Server3: http://riakDNSaddress.us-west-2.elb.amazonaws.com:8000
[RIAK DEBUG] GET: http://riakLoadBalancer-265740929.us-west-2.elb.amazonaws.com:8000/ping => OK
2019/05/03 02:24:02 Riak Ping Server1:  OK
[RIAK DEBUG] GET: http://riakLoadBalancer-265740929.us-west-2.elb.amazonaws.com:8000/ping => OK
2019/05/03 02:24:02 Riak Ping Server2:  OK
[RIAK DEBUG] GET: http://riakLoadBalancer-265740929.us-west-2.elb.amazonaws.com:8000/ping => OK
2019/05/03 02:24:02 Riak Ping Server3:  OK
[negroni] listening on :8000
```
<br>
5. Ping the application to check if it is working by opening browser and going to http://localhost:8000/ping. it should show the following

```
{
  "Test": "Location API ping Handler working!!"
}
```
