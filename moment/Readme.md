# moment
--
    import "github.com/shiaho/otto-package/moment"

Package moment contains the source for the JavaScript date library.

    import (
    	_ "github.com/shiaho/otto-package/moment"
    )
    // Every Otto runtime will now include moment

http://momentjs.com/

https://github.com/moment/moment/

By importing this package, you'll automatically load moment every time you
create a new Otto runtime.

To prevent this behavior, you can do the following:

    import (
    	"github.com/shiaho/otto-package/moment"
    )

    func init() {
    	moment.Disable()
    }

## Usage

#### func  Disable

```go
func Disable()
```
Disable moment runtime inclusion.

#### func  Enable

```go
func Enable()
```
Enable moment runtime inclusion.

#### func  Source

```go
func Source() string
```
Source returns the moment source.

--
**godocdown** http://github.com/robertkrimen/godocdown
