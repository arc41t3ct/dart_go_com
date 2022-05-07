#!/usr/local/bin/bash


DART_SDK_INCLUDE_PATH="$HOME/pkgs/flutter/bin/cache/dart-sdk/include/"

find . ! -name 'update_dart_sdk_includes.sh' -type f -exec rm -f {} +
find . ! -name 'update_dart_sdk_includes.sh' -type d -exec rm -r {} +

echo "Copying from $DART_SDK_INCLUDE_PATH contents"

cp -R $DART_SDK_INCLUDE_PATH .
