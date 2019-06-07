
```bash
bazel version
# Build label: 0.26.1

cat .bazelrc
# build --action_env=HELLO=WORLD
# run --action_env=HELLO=WORLD
# test --test_env=HELLO=WORLD
```

```bash
bazel run //:go
bazel run //:go --action_env=HELLO=WORLD
# HELLO env mismatch: ""

bazel test //:go_test
bazel test //:go_test --test_env=HELLO=WORLD
# Hello, WORLD
# //:go_test   PASSED
```

