#!/bin/bash
echo "Installing the netrogue package..."
targ=$GOPATH/src/netrogue
mkdir -p $targ
orig=$(pwd)
cp -R $orig/* $targ/
cd $targ
go install
cd $orig
echo "Done."

