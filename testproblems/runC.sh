function run_unit_test()
{
  gcc testing.c  2> compile.txt
  compile_val=$?
  # Check to see if we get any warnings with our c code
  if [[ -f "./compile.txt" && -s "./compile.txt" ]]; then
    cat_compile_file=`cat compile.txt`
    echo "$cat_compile_file"
    exit 1
  fi
  if [[ $compile_val != 0 ]]; then
    echo "compile val wrong?"
    exit 1
  fi

  # We need to capture the shell error too since segmentation fault is caused by the shell running the command ./a.out
  { ./a.out; }  2> error.txt
  execution_val=$?
  if [[ $execution_val != 0 ]]; then
    error_file=`error.txt`
    echo "$error_file"
    exit 1
  fi
}

run_unit_test