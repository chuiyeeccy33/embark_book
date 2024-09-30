# Part 3: Database

## A. 
1. **Count the total employees for each department code**

    ```sql
    SELECT COUNT(a.ID) 
    FROM Employee AS a 
    INNER JOIN Department AS b ON a.DepartmentID = b.ID 
    GROUP BY b.Code;
    ```

2. **Filter employees with salary range between 3000 and 4000, then sort them by department code and followed by Name**

    This query retrieves employees with salaries between 3000 and 4000, sorted by department code and employee name.

    ```sql
    SELECT * 
    FROM Employee AS a 
    INNER JOIN Department AS b ON a.DepartmentID = b.ID 
    WHERE a.Salary BETWEEN 3000 AND 4000 
    ORDER BY b.Code, a.Name;
    ```

## B. Design a database structure 

### Subject Table

- `subject_id` – Primary key, `INT`
- `subject_name` – `VARCHAR`

### Class Table

- `class_id` – Primary key, `INT`
- `class_name` – `VARCHAR`
- `subject_id` – Foreign key, references `Subject(subject_id)`
- `semester_no` – `INT`

### Student Table

- `id` – Primary key, `INT`
- `student_name` – `VARCHAR`
- `registered` – `BOOL`

### Student Class Table

- `id` – Primary key, `INT`
- `student_id` – Foreign key, references `Student(id)`
- `class_id` – Foreign key, references `Class(class_id)`