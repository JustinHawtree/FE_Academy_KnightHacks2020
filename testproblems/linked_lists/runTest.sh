function run_unit_test()
{
  local testcase_file=$1
  local output_file=$2

  gcc user.c $testcase_file 2> compile.txt
  compile_val=$?
  # Check to see if we get any warnings with our c code
  if [[ -f "./compile.txt" && -s "./compile.txt" ]]; then
    cat_compile_file=`cat compile.txt`
    echo "$cat_compile_file"
    exit 1
  fi
  if [[ $compile_val != 0 ]]; then
    exit 1
  fi

  # We need to capture the shell error too since segmentation fault is caused by the shell running the command ./a.out
  { ./a.out; } > output.txt 2>&1
  execution_val=$?
  if [[ $execution_val != 0 ]]; then
    cat_output_file=`cat output.txt`
    echo "$cat_output_file"
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


for i in `seq -f "%02g" 1 4`;
do
  run_unit_test "testcase$i.c" "testcase$i-output.txt"
done


function run_valgrind_test()
{
  local testcase_file=$1

  gcc user.c $testcase_file 2> compile.txt
  compile_val=$?
  # Check to see if we get any warnings with our c code
  if [[ -f "./compile.txt" && -s "./compile.txt" ]]; then
    cat_compile_file=`cat compile.txt`
    echo "$cat_compile_file"
    exit 1
  fi
  if [[ $compile_val != 0 ]]; then
    exit 1
  fi


  valgrind --leak-check=yes ./a.out > output.txt 2> err.log
  execution_val=$?
  if [[ $execution_val != 0 ]]; then
    exit 1
  fi


  grep --silent "no leaks are possible" err.log
  valgrindleak=$?
  if [[ $valgrindleak != 0 ]]; then
    echo "Memory Leak Detected: make sure to free the dequeued node"
    exit 1
  fi
}

for i in `seq -f "%02g" 4 4`
do
  run_valgrind_test "testcase$i.c"
done

echo "Passed all Test Cases!"