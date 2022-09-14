# SYNAOS Development Kata: A Reliable Service

Development kata to provide a reliable way of operating a service.

[A code kata is an exercise in programming which helps a programmer hone their skills through practice and repetition.](https://en.wikipedia.org/wiki/Kata_(programming))

* [Frame conditions](#frame-conditions)
* [Acceptance-Criteria](#acceptance-criteria)
* [License](#license)

## Frame conditions

1. You have exact 2 hours of time - no minute longer.

   If you reach this time limit stop your work immediately.
   It is one part of the kata to respect this time limit. 🙂

2. There are no restrictions on how to use the provided time.
   If you want to code the entire time, take a break or a cigarette; it’s up to you.

3. This is a real world situation. You are allowed to consult the Internet, use every library you want, call a friend ...

4. It is fully up to you how you operate this one. If you choose an AWS EC2 instance, a Kubernetes cluster, your local machine, ...; you choose download the binary directly, build your own docker build pipeline, ...; everything could make sense. 

5. Keep the following priorities in mind while you implementing - in the **mentioned order**:
    1. Reliability
    2. Maintainability
    3. Functionality

6. The service which should be operated itself can be found in [`pkg/service`](pkg/service).
   
   1. It is fully functioning as it is.
 
   2. By default, it is listening to [127.0.0.1:8000](http://127.0.0.1:8009). It can be overwritten using the environment variable `KATA_SERVICE_ADDRESS`.

   3. It is serving the following endpoints:
      * `/` a regular endpoint which prints out a simple response text.
      * `/crash` an endpoint which will let the service crash immediately.

   4. It is written in [Go](https://go.dev).

   5. It can be retrieved...
      * ...either as the latest binary retrieved directly from [our pre-packaged releases](https://github.com/synaos/reliable-service-kata/releases)
      * ...or building it yourself using (if at [Go 1.19+](https://go.dev/dl/) is installed):
        ```shell
        # Linux/macOS
        $ go build -o service ./pkg/service
        # Windows
        $ go build -o sevice.exe ./pkg/service
        ```  

## Acceptance criteria

1. The service is handling the endpoint [`/`](http://127.0.0.1:8000) continuously.
2. The service is handling the endpoint [`/crash`](http://127.0.0.1:8000/crash) continuously.

## License

See [LICENSE](LICENSE) file.