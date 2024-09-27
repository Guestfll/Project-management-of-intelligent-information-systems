# Project-management-of-intelligent-information-systems

Academic discipline


To work with the project you should compile it into a binary file:

```
go build -o project main.go
```



1. Task 1: ToUpper
    
    ToUpper converts a lowercase Latin character into an uppercase character. If the character is not a lowercase letter of the Latin alphabet, the original character is returned. 

    The program receives a user console input as a string and feeds the converted string to the output.

    The solution is presented without using the standard library.

    To work with this part of project:

    ```
    ./project uppertask
    ```

2. Task 2: Student of the Year 
    
    Student of the Year receives data from a json file, a working example of which is presented in the file list_of_group.json and outputs a table with the name, subjects and grades of the student with the best performance (several, in case there are several excellent students).

    The testdata.py script is prepared for random data generation, to run it use:

    ```
    cd scripts
    ```
    ```
    python3 testdata.py
    ```

    To work with this part of project:

    ```
    ./project beststudents
    ```




