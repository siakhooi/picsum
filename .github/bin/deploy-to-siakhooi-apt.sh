#!/bin/bash

readonly apt_repo_url="https://${PUBLISH_TO_GITHUB_REPO_TOKEN}@github.com/siakhooi/apt.git"
readonly apt_repo_branch=main
readonly apt_repo_directory=apt
readonly apt_repo_path=docs/pool/main/binary-amd64
readonly git_commit_email=picsum@siakhooi.github.io
readonly git_commit_user=picsum
git_commit_message="picsum: Auto deploy [$(date)]"
readonly git_commit_message

set -e

debian_package_path=$(ls ./dist/siakhooi-picsum_*_amd64.deb)
debian_package_source_path=$(realpath "$debian_package_path")
debian_package_filename=$(basename "$debian_package_path")
target_debian_package_path=$apt_repo_path/$debian_package_filename
working_directory=$(mktemp -d)

(
  cd "$working_directory"
  git config --global user.email "$git_commit_email"
  git config --global user.name "$git_commit_user"

  git clone -n --depth=1 -b "$apt_repo_branch" "$apt_repo_url" "$apt_repo_directory"
  cd "$apt_repo_directory"
  git remote set-url origin "$apt_repo_url"
  git restore --staged .
  mkdir -p $apt_repo_path
  cp -v "$debian_package_source_path" "$target_debian_package_path"
  git add "$target_debian_package_path"
  git status
  git commit -m "$git_commit_message"
  git push
)
find .
