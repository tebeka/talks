// +build ignore

#include <Python.h>

int main(int argc, char *argv[]) {
  Py_Initialize();
  PyRun_SimpleString("print('Hello C!')");
  if (Py_FinalizeEx() < 0) {
    return 1;
  }
  return 0;
}
