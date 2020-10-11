run_unit_test()
{
  gcc dsn.c
  compile_val=$?
  if [[ $compile_val != 0 ]]; then
    exit 1
  else
    echo "Successful?"
  fi

  ./a.out > output.txt
  execution_val=$?
  if [[ $execution_val != 0 ]]; then
    exit 1
  else 
    echo "executed!"
  fi

  diff output.txt test_output/output01.txt
  diff_val=$?
  if [[ $diff_val != 0 ]]; then
    echo "Output mismatch"
  else
    echo "PASS!"
  fi
}


#for i in 'seq -f "%02g" 1 $NUM_INPUT_TESTS';
#do
run_unit_test
#done