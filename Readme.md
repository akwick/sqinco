# sqinco - **SQ**L **In**jection **Co**mparison

A small set of test cases to compare the precision of several SQL injection detection tools.

## Installation and Usage

Instal the tool with the *go get* command.

`` go get github.com/akwick/sqinco ``

After the installation, you can to compare the programs.
We have written a shell script which helps to execute three static analysis:
* Safesql
* gas
* Gotcha

If you are aware of more tools, we are glad to hear about them.
To be able to execute all the static analysis tools, you need an installation of the tools on your computer. Check the GitHub repositories of those projects for the installation process.

As an alternative, we offer a docker image which will install these three tools.
You find the docker file in the folder *docker*.

``` 
cd docker 
docker build -t sqinco . 
```

For further information about Docker, study the tutorials for your platform. In some configurations it can be necessary that you execute the docker command with *sudo*.

One variant to run the comparision.sh file is to start the docker image and execute it within the image. 
You can do this by executing following commands. 

```
docker run -t -i sqinco 
cd src/github.com/akwick/sqinco/
./runComparison.sh 
exit
```

## Benchmark tests included

We have currently two benchmark tests:
* The first is in the folder *sqlInjection* and tests all the three tools against the analysis of all the files in this folder.
* The second is in the folder *benchmarks*. This benchmark test analyses a bigger project - the gotcha analysis.


The benchmarks can be executed in the docker image too. Be sure that you have installed the docker image for sqinco. 

```
docker run -t -i sqinco
cd src/github.com/akwick/sqinco/sqlInjections/
go test -bench=. 
exit
```

It is possible to adopt the running time of the benchmark test with a command line flag ``go test -run=XXX -bench=. -benchtime=10s``.

With the same commands it is possible to execute the benachmark tests included in the benchmark folder. Instead of changing in the folder sqlInjection, one has to change in the folder benchmakrs.



