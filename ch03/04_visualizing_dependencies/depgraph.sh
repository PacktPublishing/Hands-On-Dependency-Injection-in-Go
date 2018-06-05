#!/usr/bin/env bash

# Note:
# This script should be run in the base directory of the project/service

# Inputs
#
# This cuts down on typing by allowing you to enter only the sub-directory you wish to graph; instead of the entire
# package
prefix="./"
PKG=${1#$prefix}

# Constants
#
# Save the file on the desktop (so it's easy to find)
DEST_FILE=~/Desktop/depgraph.png

# Calculate the package in the current directory and assume this is the base or project package
BASE_PKG=$(go list)

EXCLUSIONS="$BASE_PKG/vendor"

# Generate
godepgraph -s \
        -o "$BASE_PKG" \
        $BASE_PKG/${PKG} | dot -Tpng -o $DEST_FILE

# Open
open $DEST_FILE