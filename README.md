# goLogger
# Notes
## How to use Windows named pipe in golang ?
### Windows provide the library to call Windows API
1. [https://github.com/natefinch/npipe](https://github.com/natefinch/npipe) The author is no longer maintained => [issues-16655](https://github.com/golang/go/issues/16655)
1. [https://github.com/Microsoft/go-winio](https://github.com/Microsoft/go-winio)

## How to stop net.Listener.Accept() ?
### Use net.Listener.Close()
> ### net.Listener document
> ```golang
> type Listener interface {
>     // Accept waits for and returns the next connection to the listener.
>     Accept() (Conn, error)
>
>     // Close closes the listener.
>     // Any blocked Accept operations will be unblocked and return errors.
>     Close() error
>
>     // Addr returns the listener's network address.
>     Addr() Addr
> }
> ```
 Base on [net.Listener document](https://golang.org/pkg/net/#Listener), net.Listener.Accept() will block the thread and wait for the connection. If we want to exist the program, we must need a mechanism that will stop net.Listener.Accept(). We can call **Listener.Close()** to unblock **Listener.Accept()** and get return errors.

## Stop event to all goroutine
### use close(stopChan)
