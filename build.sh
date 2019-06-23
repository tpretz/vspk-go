#!/bin/bash

rm -rf vspk/*

while read tag; do
  echo "processing $tag"

  monogen -L go -c vspk.ini -f ~/git/vsd-api-specifications -b $tag
  if [ $? -eq 0 ]; then
    mv go/vspk vspk/${tag:1}
  fi
  rm -r go
done < <(git -C ~/git/vsd-api-specifications tag | grep '^r')
