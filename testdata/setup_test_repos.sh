#!/bin/bash

rm -rf testdata/detached_tag
rm -rf testdata/single_branch_tags

cd testdata

echo 'testdata/detached_tag/ directory does not exist at the root; creating...'
git clone detached_tag.bundle
echo 'done'

echo 'testdata/detached_tag/ directory does not exist at the root; creating...'
git clone single_branch_tags.bundle
echo 'done'
