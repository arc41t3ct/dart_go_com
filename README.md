# dart_go_com
Dart Go FFI

## Notes

The include folder must match the dart sdk that comes with flutter when we want to distribute the app. That is why a shell script is included in the include folder which when run locally needs to match the directory where the dart sdk is installed. Whoever uses this package needs to either fork it and change the path or just do it locally or however you like.

We could change it to export the path to the sdk before running the script.
