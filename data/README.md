
# data
    import "github.com/britannic/blacklist/data"

Package data provides downloaded and configured data processing methods






## func DiffArray
``` go
func DiffArray(a, b []string) (diff []string)
```
DiffArray returns the delta of two arrays


## func GetExcludes
``` go
func GetExcludes(b c.Blacklist) (e c.Dict)
```
GetExcludes returns a map[string]int of excludes


## func GetHTTP
``` go
func GetHTTP(URL string) (body []byte, err error)
```
GetHTTP creates http requests to download data


## func GetIncludes
``` go
func GetIncludes(n *c.Node) (r c.Dict)
```
GetIncludes returns a map[string]int of includes


## func GetList
``` go
func GetList(cf *c.Src) (b []byte)
```
GetList returns a sorted []byte of blacklist entries


## func IsDisabled
``` go
func IsDisabled(d c.Blacklist, root string) bool
```
IsDisabled returns true if blacklist is disabled


## func ListFiles
``` go
func ListFiles(d string) (files []string)
```
ListFiles returns a list of blacklist files


## func Process
``` go
func Process(s *c.Src, dex c.Dict, ex c.Dict, d string) *c.Src
```
Process extracts hosts/domains from downloaded raw content


## func PurgeFiles
``` go
func PurgeFiles(a AreaURLs) error
```
PurgeFiles removes any files that are no longer configured


## func StripPrefix
``` go
func StripPrefix(l string, p string, rx *regx.RGX) (string, bool)
```
StripPrefix returns the modified line and true if it can strip the prefix



## type AreaURLs
``` go
type AreaURLs map[string][]*c.Src
```
AreaURLs is a map of c.Src









### func GetURLs
``` go
func GetURLs(b c.Blacklist) (a AreaURLs)
```
GetURLs returns an array of config.Src structs with active urls










- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)