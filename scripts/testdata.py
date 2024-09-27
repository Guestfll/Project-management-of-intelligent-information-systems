import json
from faker import Faker
import random

# Initialize Faker
fake = Faker()

# Number of students and teachers
num_students = 33
num_teachers = 7  # Updated to match the total subjects count

# Subjects list
subjects = ["Mathematics", "Physics", "Chemistry", "History", "Literature", "Biology", "Computer Science"]

# Generate teachers with unique subjects
teachers = [
    {
        "first_name": fake.first_name(), 
        "last_name": fake.last_name(),
        "subject": subject,
        "date_of_birth": fake.date_of_birth(minimum_age=30, maximum_age=65).strftime("%Y-%m-%d")
    }
    for subject in subjects  # Each subject is assigned a unique teacher
]

# Generate a single exam date for each subject
subject_exam_dates = {
    subject: fake.date_this_year(before_today=True, after_today=False).strftime("%Y-%m-%d")
    for subject in subjects
}

# Generate students with their details and complete credit book
students = []
for _ in range(num_students):
    # Basic student info
    student = {
        "first_name": fake.first_name(),
        "last_name": fake.last_name(),
        "date_of_birth": fake.date_of_birth(minimum_age=18, maximum_age=22).strftime("%Y-%m-%d"),
        "credit book": []
    }
    
    # Add all subjects to the student's credit book
    for subject in subjects:
        exam_date = subject_exam_dates[subject]
        grade = random.choice(["3", "4", "5"])
        # Find the teacher assigned to this subject
        teacher = next(teacher for teacher in teachers if teacher["subject"] == subject)
        
        # Adding the exam record
        exam_record = {
            "subject": subject,
            "exam_date": exam_date,
            "teacher": f"{teacher['first_name']} {teacher['last_name']}",
            "grade": grade
        }
        student["credit book"].append(exam_record)
    
    students.append(student)

# Create the final data structure
group_data = {
    "students": students,
    "teachers": teachers
}

# Save to JSON file
final_full_group_path = "../list_of_group.json"
with open(final_full_group_path, "w") as json_file:
    json.dump(group_data, json_file, indent=4)

print(f"Data successfully saved to {final_full_group_path}")
