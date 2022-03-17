#!/bin/bash
#
# Usage: checks.sh output_dir type...

EXIT_CODE=0
OUTPUT_DIR=$1
shift 1
for type in "$@"
do
  case "$type" in
    "lint")
      [ "$(grep -Ec "Grade: A\+?" "${OUTPUT_DIR}report.out")" == "0" ] && echo "Grade must be A or A+" && EXIT_CODE=1
      [ "$(grep -Ec "golint: (100|9[0-9])%" "${OUTPUT_DIR}report.out")" == "0" ] && echo "golint must be greater or equal than 90%" && EXIT_CODE=1
    ;;
    "test-unit")
      [ "$(grep -Ec "FAIL" "${OUTPUT_DIR}unit-run.out")" != "0" ] && echo "Unit tests failing" && EXIT_CODE=1
      [ "$(grep -Ec "total:\s+\(statements\)\s+(100|9[0-9])\.[0-9]%" "${OUTPUT_DIR}parsed_coverage.out")" == "0" ] && echo "Unit tests coverage must be >= 90%" && EXIT_CODE=1
    ;;
    "test-component")
      [ "$(grep -Ec "FAIL" "${OUTPUT_DIR}component-run.out")" != "0" ] && echo "Component tests failing" && EXIT_CODE=1
      [ "$(grep -Ec "\?\s+.+\s+\[no test files\]" "${OUTPUT_DIR}component-run.out")" != "0" ] && echo "Component tests, all lambdas must have the component test" && EXIT_CODE=1
    ;;
    *)
      echo "Type '$1' not defined" && EXIT_CODE=1
    ;;
  esac
  shift 1
done
exit ${EXIT_CODE}
