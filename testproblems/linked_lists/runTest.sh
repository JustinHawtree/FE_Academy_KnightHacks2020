run_unit_test()
{
  local testcase_file=$1
  local output_file=$2

  gcc user.c $testcase_file 2> compile.txt
  compile_val=$?
  if [[ $compile_val != 0 ]]; then
    cat_compile_file=`cat compile.txt`
    echo "$cat_compile_file"
    #echo "compilation $testcase_file"
    exit 1
  fi

  # We need to capture the shell error too since segmentation fault is caused by the shell running the command ./a.out
  { ./a.out; } > output.txt 2> err.txt
  execution_val=$?
  if [[ $execution_val != 0 ]]; then
    cat_err_file=`cat err.txt`
    echo "$cat_err_file"
    #echo "execution $testcase_file"
    exit 1
  fi

  diff output.txt testcase_output/$output_file > /dev/null
  diff_val=$?
  if [[ $diff_val != 0 ]]; then
    echo "Output mismatch"
    #echo "diff $testcase_file"
    exit 1
  fi
}


for i in `seq -f "%02g" 1 3`;
do
  run_unit_test "testcase$i.c" "testcase$i-output.txt"
done

echo "Passed all Test Cases!"