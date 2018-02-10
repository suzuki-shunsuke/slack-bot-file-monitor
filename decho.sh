decho() {
  echo "+ $@"
  eval "$@"
}

dexec() {
  echo "+ $@"
  exec "$@"
}
