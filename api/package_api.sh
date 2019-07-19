#!/usr/bin/env bash

TARGET=out
ZIP_NAME=myfunct.zip

if [[ -d $TARGET ]]; then
  rm -rf $TARGET
fi
mkdir -p $TARGET
pip3 install -r requirements.txt --target=$TARGET
cp -r src/  $TARGET/

pushd $TARGET || exit 1
# clean up cache files
find . | grep -E "(__pycache__|\.pyc|\.pyo$)" | xargs rm -rf
if [[ -f $ZIP_NAME ]]; then
  rm $ZIP_NAME
fi
# shellcheck disable=SC2046
zip -r $ZIP_NAME $(ls)

popd ||: