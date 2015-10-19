# Goony

&ndash; Simple Load Testing


## Contribute

I prefer having my development projects outside of my Go workspace.
The file `link_project_in_workspace.sh` has been added for that.
It will create a symlink in your `$GOPATH` to the directory you cloned this repository into.
This way all import statements will work as expected.

### Running the Tests

Goony uses [gom](https://github.com/mattn/gom) for its dependencies.
Make sure to run `gom -test install` before running the tests.

Execute `gom exec ginkgo -r` to run the test suite.
