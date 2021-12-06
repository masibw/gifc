# gifc
gifc generates github issues from TODO comment in source code.

# Install
```bash
go install github.com/masibw/gifc@latest
```
# Usage
You must provide your github private token(with repo scope).

``` bash
GITHUB_TOKEN=your_token gifc
```

If the source code contains the following TODO comment, a github issue will be created automatically.

```main.go
// TODO: we must implement this.
```
