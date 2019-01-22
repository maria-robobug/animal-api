#!/usr/bin/env bash

git fetch --tags
latest=$(git tag  -l --merged master --sort='-*authordate' | head -n1)

latest=$(git tag  -l --merged master --sort='-*authordate' | head -n1)
semver_parts=(${latest//./ })
major=${semver_parts[0]}
minor=${semver_parts[1]}
patch=${semver_parts[2]}

version=${major}.${minor}.$((patch+1))

echo $version
